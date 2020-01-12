package services

import "fmt"

type Service interface {
	SayHi()
}

type MyService struct{}

func (s MyService) SayHi() {
	fmt.Println("Hi")
}

type SecondService struct{}

func (s SecondService) SayHi() {
	fmt.Println("Hello From the 2nd Service")
}
