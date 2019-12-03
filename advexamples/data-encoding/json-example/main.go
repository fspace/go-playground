package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

func main() {
	Demo1()
}

type Character struct {
	Name        string `json:"name" tag:"foo"`
	Surname     string `json:"surname"`
	Job         string `json:"job,omitempty"`
	YearOfBirth int    `json:"year_of_birth,omitempty"`
}

// ........................................................................................ +|
// 					## 自己实现接口

func (c *Character) UnmarshalJSON(b []byte) error {
	type C Character
	var v C
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	*c = Character(v)
	if c.Job == "" {
		c.Job = "unknown"
	}
	return nil
}
func (c Character) MarshalJSON() ([]byte, error) {
	type C Character
	v := C(c)
	if v.Job == "" {
		v.Job = "unknown"
	}
	return json.Marshal(v)
}

// ........................................................................................ +|
func Demo1() {
	r := strings.NewReader(`{
"name":"Lavinia",
"surname":"Whateley",
"year_of_birth":1878
}`)
	d := json.NewDecoder(r)
	var c Character
	if err := d.Decode(&c); err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v", c)
}
func EncodeDemo() {
	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "\t")
	c := Character{
		Name:        "Charles Dexter",
		Surname:     "Ward",
		YearOfBirth: 1902,
	}
	if err := e.Encode(c); err != nil {
		log.Fatalln(err)
	}
}

/**
-  转换工具 https://mholt.github.io/json-to-go/
-  模式检验 https://github.com/xeipuuv/gojsonschema
*/
