package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//basic3()
	basic4()
	demo5()
}
func basic() {
	a, b := 1, 0
	ans := a / b
	fmt.Println(ans)
}
func basic2() {
	fmt.Println("start")
	panic("something bad happened")
	fmt.Println("end")
}
func basic3() {
	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("something bad happened")
	fmt.Println("end")
}
func basic4() {
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			// fmt.Println("Error: ", err)
			log.Println("Error: ", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("end")
}
func demo5() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}
func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			// 可以重抛: panic(err)
		}
	}()
	panic("something bad happend")
	fmt.Println("done panicking")
}
func basicHttpServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}
