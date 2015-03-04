package main

import (
	. "../redis/protocol"
	"log"
	"net"
	"strings"
)

func main() {
	log.Println("hello")

	cmd := NewCommand([]byte("SYNC"))
	log.Println(strings.Replace(string(cmd.Bytes()), CRLF, " ", -1), cmd.Bytes())

	testLocal()
}

func testLocal() {
	conn, err := net.Dial("tcp", ":6379")
	if err != nil {
		log.Panicln(err)
	}

	session := NewSession(conn)
	// cmd := NewCommand([]byte("get"), []byte("name"))
	cmd := NewCommand([]byte("HGETALL"), []byte("meeting:users"))
	session.Write(cmd.Bytes())
	reply, _ := session.ReadReply()
	log.Println(formatReply(reply))
}

func formatReply(r Reply) string {
	return strings.Replace(string(r.Bytes()), CRLF, " ", -1)
}
