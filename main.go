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
	if cfg.Debug {
		fmt.Printf("\n App started with params:\n %+v\n", cfg)
	}
	errCh := make(chan error, 10)
	zabbixSender := zabbix.NewZabbixSender(cfg.ZabbixHost, cfg.ZabbixPort)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go startZabbixSender(ctx, zabbixSender)
	go http.RunServer(ctx, cfg.HttpPort, zabbixSender.MsgCh, errCh, cfg.Debug)
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
