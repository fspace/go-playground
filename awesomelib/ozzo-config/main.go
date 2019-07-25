package main

import (
	"fmt"
	"github.com/go-ozzo/ozzo-config"
)
/**
相似功能的库
- https://github.com/zpatrick/go-config
- https://github.com/jinzhu/configor
 */
func main() {
	// create a Config object
	c := config.New()

	// load configuration from a JSON string
	c.LoadJSON([]byte(`{
        "Version": "2.0",
        "Author": {
            "Name": "Foo",
            "Email": "bar@example.com"
        }
    }`))

	// get the "Version" value, return "1.0" if it doesn't exist in the config
	version := c.GetString("Version", "1.0")

	var author struct {
		Name, Email string
	}
	// populate the author object from the "Author" configuration
	c.Configure(&author, "Author")

	fmt.Println(version)
	fmt.Println(author.Name)
	fmt.Println(author.Email)
	// Output:
	// 2.0
	// Foo
	// bar@example.com
}