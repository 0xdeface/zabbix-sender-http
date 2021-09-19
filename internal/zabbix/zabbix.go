package zabbix

import (
	"encoding/binary"
	"encoding/json"
	"log"
	"sync"
	"time"
)

var header = []byte("ZBXD\x01")

const HeaderLength = 5
const DataLength = 8

type Message struct {
	Host  string `json:"host"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Packet struct {
	mu      sync.Mutex
	Request string    `json:"request"`
	Data    []Message `json:"data"`
	Clock   int64     `json:"clock"`
}

func CreateMessage(host, key, value string) Message {
	return Message{
		Host:  host,
		Key:   key,
		Value: value,
	}
}
func NewPacket(messages []Message, clock ...int64) *Packet {
	zp := &Packet{Request: `sender data`, Data: messages}
	if zp.Clock = time.Now().Unix(); len(clock) > 0 {
		zp.Clock = clock[0]
	}
	return zp
}

func NewPacketWithMessage(host, key, value string, clock ...int64) *Packet {
	zp := &Packet{Request: `sender data`, Data: []Message{CreateMessage(host, key, value)}}
	if zp.Clock = time.Now().Unix(); len(clock) > 0 {
		zp.Clock = clock[0]
	}
	return zp
}

func (zp *Packet) AddMessage(host, key, value string) {
	zp.mu.Lock()
	zp.Data = append(zp.Data, Message{
		Host:  host,
		Key:   key,
		Value: value,
	})
	zp.mu.Unlock()
}
func (zp *Packet) Prepare() []byte {
	serialized, _ := json.Marshal(zp)
	dataLen := make([]byte, DataLength)
	binary.LittleEndian.PutUint32(dataLen, uint32(len(serialized)))
	buf := append(header, dataLen...)
	buf = append(buf, serialized...)
	zp.mu.Lock()
	zp.Data = make([]Message, 0)
	zp.mu.Unlock()
	return buf
}
func (zp *Packet) ResponseDecode(response []byte) []byte {
	if len(response) <= HeaderLength+DataLength {
		log.Printf("%v", response)
		return response
	}
	return response[HeaderLength+DataLength:]
}
