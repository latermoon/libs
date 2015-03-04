package protocol

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"sync"
)

//
// session.Write(cmd.Bytes())
// session.Write(reply.Bytes())
// cmd, err := codec.Read(session, &cmd)
// reply, err := codec.Read(session, &reply)
// rdb, err := codec.ReadRDB(sessoin, rdbHandler)
type Session struct {
	net.Conn
	rw    *bufio.Reader
	wlock sync.Mutex
	rlock sync.Mutex
}

func NewSession(conn net.Conn) *Session {
	return &Session{
		Conn: conn,
		rw:   bufio.NewReader(conn),
	}
}

func (s *Session) ReadCommand() (Command, error) {
	s.rlock.Lock()
	defer s.rlock.Unlock()
	// Read ( *<number of arguments> CR LF )
	if err := s.skipByte('*'); err != nil { // io.EOF
		return nil, err
	}
	// number of arguments
	argCount, err := s.readInt()
	if err != nil {
		return nil, err
	}
	args := make([][]byte, argCount)
	for i := 0; i < argCount; i++ {
		// Read ( $<number of bytes of argument 1> CR LF )
		if err := s.skipByte('$'); err != nil {
			return nil, err
		}

		var argSize int
		argSize, err = s.readInt()
		if err != nil {
			return nil, err
		}

		// Read ( <argument data> CR LF )
		args[i] = make([]byte, argSize)
		_, err = io.ReadFull(s, args[i])
		if err != nil {
			return nil, err
		}

		err = s.skipBytes([]byte{CR, LF})
		if err != nil {
			return nil, err
		}
	}
	return NewCommand(args...), nil
}

func (s *Session) ReadReply() (Reply, error) {
	s.rlock.Lock()
	defer s.rlock.Unlock()

	reader := s.rw
	c, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}

	switch c {
	case '+':
		if value, err := s.readString(); err != nil {
			return nil, err
		} else {
			return StatusReply(value), nil
		}
	case '-':
		if value, err := s.readString(); err != nil {
			return nil, err
		} else {
			return ErrorReply(value), nil
		}
	case ':':
		if value, err := s.readInt(); err != nil {
			return nil, err
		} else {
			return IntegerReply(value), nil
		}
	case '$':
		var bufsize int
		bufsize, err := s.readInt()
		if err != nil {
			return nil, err
		}
		buf := make([]byte, bufsize)
		_, err = io.ReadFull(s, buf)
		if err != nil {
			return nil, err
		}
		s.skipBytes([]byte{CR, LF})
		return BulkReply(buf), nil
	case '*':
		var argCount int
		argCount, err = s.readInt()
		if err != nil {
			return nil, err
		}
		if argCount == -1 {
			return MultiBulkReply(nil), nil // *-1
		} else {
			args := make([]interface{}, argCount)
			for i := 0; i < argCount; i++ {
				// TODO multi bulk 的类型 $和:
				if err := s.skipByte('$'); err != nil {
					return nil, err
				}
				if argSize, err := s.readInt(); err != nil {
					return nil, err
				} else if argSize == -1 {
					args[i] = nil
				} else {
					arg := make([]byte, argSize)
					_, err = io.ReadFull(s, arg)
					if err != nil {
						return nil, err
					}
					args[i] = arg
				}
				s.skipBytes([]byte{CR, LF})
			}
			return MultiBulkReply(args), nil
		}
	default:
		err = errors.New("Bad Reply Flag:" + string([]byte{c}))
		return nil, err
	}
}

func (s *Session) ReadRDB(w io.Writer) (err error) {
	// Read ( $<number of bytes of RDB> CR LF )
	if err = s.skipByte('$'); err != nil {
		return
	}

	var rdbSize int64
	if rdbSize, err = s.readInt64(); err != nil {
		return
	}

	var c byte
	for i := int64(0); i < rdbSize; i++ {
		c, err = s.rw.ReadByte()
		if err != nil {
			return
		}
		if w != nil {
			w.Write([]byte{c})
		}
	}
	return
}

func (s *Session) Close() error {
	return nil
}

func (s *Session) String() string {
	return fmt.Sprintf("<Session:%s>", s.RemoteAddr())
}

// ============================================

func (s *Session) Read(p []byte) (n int, err error) {
	return s.rw.Read(p)
}

func (s *Session) ReadByte() (c byte, err error) {
	return s.rw.ReadByte()
}

func (s *Session) PeekByte() (c byte, err error) {
	if b, e := s.rw.Peek(1); e == nil {
		c = b[0]
	}
	return
}

// ============================================

func (s *Session) skipByte(c byte) (err error) {
	var tmp byte
	tmp, err = s.rw.ReadByte()
	if err != nil {
		return
	}
	if tmp != c {
		err = errors.New(fmt.Sprintf("Illegal Byte [%d] != [%d]", tmp, c))
	}
	return
}

func (s *Session) skipBytes(bs []byte) (err error) {
	for _, c := range bs {
		err = s.skipByte(c)
		if err != nil {
			break
		}
	}
	return
}

func (s *Session) readLine() (line []byte, err error) {
	line, err = s.rw.ReadSlice(LF)
	if err == bufio.ErrBufferFull {
		return nil, errors.New("line too long")
	}
	if err != nil {
		return
	}
	i := len(line) - 2
	if i < 0 || line[i] != CR {
		err = errors.New("bad line terminator:" + string(line))
	}
	return line[:i], nil
}

func (s *Session) readString() (str string, err error) {
	var line []byte
	if line, err = s.readLine(); err != nil {
		return
	}
	str = string(line)
	return
}

func (s *Session) readInt() (i int, err error) {
	var line string
	if line, err = s.readString(); err != nil {
		return
	}
	i, err = strconv.Atoi(line)
	return
}

func (s *Session) readInt64() (i int64, err error) {
	var line string
	if line, err = s.readString(); err != nil {
		return
	}
	i, err = strconv.ParseInt(line, 10, 64)
	return
}
