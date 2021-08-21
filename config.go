package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
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
	serverAddr string
	zabbixPort string
	httpPort   string
}

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

func getConfig() *Config {
	serverEnv := os.Getenv("SERVER")
	zabbixPortEnv := os.Getenv("ZABBIX_PORT")
	httpPortEnv := os.Getenv("HTTP_PORT")
	serverFlag := flag.String("server", "127.0.0.1", "set zabbix server address, default 127.0.0.1")
	zabbixPortFlag := flag.Int("zabbix-port", 10051, "set zabbix server port, default 10051")
	httpPortFlag := flag.Int("http-port", 8080, "http server port, default 8080")
	flag.Parse()
	config := &Config{serverAddr: getPriorityFilled(serverEnv, *serverFlag),
		zabbixPort: getPriorityFilled(zabbixPortEnv, strconv.Itoa(*zabbixPortFlag)),
		httpPort:   getPriorityFilled(httpPortEnv, strconv.Itoa(*httpPortFlag))}
	announceConfig(config)
	return config
}
func announceConfig(cfg *Config) {
	colorize(ColorYellow, fmt.Sprintf("zabbix server addr: %v:%v", cfg.serverAddr, cfg.zabbixPort))
	colorize(ColorYellow, fmt.Sprintf("http server port: %v", cfg.httpPort))
}

func getPriorityFilled(one, two string) string {
	if len(one) == 0 {
		return two
	} else {
		return one
	}
}
