package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"time"
)

const responsePrefix = "ZBXDZ\x01"

type ZabbixConnection struct {
	c net.Conn
}

func CreateConnection(server string, port int) (*ZabbixConnection, error) {
	con, err := getConnection(server, port)
	return &ZabbixConnection{c: con}, err
}
func getConnection(server string, port int) (net.Conn, error) {
	type DialResp struct {
		Conn  net.Conn
		Error error
	}
	dialCh := make(chan DialResp, 0)
	go func() {
		conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", server, port))
		defer close(dialCh)
		dialCh <- DialResp{Conn: conn, Error: err}
	}()
	select {
	case dial := <-dialCh:
		return dial.Conn, dial.Error
	case <-time.After(5 * time.Second):
		return nil, fmt.Errorf("connection to %v:%v timeout", server, port)
	}
}
func (c *ZabbixConnection) Send(preparedPackage ZabbixPreparedPackage) (string, error) {
	if _, err := c.c.Write(preparedPackage); err != nil {
		return "", err
	}
	res := make([]byte, 1024)
	res, _ = ioutil.ReadAll(c.c)

	return string(res[len(responsePrefix):]), nil

}
func (c *ZabbixConnection) Close() {
	err := c.c.Close()
	fmt.Println(err)
}
