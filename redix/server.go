package redix

import (
	"net"
	"strings"
)

var (
	handlers = map[string]Handler{}
)

const (
	CR   = '\r'
	LF   = '\n'
	CRLF = "\r\n"
)

// Unregistered command
const Other = ""

type Handler func(Command) Reply

func init() {
	Handle(Other, func(c Command) Reply {
		return ErrorReply("Unregistered " + strings.ToUpper(string(c[0])))
	})
}

// Register command
func Handle(cmd string, handler Handler) {
	handlers[strings.ToUpper(cmd)] = handler
}

func ListenAndServe(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleConnection(NewSession(conn))
	}
	return nil
}

func handleConnection(session *Session) {
	for {
		cmd, err := session.ReadCommand()
		if err != nil {
			session.Close()
			break
		}

		name := strings.ToUpper(string(cmd[0]))
		handler, ok := handlers[name]
		if !ok {
			handler, ok = handlers[Other]
		}

		if ok {
			reply := handler(cmd)
			if reply != nil {
				if _, err := session.Write(reply.Bytes()); err != nil {
					session.Close()
					break
				}
			}
		}
	}
}
