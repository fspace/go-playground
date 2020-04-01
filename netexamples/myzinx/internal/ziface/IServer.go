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
}
