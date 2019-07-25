package main

import (
	"sync"
	"sync/atomic"
)

type singleton struct {

}
var (
	instance *singleton
	initialized uint32
	mu sync.Mutex
)

func Instance() *singleton  {
	if atomic.LoadUint32(&initialized) == 1{
		return instance
	}
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized,1)
		instance = &singleton{}
	}
	return instance
}
// ---------------------------------------------------------------------\
//					Once 模拟实现
type Once struct {
	m sync.Mutex
	done uint32
}

func (o *Once)Do(f func())  {
	if atomic.LoadUint32(&o.done) == 1{
		return
	}

	o.m.Lock()
	defer o.m.Unlock()

	if o.done == 0 {
		defer atomic.StoreUint32(&o.done , 1)
		f()
	}
}
// ---------------------------------------------------------------------/
// ---------------------------------------------------------------------\
// 					# 基于Once 实现单例
var (
	instance2 *singleton
	once Once
)

func Instance() *singleton  {
	once.Do(func() {
		instance2 = &singleton{}
	})
	return instance2
}
/**
sync/atomic 包对基本的数值类型及复杂对象的读写都提供了原子操作的支
持。 atomic.Value 原子对象提供了 Load 和 Store 两个原子方法，分别用于加
载和保存数据，返回值和参数都是 interface{} 类型，因此可以用于任意的自定
义复杂类型。
 */
// ---------------------------------------------------------------------/
func main() {
	
}
