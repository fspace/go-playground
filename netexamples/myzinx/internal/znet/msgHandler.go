package znet

import (
	"fmt"
	"playgo/netexamples/myzinx/internal/ziface"
	"playgo/netexamples/myzinx/utils"
	"strconv"
)

/**
消息处理模块的实现
*/
type MsgHandle struct {
	// 存放每个MsgID 对应的处理
	Apis map[uint32]ziface.IRouter
	// 负责Worker取任务的消息队列
	TaskQueue []chan ziface.IRequest
	// 业务工作worker池的worker数量
	WorkerPooSize uint32
}

//
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:          make(map[uint32]ziface.IRouter),
		WorkerPooSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:     make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
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

// 启动Worker工作池 （只需要一次)
func (mh *MsgHandle) StartWorkerPool() {
	for i := 0; i < int(mh.WorkerPooSize); i++ {
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		// 启动当前worker 阻塞等待channel中的消息
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}
func (mh *MsgHandle) StartOneWorker(workerId int, taskQueue chan ziface.IRequest) {
	fmt.Println("Worker ID=", workerId, "is started ...")
	for {
		select {
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

// 将消息交给TaskQueue
func (mh *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	// 根据连接id来分配 将消息平均分配给worker
	workerId := request.GetConnection().GetConnID() % mh.WorkerPooSize
	fmt.Println("Add ConnID: ", request.GetConnection().GetConnID(),
		"request MsgID: ", request.GetMsgID(), "to WorkerID: ", workerId)

	// 将消息发送给对应的队列
	mh.TaskQueue[workerId] <- request
}

var _ ziface.IMsgHandle = &MsgHandle{}
