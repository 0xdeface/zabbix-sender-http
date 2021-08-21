package main

import (
	"encoding/binary"
	"encoding/json"
	"time"
)

var header = []byte("ZBXD\x01")

type ZabbixPreparedPackage []byte

type ZabbixMessage struct {
	Host  string `json:"host"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ZabbixPacket struct {
	Request string          `json:"request"`
	Data    []ZabbixMessage `json:"data"`
	Clock   int64           `json:"clock"`
}

type ZabbixResponse struct {
	Response string
	Info     string
}

func Message(host, key, value string) ZabbixMessage {
	return ZabbixMessage{
		Host:  host,
		Key:   key,
		Value: value,
	}
}
func NewZabbixPacket(messages []ZabbixMessage, clock ...int64) *ZabbixPacket {
	zp := &ZabbixPacket{Request: `sender data`, Data: messages}
	if zp.Clock = time.Now().Unix(); len(clock) > 0 {
		zp.Clock = clock[0]
	}
	return zp
}
func (zp *ZabbixPacket) AddMessage(host, key, value string) {
	zp.Data = append(zp.Data, ZabbixMessage{
		Host:  host,
		Key:   key,
		Value: value,
	})
}
func (zp *ZabbixPacket) Prepare() ZabbixPreparedPackage {
	serialized, _ := json.Marshal(zp)
	dataLen := make([]byte, 8)
	binary.LittleEndian.PutUint32(dataLen, uint32(len(serialized)))
	buf := append(header, dataLen...)
	buf = append(buf, serialized...)
	return buf
}
