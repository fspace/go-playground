package ziface

type IRequest interface {
	// 得到当前连接
	GetConnection() IConnection

	// 得到当前数据
	GetData() []byte

	// 得到当前请求消息的ID
	GetMsgID() uint32
}
