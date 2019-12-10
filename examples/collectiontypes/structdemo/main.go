package main

import (
	"fmt"
	"playgo/examples/collectiontypes/structdemo/lib"
	"reflect"
)
import "github.com/fatih/structs"

type (
	Doctor struct {
		number     int
		actorName  string
		companions []string
	}
)

func main() {
	//basic()

	embedDemo()

	//
	usecaseTag()
}
func basic() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"jo Grant",
		},
	}
	fmt.Println(aDoctor)

	// 读元素
	fmt.Println(aDoctor.companions[1])

	// structInfo(&aDoctor)
	structInfo(aDoctor)
	// 匿名struct
	s := struct {
		name string
	}{name: "Yiqing"}
	fmt.Println(s)
	structInfo(s)
	// 值类型
	s2 := s
	s2.name = "qing"
	structInfo(s)
	structInfo(s2)

	// 指针赋值
	s3 := &s
	s3.name = "yi"
	fmt.Println(s3)
	fmt.Println(s)
}

// 对某个目标的全方位审查 一般手段就是 1. 动态 运行时审查 一般使用反射 埋点  2. 静态 设计时分析 一般使用AST抽象语法树 依赖分析
//func structInfo(d Doctor)  {
func structInfo(d interface{}) {
	fmt.Printf("%v \n\n", d)
	// Convert a struct to a map[string]interface{}
	// => {"Name":"gopher", "ID":123456, "Enabled":true}
	m := structs.Map(d)

	// Convert the values of a struct to a []interface{}
	// => ["gopher", 123456, true]
	v := structs.Values(d)

	// Convert the names of a struct to a []string
	// (see "Names methods" for more info about fields)
	n := structs.Names(d)

	// Convert the values of a struct to a []*Field
	// (see "Field methods" for more info about fields)
	f := structs.Fields(d)

	// Return the struct name => "d"
	nm := structs.Name(d)

	// Check if any field of a struct is initialized or not.
	h := structs.HasZero(d)

	// Check if all fields of a struct is initialized or not.
	z := structs.IsZero(d)

	// Check if d is a struct or a pointer to struct
	i := structs.IsStruct(d)

	fmt.Printf("Convert a struct to a map[string]interface{}: %v, %T \n", m, m)
	fmt.Printf("Convert the values of a struct to a []interface{}: %v, %T \n", v, v)
	fmt.Printf("Convert the names of a struct to a []string: %v, %T \n", n, n)
	fmt.Printf(" Convert the values of a struct to a []*Field: %v, %T \n", f, f)
	fmt.Printf(" Return the struct name => : %v, %T \n", nm, nm)
	fmt.Printf("Check if any field of a struct is initialized or not.: %v, %T \n", h, h)
	fmt.Printf(" Check if all fields of a struct is initialized or not..: %v, %T \n", z, z)
	fmt.Printf(" Check if d is a struct or a pointer to struct: %v, %T \n", i, i)
	//

	s := structs.New(d)

	for _, f := range s.Fields() {
		fmt.Printf("field name: %+v\n", f.Name())

		if f.IsExported() {
			fmt.Printf("value   : %+v\n", f.Value())
			fmt.Printf("is zero : %+v\n", f.IsZero())
		} else {
			fmt.Printf("not exported   : %+v\n", f.Name())
		}
	}
}

func structFieldInfo(inst interface{}, field string) {
	s := structs.New(inst)

	// Get the Field struct for the "Name" field
	name := s.Field(field)

	// Get the underlying value,  value => "gopher"
	//value := name.Value().(string)
	// Set the field's value
	// name.Set("another gopher")

	// Get the field's kind, kind =>  "string"
	fmt.Printf("field :%s kind : ", name.Kind())

	// Check if the field is exported or not
	if name.IsExported() {
		fmt.Println("Name field is exported")
	}

	// Check if the value is a zero value, such as "" for string, 0 for int
	if !name.IsZero() {
		fmt.Println("Name is initialized")
	}

	// Check if the field is an anonymous (embedded) field
	if !name.IsEmbedded() {
		fmt.Println("Name is not an embedded field")
	}

	// Get the Field's tag value for tag name "json", tag value => "name,omitempty"
	// tagValue := name.Tag("json")
}

func embedDemo() {
	b := lib.Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)
}

func usecaseTag() {
	// 使用反射获取结构体的某个字段上依附的tag
	t := reflect.TypeOf(lib.Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
