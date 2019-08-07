package main

import "os"

func deferExample() error {
	f, err := os.Open("config.txt")
	if err != nil {
		return err
	}
	defer f.Close() // 无论如何总是会关闭的   defer的列表会以栈的形式执行--FILO 先进后出
	// 用f做一些操作
	return nil
}

func main() {

}
