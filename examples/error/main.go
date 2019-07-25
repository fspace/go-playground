package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"syscall"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal("top level recover:", r)
		}
	}()

	m := map[string]string{
		"key1": "value1",
		"key":  "value",
	}
	if v, ok := m["key"]; ok {
		fmt.Println("key exists:", "key", "=>", v)
	}

	// 测试下异常回复
	err := SomeAPI()
	if err != nil {
		fmt.Println("some error occurred when invoke the api: ",err)
	}
	err = syscall.Chmod(":invalid path", 0666)
	if err != nil {
		log.Fatalln(err.(syscall.Errno))
	}
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
// SomeAPI 导出的API方法
// Go语言库的实现习惯: 即使在包内部使用了 panic ，但是在导出函数时会被转化为明确的错误值。
func SomeAPI()  (err error){
	// panic 转换为错误返回
	/**
	必须要和有异常的栈帧只隔一个栈帧， recover 函数才能正常捕获异常。换言
	之， recover 函数捕获的是祖父一级调用函数栈帧的异常（刚好可以跨越一
	层 defer 函数）！
	 */
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("API: internal error: %v", p)
		}
	}()
	// _ = 1/0
	panic("hi i am test !")
	return
}