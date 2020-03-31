package ziface

type IRequest interface {
	// 得到当前连接
	GetConnection() IConnection

	// 得到当前数据
	GetData() []byte
}
