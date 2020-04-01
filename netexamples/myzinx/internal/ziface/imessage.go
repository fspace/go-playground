package ziface

type IMessage interface {
	GetMsgId() uint32
	GetMsgLen() uint32
	GetData() []byte

	SetMsgId(uint32)
	SetData([]byte)
	SetDataLen(uint32)
}

// TLV(Type Len Value) 发送数据包的惯用法  主要想解决粘包问题
