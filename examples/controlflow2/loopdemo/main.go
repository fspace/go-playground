package main

import "fmt"

func main() {
	simpleLoops()
	exitingEarly()

	simpleLoops2()
	continueDemo()
	nestedLoops()

	loopingThroughCollections()
	loopingThroughCollections2()

}

func simpleLoops() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}
}
func simpleLoops2() {
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}
	fmt.Println(i, j)
}
func infiniteLoops() {
	i := 0
	for {
		fmt.Println(i)
		i++
	}
}
func breakLoops() {
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}
}
func continueDemo() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
func nestedLoops() {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
		}
	}
}
func breakNestedLoops() {
Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				//break // 跳出最近的循环体
				break Loop // 跳出label标记的循环体 本例中即最外层的循环体
			}
		}
	}
}
func exitingEarly() {

}
func loopingThroughCollections() {
	// for range loop
	s := []int{1, 2, 3, 1000} // slice| array 都适用
	for k, v := range s {
		fmt.Println(k, v)
	}

	for i := range s {
		fmt.Println(s[i])
	}
	for idx, itm := range s {
		fmt.Printf("idx(%d) => %v \n", idx, itm)
	}
}
func loopingThroughCollections2() {
	// for range loop
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
	}
	for k, v := range statePopulations {
		fmt.Println(k, v)
	}
	for k := range statePopulations {
		fmt.Println(k)
	}
	for _, v := range statePopulations {
		fmt.Println(v)
	}

	s := "hello GO!"
	for k, v := range s {
		fmt.Println(k, v)
		fmt.Println(k, "=>", string(v))
	}

}
func loopingThroughChannels() {
	// to be continue ...

}
