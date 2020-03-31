package znet

import "playgo/netexamples/myzinx/internal/ziface"

type BaseRouter struct {
}

func (br *BaseRouter) PreHandle(request ziface.IRequest) {
	//panic("implement me")
}

func (br *BaseRouter) Handle(request ziface.IRequest) {
	//panic("implement me")
}

func (br *BaseRouter) PostHandle(request ziface.IRequest) {
	//panic("implement me")
}

var _ ziface.IRouter = &BaseRouter{}
