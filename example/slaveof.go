package main

import (
	. "../redis/protocol"
	"../redis/slaveof"
	"log"
	"net"
	"strings"
)

func main() {
	log.Println("hello")

	cmd := NewCommand([]byte("SYNC"))
	log.Println(strings.Replace(string(cmd.Bytes()), CRLF, " ", -1), cmd.Bytes())

	tectSlaveOf()
}

func tectSlaveOf() {
	client := slaveof.New(":6379", &handler{})
	if err := client.Connect(); err != nil {
		log.Panicln(err)
	}
}

type handler struct {
	slaveof.SlaveOfHandler
}

func (h *handler) OnCommand(cmd Command) {
	log.Println(cmd)
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
