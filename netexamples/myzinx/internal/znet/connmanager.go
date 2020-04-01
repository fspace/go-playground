package znet

import (
	"errors"
	"fmt"
	"playgo/netexamples/myzinx/internal/ziface"
	"sync"
)

type ConnManager struct {
	// 管理的连接集合
	connections map[uint32]ziface.IConnection
	// 保护连接集合的互斥锁
	connLock sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}

func (cm *ConnManager) Add(conn ziface.IConnection) {
	// 加写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	cm.connections[conn.GetConnID()] = conn
	fmt.Println("connID= ", conn.GetConnID(), "Connection add to ConnManager : conn num = ", cm.Len())
}

func (cm *ConnManager) Remove(conn ziface.IConnection) {
	// 加写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	delete(cm.connections, conn.GetConnID())
	fmt.Println("connID= ", conn.GetConnID(), "Connection  remove from  connManager successfully : conn num = ", cm.Len())
}

func (cm *ConnManager) Get(connID uint32) (ziface.IConnection, error) {
	// 加读锁
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()

	if conn, ok := cm.connections[connID]; ok {
		// 找到了
		return conn, nil
	}
	return nil, errors.New("connection Not FOUND")
}

func (cm *ConnManager) Len() int {
	// 不加锁？
	return len(cm.connections)
}

func (cm *ConnManager) ClearConn() {
	// 加写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	for connId, conn := range cm.connections {
		// 停止
		conn.Stop()
		// 删除
		delete(cm.connections, connId)
	}
	fmt.Println("Clear All connections succ! conn num = ", cm.Len())
}

var _ ziface.IConnManager = &ConnManager{}
