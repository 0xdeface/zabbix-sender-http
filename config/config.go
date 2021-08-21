package config

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type Color string

func colorize(color Color, message string) {
	if runtime.GOOS == "windows" {
		fmt.Println(message)
		return
	}
	fmt.Println(string(color), message, string(ColorReset))
}

type Config struct {
	ServerAddr string
	ZabbixPort string
	HttpPort   string
}
type Params struct {
	name        string
	description string
	predefined  string
}

func (p *Params) envNotation() string {
	return strings.ReplaceAll(strings.ToUpper(p.name), "-", "_")
}

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

func GetConfig() *Config {
	params := []Params{
		{
			name:        "zabbix-server",
			description: "set zabbix server address, default 127.0.0.1",
			predefined:  "127.0.0.1",
		},
		{
			name:        "zabbix-port",
			description: "set zabbix server port, default 10051",
			predefined:  "10051",
		},
		{
			name:        "http-port",
			description: "http server port, default 8080",
			predefined:  "8080",
		},
	}
	filledParams := make(map[string]*string, len(params))
	for _, param := range params {
		filledParams[param.name] = flag.String(param.name, envOrDefault(param.envNotation(), param.predefined), param.description)
	}
	flag.Parse()
	config := &Config{
		ServerAddr: *filledParams["zabbix-server"],
		ZabbixPort: *filledParams["zabbix-port"],
		HttpPort:   *filledParams["http-port"],
	}
	announceConfig(config)
	return config
}
func announceConfig(cfg *Config) {
	colorize(ColorYellow, fmt.Sprintf("zabbix server addr: %v:%v", cfg.ServerAddr, cfg.ZabbixPort))
	colorize(ColorYellow, fmt.Sprintf("http server port: %v", cfg.HttpPort))
}

func envOrDefault(env, predefined string) string {
	if val, exist := os.LookupEnv(env); exist {
		return val
	}
	return predefined
}

