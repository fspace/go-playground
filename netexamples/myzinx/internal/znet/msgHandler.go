package znet

import (
	"fmt"
	"playgo/netexamples/myzinx/internal/ziface"
	"strconv"
)

/**
消息处理模块的实现
*/
type MsgHandle struct {
	// 存放每个MsgID 对应的处理
	Apis map[uint32]ziface.IRouter
}

//
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

func (mh *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	//
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api MsgID=", request.GetMsgID(), " is NOT FOUND! Need register")
		panic("....")
	}
	// 调度对应的业务方法
	// TODO 可以改为链式执行呀！
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

func (mh *MsgHandle) AddRouter(msgID uint32, router ziface.IRouter) {
	//
	if _, ok := mh.Apis[msgID]; ok {
		// 对应的ID已经注册
		panic("repeat api, msg id=" + strconv.Itoa(int(msgID)))
	}
	// 添加
	mh.Apis[msgID] = router
	fmt.Println("Add api MsgID=", msgID, " succ!")
}

var _ ziface.IMsgHandle = &MsgHandle{}
