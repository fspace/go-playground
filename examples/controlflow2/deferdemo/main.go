package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	basic()
	basic2()

	resourceRelease()
	resourceRelease2()

	basic3()
}

func basic() {
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
}
func basic2() {
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")
}

func basic3() {
	a := "start"
	defer fmt.Println(a)
	a = "end"
}

func resourceRelease() {
	res, err := http.Get("http://www.baidu.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
func resourceRelease2() {
	res, err := http.Get("http://www.baidu.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
