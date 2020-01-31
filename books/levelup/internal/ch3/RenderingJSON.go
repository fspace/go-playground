package ch3

import (
	"encoding/json"
	"fmt"
)

func JsonMarshal() {
	article := Article{
		Name:       "JSON in Go",
		AuthorName: "Mal Curtis",
		Draft:      true,
	}
	//data, err := json.Marshal(article)
	data, err := json.MarshalIndent(article, "", " ")
	/**
		The
	most common JSON indentation is two spaces, but you could use four, or a tab, or
	the emoji symbol for a cat—but other software might not be able to read it
	*/
	if err != nil {
		fmt.Println("Couldn’t marshal article:", err)
	} else {
		fmt.Println(string(data))
	}
}

// Product2
type Product2 struct {

	// Field appears in JSON with the key "name".
	Name string `json:"name"`
	// Field appears in the JSON with the key "author_name",
	// but doesn’t appear at all if its value is empty.
	AuthorName string `json:"author_name,omitempty"`
	// Field will not appear in the JSON representation.
	CommissionPrice float64 `json:"-"`
}

func CustomJSONKeys() {
	p := Product2{
		Name:            "desk",
		AuthorName:      "yiqing",
		CommissionPrice: 22.2,
	}

	data, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		fmt.Printf("Couldn’t marshal %T: %v", p, err)
	} else {
		fmt.Println(string(data))
	}
}

type Article2 struct {
	Name string
}
type ArticleCollection struct {
	Articles []Article2 `json:"articles"`
	Total    int        `json:"total"`
}

func NestedTypes() {
	p1 := Article2{Name: "JSON in Go"}
	p2 := Article2{Name: "Marshaling is easy"}
	articles := []Article2{p1, p2}
	collection := ArticleCollection{
		Articles: articles,
		Total:    len(articles),
	}
	data, err := json.MarshalIndent(collection, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func UnknownJSONStructure() {
	FooJSON(`{
"foo": 123
}`)
	FooJSON(`{
"foo": "bar"
}`)
	FooJSON(`{
"foo": []
}`)
}
func FooJSON(input string) {
	data := map[string]interface{}{}
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		panic(err)
	}
	foo, _ := data["foo"]
	switch foo.(type) {
	case float64:
		/**
				Note how the JSON input with the integer ended up with type float64. This is because
		JSON doesn’t differentiate between integers and floats, it just has the type
		“number”. Since Go is unable to infer the expected data type, it defaults to float64.
		*/
		fmt.Printf("Float %f\n", foo)
	case string:
		fmt.Printf("String %s\n", foo)
	default:
		fmt.Printf("Something else\n")
	}
}

/**
-  https://github.com/fatih/gomodifytags  利用工具给struct添加tag 会方便很多哦！
*/
