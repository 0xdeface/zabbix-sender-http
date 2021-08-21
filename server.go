package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var required = []string{"server", "key", "value"}

func RunServer(cfg *Config) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Body)
		q := request.URL.Query()
		for _, val := range required {
			if _, ok := q[val]; !ok {
				fmt.Fprintf(writer, "%v: shouldn t be empty", val)
				return
			}
		}
		zabbixPort, _ := strconv.Atoi(cfg.zabbixPort)
		con, err := CreateConnection(cfg.serverAddr, zabbixPort)
		if err != nil {
			fmt.Fprint(writer, err)
			return
		}
		z := NewZabbixPacket([]ZabbixMessage{})
		z.AddMessage(q["server"][0], q["key"][0], q["value"][0])
		con.Send(z.Prepare())
		con.Close()
	})

	err := http.ListenAndServe(":"+cfg.httpPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
