package main

import (
	"fmt"
	"github.com/micro/go-micro/config"
	"log"
)


func main() {
	//conf := config.NewConfig()
	//println(conf)
	println("loading config file")
	err := config.LoadFile("./conf/config.json")
	if err != nil {
		log.Fatalln(err)
	}
	println("to map")
	conf := config.Map()
	fmt.Println(conf)
	hostsDbPort := config.Get("hosts","database","port").Int(80)
	fmt.Println("db port :", hostsDbPort)
}
