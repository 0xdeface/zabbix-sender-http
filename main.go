package main

import (
	configurator "zabbix-http/config"
	"zabbix-http/internal/connection"
	"zabbix-http/internal/domain"
	"zabbix-http/internal/http"
	logger "zabbix-http/internal/logger"
	"zabbix-http/internal/zabbix"
)

func main() {
	log := logger.NewLogger(domain.DEBUG)
	config := configurator.GetConfig(log)
	con := connection.CreateConnection(config.ServerAddr, config.ZabbixPort)
	packet := zabbix.NewPacket([]zabbix.Message{})
	zabbixSender := domain.NewZabbixSender(con, packet, log)
	http.RunServer(config, zabbixSender, log)
}
