package ch2

import (
	"fmt"
	"github.com/micro/go-micro/util/log"
)

type Fruit interface {
	String() string
}
type Apple struct {
	Variety string
}

func (a Apple) String() string {
	return fmt.Sprintf("A %s apple", a.Variety)
}

type Orange struct {
	Size string
}

func (o Orange) String() string {
	return fmt.Sprintf("A %s orange", o.Size)
}
func PrintFruit(fruit Fruit) {
	fmt.Println("I have this fruit:", fruit.String())
}

// -------------------------------------------------------------------------
type Product struct {
}
type ProductCatalogue interface {
	All() ([]Product, error)
	Find(string) (Product, error)
}

type FileProductCatalogue struct {
	// Some fields would be here, perhaps the file location
}

func (c FileProductCatalogue) All() (ps []Product, err error) {
	// Implementation omitted
	return
}
func (c FileProductCatalogue) Find(id string) (p Product, err error) {
	// Implementation omitted
	return
}

type DBProductCatalogue struct {
	// Some fields would be here, perhaps a database connection
}

func (c DBProductCatalogue) All() (ps []Product, err error) {
	// Implementation omitted
	return
}
func (c DBProductCatalogue) Find(id string) (p Product, err error) {
	// Implementation omitted
	return
}

/**
func main() {
var myCatalogue ProductCatalogue
myCatalogue = DBProductCatalogue{}
…
}
*/
func LoadProducts(catalogue ProductCatalogue) {
	products, err := catalogue.All()
	if err != nil {
		log.Fatal(err)
		return
	}
	_ = products
}

// ================================================================================================
// ## API
func InterfacesImplements() {
	apple := Apple{"Golden Delicious"}
	orange := Orange{"large"}
	PrintFruit(apple)
	PrintFruit(orange)
}

// ================================================================================================
