package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"os"
	"time"
)

func main() {
	os.Exit(realMain())
}
func realMain() (exitCode int) {
	app := cli.App("timer-demo", "demo for timer")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	app.Command("bs", "basic syntax", cli.ActionCommand(basics))
	app.Command("ticker", "basic syntax", cli.ActionCommand(ActionTicker))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================

// ==============================================================================
func basics() {
	c := make(chan int)

	//使用time.AfterFunc：
	d := time.Second * 6
	f := func() {
		fmt.Println("hi this is function called by a timer!")
	}

	t := time.AfterFunc(d, f) //time.AfterFunc这种方式创建的Timer，在到达超时时间后会在单独的goroutine里执行函数f
	_ = t                     // t 基本没啥用了

	//使用time.After：
	select {
	case m := <-c:
		handle(m)
	case <-time.After(1 * time.Minute):
		fmt.Println("timed out after 1 minute")
	}

	// 使用time.NewTimer:
	t2 := time.NewTimer(5 * time.Minute)
	select {
	case m := <-c:
		handle(m)
	case <-t2.C:
		fmt.Println("after 4 minute : timed out")
	}
}

func handle(obj int) {
	fmt.Println(obj)
}

func ActionTicker() {
	// 使用time.Tick:
	go func() {
		// time.Tick底层的Ticker不能被垃圾收集器恢复；
		// 所以使用time.Tick时一定要小心，为避免意外尽量使用time.NewTicker返回的Ticker替代。
		for t := range time.Tick(time.Minute) {
			fmt.Println("Tick at", t)
		}
	}()

	// 使用time.Ticker
	var ticker *time.Ticker = time.NewTicker(1 * time.Second)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Second * 5)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

func ActionReset() {
	/**
	重置计时器时必须注意不要与当前计时器到期发送时间到t.C的操作产生竞争。如果程序已经从t.C接收到值，则计时器是已知的已过期，
	并且t.Reset可以直接使用。如果程序尚未从t.C接收值，计时器必须先被停止，并且-如果使用t.Stop时报告计时器已过期，那么请排空其通道中值。

	例如：

	if !t.Stop() {
	  <-t.C
	}
	t.Reset(d)
	*/
	c := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 1)
			c <- false
		}

		time.Sleep(time.Second * 1)
		c <- true
	}()

	go func() {
		// try to read from channel, block at most 5s.
		// if timeout, print time event and go on loop.
		// if read a message which is not the type we want(we want true, not false),
		// retry to read.
		timer := time.NewTimer(time.Second * 5)
		for {
			// timer is active , not fired, stop always returns true, no problems occurs.
			if !timer.Stop() {
				<-timer.C
			}
			timer.Reset(time.Second * 5)
			select {
			case b := <-c:
				if b == false {
					fmt.Println(time.Now(), ":recv false. continue")
					continue
				}
				//we want true, not false
				fmt.Println(time.Now(), ":recv true. return")
				return
			case <-timer.C:
				fmt.Println(time.Now(), ":timer expired")
				continue
			}
		}
	}()

	//to avoid that all goroutine blocks.
	var s string
	fmt.Scanln(&s)
}
