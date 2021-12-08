package zabbix

import (
	"context"
	"fmt"
)

type Sender struct {
	server string
	port   string
	MsgCh  chan Message
	ErrCh  chan error
	con    Connection
}

func NewZabbixSender(addr, port string) *Sender {
	zs := &Sender{server: addr, port: port}
	zs.MsgCh = make(chan Message, 10)
	zs.ErrCh = make(chan error, 10)
	return zs
}
func (z *Sender) SetErrChan(ech chan error) {
	z.ErrCh = ech
}
func (z *Sender) Start(ctx context.Context) (err error) {
	z.MsgCh = make(chan Message, 1)
	z.con = CreateConnection(z.server, z.port)
	err = z.con.Open()
	if err != nil {
		return err
	}
	go func() {
		for {
			select {
			case m := <-z.MsgCh:
				fmt.Println("received new message", m)
				packet := NewPacket([]Message{m})
				_, err := z.con.Send(packet.Prepare())
				if err != nil {
					z.ErrCh <- err
				}
			case <-ctx.Done():
				err = z.stop()
				close(z.MsgCh)
				if err != nil {
					z.ErrCh <- err
				}
				return
			}
		}
	}()
	close(z.ErrCh)
	return nil
}

func (z *Sender) stop() (err error) {
	defer func() {
		if r, ok := recover().(error); ok {
			err = r
		}
	}()
	fmt.Println(z.con.c)
	return z.con.Close()
}
