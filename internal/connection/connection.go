package connection

import (
	"fmt"
	"io/ioutil"
	"net"
	"time"
)

type Connection struct {
	c      net.Conn
	server string
	port   string
}

func CreateConnection(server string, port string) *Connection {
	return &Connection{c: nil, server: server, port: port}
}
func (c *Connection) getConnection() error {
	type DialResp struct {
		Conn  net.Conn
		Error error
	}
	dialCh := make(chan DialResp, 0)
	go func() {
		conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", c.server, c.port))
		defer close(dialCh)
		dialCh <- DialResp{Conn: conn, Error: err}
	}()
	select {
	case dial := <-dialCh:
		if dial.Error != nil {
			return dial.Error
		}
		c.c = dial.Conn
		return nil
	case <-time.After(5 * time.Second):
		return fmt.Errorf("connection to %v:%v timeout", c.server, c.port)
	}
}

func (c *Connection) Send(preparedPackage []byte) ([]byte, error) {
	if err := c.getConnection(); err != nil {
		return nil, err
	}
	defer c.Close()
	if _, err := c.c.Write(preparedPackage); err != nil {
		return nil, err
	}
	res := make([]byte, 1024)
	res, err := ioutil.ReadAll(c.c)
	return res, err

}
func (c *Connection) Close() error {
	return c.c.Close()
}
