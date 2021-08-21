package main

import "zabbix-http/config"

func main() {
	RunServer(config.GetConfig())
}
