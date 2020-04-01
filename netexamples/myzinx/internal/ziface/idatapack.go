package ziface

// IDataPack 解决粘包问题的 封包 拆包
type IDataPack interface {
	GetHeadLen() uint32
	Pack(msg IMessage) ([]byte, error)
	Unpack([]byte) (IMessage, error)
}

/**
封包  拆包 模块
直接面向TCP连接中的数据流 ， 用于处理tcp粘包问题
*/
