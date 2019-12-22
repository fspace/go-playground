package main

import (
	"fmt"
	"github.com/go-ozzo/ozzo-config"
)

func main() {
	c := config.New()

	// =========================================================================================

	// ## 多种加载源
	// load from one or multiple JSON, YAML, or TOML files.
	// file formats are determined by their extensions: .json, .yaml, .yml, .toml
	// c.Load("app.json", "app.dev.json")
	c.Load("app.toml")

	// load from one or multiple JSON strings
	// c.LoadJSON([]byte(`{"Name": "abc"}`), []byte(`{"Age": 30}`))

	// load from one or multiple variables
	//data1 := struct {
	//	Name string
	//} { "abc" }
	//data2 := struct {
	//	Age int
	//} { 30 }
	//c.SetData(data1, data2)
	// ========================================================================================= />
	// ##  Accessing Configuration
	name := c.GetString("Name", "Yiqing")
	fmt.Println(name)

	pg_port := c.Get("Postgres.Port", 5431)
	fmt.Println("pg port is :", pg_port)

	// ## Changing Configuration
	c.Set("Postgres.Port", 5488)
	fmt.Println("new pg port is :", c.Get("Postgres.Port"))

	// ## Configuring Objects
	var pgConfig struct {
		Enabled           bool
		Port              int
		Hosts             []string
		AvailabilityRatio float32
	}
	// populate the author object from the "Author" configuration
	c.Configure(&pgConfig, "Postgres")
	//fmt.Printf("type: %T , value: %v ", pgConfig, pgConfig)
	fmt.Println(pgConfig)
	// ========================================================================================= />

}
