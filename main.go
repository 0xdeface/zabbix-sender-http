package main

import (
	"context"
	"fmt"
	zabbix "github.com/0xdeface/zabbix/sender"
	"log"
	"os"
	"os/signal"
	"zabbix-http/config"
	"zabbix-http/internal/http"
)

func main() {
	cfg := config.GetConfig()
	errCh := make(chan error, 10)
	zabbixSender := zabbix.NewZabbixSender(cfg.ServerAddr, cfg.ZabbixPort)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	//ctx, cancel := context.WithTimeout(ctx, time.Second * 1)
	go startZabbixSender(ctx, zabbixSender)
	go startHttpServer(ctx, cfg.HttpPort, zabbixSender.MsgCh, errCh)
	go func() {
		for {
			select {
			case e := <-errCh:
				fmt.Println(e)
			case e := <-zabbixSender.ErrCh:
				fmt.Println(e)
			case <-ctx.Done():
				return
			}
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	cancel()
}

func startZabbixSender(ctx context.Context, s *zabbix.Sender) {
	err := s.Start(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}

func startHttpServer(ctx context.Context, port string, msgCh chan zabbix.Message, errCh chan error) {
	http.RunServer(ctx, port, msgCh, errCh)
}
