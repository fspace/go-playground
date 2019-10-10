package main

import (
	"fmt"
	"github.com/go-ozzo/ozzo-dbx"
)

func main() {
	//db, _ := dbx.Open("mysql", "yiqing:yiqing@/yii_space")
	//
	//lk := dbx.Like("name", "Charles")
	fmt.Println(dbx.Like("name", "Charles"))

}
