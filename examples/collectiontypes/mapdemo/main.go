package main

import (
	"fmt"
)

func main() {
	// 集合类型 可以有类似 CRUD的操作 涉及元素的增删改查
	basicDemo()

	creation()
}

func basicDemo() {
	// 无序性 key的出现顺序不代表什么
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
	}
	// 什么类型可以做key！
	m := map[[3]int]string{}
	fmt.Println(statePopulations, m)

	// manipulate
	fmt.Println(statePopulations["Ohio"], statePopulations["New York"]) // 读元素
	// 写
	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations)
	// 删除
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Georgia"]) // 已删除的可以或者不存在的key 将得到零值（k=>v 中 v的类型零值)

	// 查询存在性 comma ok 惯用法
	pop, ok := statePopulations["Ohio"]
	fmt.Println(pop, ok)
	statePopulations["Ohio"] = 11614373
	pop, ok = statePopulations["Ohio"]
	fmt.Println(pop, ok)
	if ok {
		fmt.Println("yes exists this key :", "Ohio")
	} else {
		fmt.Println("not exists this key :", "Ohio")
	}
	// 仅仅为了检测存在性 可以忽略第一个值
	k := "SomeKey"
	_, present := statePopulations[k]
	if present {
		fmt.Println("exists : ", k)
	} else {
		fmt.Println("not exists this key :", k)
	}
	// len
	fmt.Println(len(statePopulations))

	// 赋值 指针复制
	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations)
	fmt.Println(len(statePopulations))
}

func creation() {
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"Illinois": 12801539,
		"Ohio":     11614373,
	}
	fmt.Println(statePopulations)
}
