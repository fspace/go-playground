package ziface

type IConnManager interface {
	// 添加
	Add(conn IConnection)
	// 删除
	Remove(conn IConnection)
	// 根据Id 获取连接
	Get(connID uint32) (IConnection, error)
	// 获取当前连接总数
	Len() int
	// 清除并终止所有的连接
	ClearConn()
}
