package slaveof

import (
	. "../protocol"
	"net"
	"os"
)

type SlaveOfHandler interface {
	OnCommand(cmd Command)
}

// Client
type Client struct {
	address string
	handler SlaveOfHandler
	session *Session
}

func New(address string, handler SlaveOfHandler) *Client {
	return &Client{
		address: address,
		handler: handler,
	}
}

func (c *Client) Connect() error {
	conn, err := net.Dial("tcp", c.address)
	if err != nil {
		return err
	}
	c.session = NewSession(conn)
	cmd := NewCommand([]byte("SYNC"))
	c.session.Write(cmd.Bytes())
	if err := c.recvRDB(); err != nil {
		return err
	}
	if err := c.recvCommand(); err != nil {
		return err
	}
	return nil
}

func (c *Client) Disconnect() error {
	c.session.Close()
	return nil
}

func (c *Client) recvRDB() error {
	c.session.ReadRDB(nil)
	return nil
}

func (c *Client) recvCommand() error {
	for {
		cmd, err := c.session.ReadCommand()
		if err != nil {
			return err
		}
		c.handler.OnCommand(cmd)
	}
}
