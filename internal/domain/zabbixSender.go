package domain

type ZabbixSenderPort interface {
	SendToZabbix(server, key, value string) ([]byte, error)
}

type ConnectionPort interface {
	Send([]byte) ([]byte, error)
	Close() error
}

type ZabbixPort interface {
	AddMessage(host, key, value string)
	Prepare() []byte
	ResponseDecode([]byte) []byte
}

type ZabbixSender struct {
	connection ConnectionPort
	zabbix     ZabbixPort
	logger     Logger
}

func NewZabbixSender(con ConnectionPort, zpp ZabbixPort, logger Logger) *ZabbixSender {
	return &ZabbixSender{connection: con, zabbix: zpp, logger: logger}
}

func (zs *ZabbixSender) SendToZabbix(server, key, value string) ([]byte, error) {
	(zs.zabbix).AddMessage(server, key, value)
	message := (zs.zabbix).Prepare()
	response, err := (zs.connection).Send(message)
	return zs.zabbix.ResponseDecode(response), err
}
