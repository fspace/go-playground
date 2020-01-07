package lib

import "fmt"

type Guitarist struct {
}

type ParamsStruct struct {
}

func UpdateGuitarist(guitarist *Guitarist, params ParamsStruct) {
	fmt.Println("This is a simple function")
}

func (g *Guitarist) Update(params ParamsStruct) {
	fmt.Println("This is a simple method")
}
