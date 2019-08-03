package main

import (
	"fmt"
	_ "github.com/afex/hystrix-go/hystrix"
	"github.com/prometheus/common/log"
	"math/rand"
	"time"
)

func getInfo() (string, error) {
	r := rand.Intn(10)
	if r < 6 {
		time.Sleep(time.Second * 3)
	}

	return "some api result", nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	print("hi hystrix")

	for {
		rslt, err := getInfo()
		if err != nil {
			log.Fatalf("getInfo err: %v", err)
		}

		fmt.Println("info is : ", rslt)

		time.Sleep(time.Second * 1)
	}

}
