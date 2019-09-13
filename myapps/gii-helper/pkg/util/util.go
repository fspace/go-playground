package util

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus" // replace the std log package
)

func PrettyPrint(v interface{}) {
	//   fmt.Printf("%#v", p) //with name, value and type
	// b, _ := json.MarshalIndent(v, "", "  ")
	// println(string(b))
	spew.Dump(v)
}
func PrintJson(v interface{}) {
	//   fmt.Printf("%#v", p) //with name, value and type
	b, _ := json.MarshalIndent(v, "", "  ")
	//fmt.Println(b)
	fmt.Println(string(b))
}

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
