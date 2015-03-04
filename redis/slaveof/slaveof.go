package slaveof

import (
	"net"
)

type SlaveOfHandler interface {
	OnCommand(args [][]byte)
}

// Client
type Client struct {
	address string
	handler SlaveOfHandler
	conn    net.Conn
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
	c.conn = conn

	return nil
}

func (c *Client) Disconnect() error {
	return nil
}
