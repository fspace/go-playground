package ziface

type IServer interface {
	// 启动
	Start()
	// 停止
	Stop()
	// 运行
	Serve()

	// 路由功能
	AddRouter(msgID uint32, router IRouter)
	// 获取当前服务器的连接管理器
	GetConnMgr() IConnManager

	// 事件支持
	SetOnConnStart(func(connection IConnection))
	SetOnConnStop(func(connection IConnection))
	CallOnConnStart(connection IConnection)
	CallOnConnStop(connection IConnection)
}
