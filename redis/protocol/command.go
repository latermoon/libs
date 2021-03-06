package protocol

import (
	"bytes"
	"encoding/json"
)

type Command [][]byte

func NewCommand(args ...[]byte) Command {
	return Command(args)
}

func (c Command) Len() int {
	return len(c)
}

func (c Command) Bytes() []byte {
	buf := &bytes.Buffer{}
	buf.WriteByte('*')
	argCount := c.Len()
	buf.WriteString(itoa(argCount)) //<number of arguments>
	buf.WriteString(CRLF)
	for i := 0; i < argCount; i++ {
		buf.WriteByte('$')
		buf.WriteString(itoa(len(c[i]))) //<number of bytes of argument i>
		buf.WriteString(CRLF)
		buf.Write(c[i]) //<argument data>
		buf.WriteString(CRLF)
	}
	return buf.Bytes()
}

func (c Command) String() string {
	arr := make([]string, len(c))
	for i := range c {
		arr[i] = string(c[i])
	}
	b, _ := json.Marshal(arr)
	return string(b)
}
