package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"os"
	"sync"
	"time"
)

func main() {
	os.Exit(realMain())
}
func realMain() (exitCode int) {
	app := cli.App("fuction-demo", "demo for function type")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	app.Command("bs", "basic syntax", cli.ActionCommand(basics))
	app.Command("bs2", "basic syntax: ", cli.ActionCommand(basics2))
	app.Command("bs3", "basic syntax: 同时充当读者|写者", cli.ActionCommand(basics3))

	app.Command("bc", "buffered channels: 缓冲型channels", cli.ActionCommand(bufferedChannels))
	app.Command("fr", "forRangeLoopsWithChannels: 缓冲型channels: 从channel中读不确定个数的元素", cli.ActionCommand(forRangeLoopsWithChannels))
	app.Command("cc", "closing channels : comma Ok ", cli.ActionCommand(closingChannels2))

	app.Command("ss", "select Statement :   ", cli.ActionCommand(selectStatement))
	app.Command("ss2", "select Statement2 : 优雅关闭channel|使用select语句读不同的channel   ", cli.ActionCommand(selectStatement2))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================
// ## AGENDA

var wg = sync.WaitGroup{}

func basics() {
	ch := make(chan int)

	wg.Add(2)

	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	go func() {
		ch <- 42
		wg.Done()
	}()

	wg.Wait()
}
func basics2() {
	ch := make(chan int)
	for j := 0; j < 5; j++ {
		wg.Add(2)
		go func() {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()

		go func() {
			ch <- 42
			wg.Done()
		}()
	}
	wg.Wait()
}
func basics3() {
	ch := make(chan int)
	wg.Add(2)

	go func() {
		i := <-ch
		fmt.Println(i)
		ch <- 27
		wg.Done()
	}()

	go func() {
		ch <- 42
		fmt.Println(<-ch)
		wg.Done()
	}()

	wg.Wait()
}

// 窄化channel方向
func restrictingDataFlow() {
	ch := make(chan int)
	wg.Add(2)
	// 只读型channel
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	// 只写
	go func(ch chan<- int) {
		ch <- 42
		wg.Done()
	}(ch)

	wg.Wait()
}

func bufferedChannels() {
	ch := make(chan int, 50)
	wg.Add(2)
	// 只读型channel
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)

		i = <-ch
		fmt.Println(i)

		wg.Done()
	}(ch)

	// 只写
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		wg.Done()
	}(ch)

	wg.Wait()
}
func closingChannels2() {
	ch := make(chan int, 50)
	wg.Add(2)
	// 只读型channel
	go func(ch <-chan int) {

		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}

		}

		wg.Done()
	}(ch)

	// 只写
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		// NOTE 这里的关闭很重要 不然读的协程就死锁了
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}
func forRangeLoopsWithChannels() {
	ch := make(chan int, 50)
	wg.Add(2)
	// 只读型channel
	go func(ch <-chan int) {

		for i := range ch {
			fmt.Println(i)
		}

		wg.Done()
	}(ch)

	// 只写
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		// NOTE 这里的关闭很重要 不然读的协程就死锁了
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}

// --------------------------------------------
const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)

var doneCh = make(chan struct{})

func logger() {
	for entry := range logCh {
		fmt.Printf("%v - [%v] %v\n ", entry.time.Format("2006-01-02T15:04:05"),
			entry.severity, entry.message)
	}
}
func logger2() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v] %v\n ", entry.time.Format("2006-01-02T15:04:05"),
				entry.severity, entry.message)
		case <-doneCh:
			break
		}

	}
}
func selectStatement() {
	go logger()

	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
}
func selectStatement2() {
	go logger2()

	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{}
}

// --------------------------------------------
