package http

import (
	"fmt"
	"log"
	"net/http"
	"zabbix-http/config"
	"zabbix-http/internal/domain"
)

var required = []string{"server", "key", "value"}

func RunServer(cfg *config.Config, zabbixSender domain.ZabbixSenderPort, logger domain.Logger) {
	handler(zabbixSender, logger)
	err := http.ListenAndServe(":"+cfg.HttpPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(zabbixSender domain.ZabbixSenderPort, logger domain.Logger) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		q := request.URL.Query()

		for _, val := range required {
			if _, ok := q[val]; !ok {
				logger.Log(domain.INFO, fmt.Sprintf("request with empty parameter %v", val))
				fmt.Fprintf(writer, "%v: shouldn t be empty \n", val)
				return
			}
		}

		response, err := zabbixSender.SendToZabbix(q["server"][0], q["key"][0], q["value"][0])
		if err != nil {
			logger.Log(domain.ERROR, err.Error())
			fmt.Fprintln(writer, "an error occurred while sending the request, see log for detail")
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.Write(response)

	})
}
