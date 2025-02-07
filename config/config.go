package config

import (
	"flag"
	"os"
	"strconv"
)

type Config struct {
	ZabbixHost string
	ZabbixPort string
	HttpPort   string
	Debug      bool
}

func GetConfig() *Config {
	zabbixHost := flag.String("zabbix-server", getDefault("ZABBIX_HOST", "127.0.0.1"), "set zabbix server address, default 127.0.0.1")
	zabbixPort := flag.String("zabbix-port", getDefault("ZABBIX_PORT", "10051"), "set zabbix server port, default 10051")
	httpPort := flag.String("http-port", getDefault("HTTP_PORT", "8080"), "set http server port, default 8080")
	debug := flag.Bool("debug", getDefault("DEBUG", false), "set debug mode")
	flag.Parse()
	config := &Config{
		ZabbixHost: *zabbixHost,
		ZabbixPort: *zabbixPort,
		HttpPort:   *httpPort,
		Debug:      *debug,
	}
	return config
}

// getDefault return value from environment if exists
// or return second param
func getDefault[T string | bool](envName string, defaultValue T) T {
	var (
		ok     bool
		value  string
		result any
		err    error
	)
	if value, ok = os.LookupEnv(envName); !ok {
		return defaultValue
	}
	switch any(defaultValue).(type) {
	case string:
		result = value
	case bool:
		if result, err = strconv.ParseBool(value); err != nil {
			panic(err)
		}
	}
	return result.(T)
}
