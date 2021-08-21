package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"zabbix-http/config"
)

var required = []string{"server", "key", "value"}

func RunServer(cfg *config.Config) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Body)
		q := request.URL.Query()
		for _, val := range required {
			if _, ok := q[val]; !ok {
				fmt.Fprintf(writer, "%v: shouldn t be empty", val)
				return
			}
		}
		zabbixPort, _ := strconv.Atoi(cfg.ZabbixPort)
		con, err := CreateConnection(cfg.ServerAddr, zabbixPort)
		defer con.Close()
		if err != nil {
			fmt.Fprint(writer, err)
			return
		}
		z := NewZabbixPacket([]ZabbixMessage{})
		z.AddMessage(q["server"][0], q["key"][0], q["value"][0])
		if response, err := con.Send(z.Prepare()); err != nil {
			fmt.Fprint(writer, err)
			return
		} else {
			fmt.Fprint(writer, response)
		}

	})

	err := http.ListenAndServe(":"+cfg.HttpPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
