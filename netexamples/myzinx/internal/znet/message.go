package znet

import "playgo/netexamples/myzinx/internal/ziface"

type Message struct {
	Id uint32
	// DataLen unit32        坑爹的相似性
	DataLen uint32
	Data    []byte
}

func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

func (m *Message) GetMsgId() uint32 {
	return m.Id
}

func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMsgId(id uint32) {
	m.Id = id
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}

func (m *Message) SetDataLen(dataLen uint32) {
	m.DataLen = dataLen
}

var _ ziface.IMessage = &Message{}
