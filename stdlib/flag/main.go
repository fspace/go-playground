package main

import (
	"flag"
	"fmt"
)

type Config struct {
	Addr      string
	StaticDir string
}

/**
- go run main.go -static-dir=../public/static -addr=:9000

*/
func main() {
	// flag 解析到已经存在的变量地址
	cfg := new(Config)
	flag.StringVar(&cfg.Addr, "addr", ":4000", "Http network address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "../ui/static", "static html dir")
	//
	// addr := flag.String("addr",":4000","HTTP network address")
	flag.Parse()
	// fmt.Println("the server addr is :", *addr)
	fmt.Printf("the config is %#v", cfg)

}
