package main

import (
	"fmt"
	"strconv"
	"time"
)

/**
## 可以看下 php2go 库中的玩法
*/

func time1() {
	now := time.Now()
	fmt.Printf("now is : %v \n type is : %T \n", now, now)
	// 获取年月日时分秒
	fmt.Printf("年: %v\t月: %v\t日: %v\t时: %v\t分: %v\t秒: %v\t \n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("微秒：%v \n", now.Nanosecond())
	// 月份名强制转换 底层类型都是int型 ：// A Month specifies a month of the year (January = 1, ...).
	//type Month int
	fmt.Println("月：", int(now.Month()))
}

// 格式化输出
func time2() {
	now := time.Now()
	// 获取年月日时分秒
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("%d/%d/%d %d:%d:%d\n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	dateStr := fmt.Sprintf("%d/%d/%d %d:%d:%d\n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println("data str is :", dateStr)
}

const (
	FMT_STR     = "2006/01/02 15:04:05"
	TIME_FMT_Y  = "2006"
	TIME_FMT_M  = "01"
	TIME_FMT_D  = "02"
	TIME_FMT_H  = "15"
	TIME_FMT_I  = "04"
	TIME_FMT_S  = "05"
	TIME_FMT_NS = "-0700" // 什么东东 这个啥呀！好像不是纳秒
)

func time3() {
	now := time.Now()
	fmt.Println(now.Format(FMT_STR))
	fmt.Println("YEAR:  ", now.Format(TIME_FMT_Y))
	fmt.Println("Month:  ", now.Format(TIME_FMT_M))
	fmt.Println("Nanosecond:  ", now.Format(TIME_FMT_NS))

}

// 时间常量本质上是int64 +|-|* 是没问题的 但不可用除 除就变类型了
func constDemo() {
	fmt.Println("1 Hour:", 60*time.Minute)

	var i = 0
	for {
		i++
		fmt.Println(i)
		time.Sleep(time.Second)
		if i == 10 {
			break
		}
	}
}

// most commonly used ;  most used; most frequently used
func mostUsed1() {
	// Unix 时间戳  1970年1月1日 至今经历过的秒数 纳秒数
	now := time.Now()
	curr := time.Now().Unix()
	fmt.Println(curr)

	fmt.Println("as random number: ", now.UnixNano())
}

// in practice 实战：
// - 统计某个函数运行的时间
func inAction1() {
	testFunc := func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(i)
		}
	}
	_ = testFunc

	testFunc2 := func() {
		str := ""
		for i := 0; i < 1000; i++ {
			str += "hello " + strconv.Itoa(i)
		}
	}
	start := time.Now().UnixNano()
	//testFunc()
	testFunc2()
	end := time.Now().UnixNano()

	span := end - start
	fmt.Println("耗时：", span, "纳秒")

}

func main() {
	//time1()
	//time2()
	//time3()
	//
	//constDemo()
	//
	//mostUsed1()
	inAction1()
}
