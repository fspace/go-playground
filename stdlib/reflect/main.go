package main

import (
	"fmt"
	"log"
	"playgo/pkg/models"
	"reflect"
	"strings"
)

type Person struct {
	Name string  `key:"value";json:"name"`
	Age int `json:"age"`
}

func (p Person)Eat()  {
	fmt.Println("hi i am eating !")
}

func (p Person)EatAnything(i ...interface{})  {
	/*
	var params  []interface{}
	params = append(params,"hi i am eating")
	params = append(params,i...)
	fmt.Println(params)
	*/
	fmt.Println(i...)
}

var foo = func(){
	fmt.Println("foo is called!")
	}

func reflectMe(va interface{}){
	fmt.Println("_>  enter reflectMe")
	rTpy := reflect.TypeOf(va)
	//fmt.Println(rTpy.Name())
	fmt.Println("rTpy.PkgPath is: ",rTpy.PkgPath())

	fmt.Println("TYPE ： ", rTpy.Name())
	fmt.Println("KIND ： ", rTpy.Kind())

	methodNum := rTpy.NumMethod()
	fmt.Println("mum of method is : ",methodNum)
   for i:= 0 ; i<methodNum ; i++ {
   	  fmt.Printf("the %d method of %s: %s \n ",i, rTpy.Name() , rTpy.Method(i).Name)
   }


   // ===========================================================================  +|

	rVal := reflect.ValueOf(va)
	fmt.Println("rVal is : ", rVal)

	// 基本类型 也可以直接读
	if rTpy.Kind() == reflect.Int {
		fmt.Println("is int and the value is :", rVal.Int())
	}

	va2 := rVal.Interface()
	switch va2.(type) {
	case int:
		fmt.Println("i am int :",va2)
	case Person:
		fmt.Println("I am a Person and my name is :",va2.(Person).Name)
	}

	fmt.Println(strings.Repeat("=",70),"end reflectMe")
}


func modify(x interface{}){
	rTpy := reflect.TypeOf(x)
	rVal := reflect.ValueOf(x)
	if rTpy.Kind() == reflect.Ptr {
		 fmt.Println("ya it is a ptr!")
		 v := rVal.Elem()
		 if v.Kind() == reflect.Int {
			 v.SetInt(3)
		 }

	}
}

func call(i interface{}){
	rType := reflect.TypeOf(i)
	rValue := reflect.ValueOf(i)

	if rType.Kind() != reflect.Func {
		fmt.Println("expect a function")
		return
	}
	// 入参 出参判断
	fmt.Printf("in num is %d and the num of out is %d  \n ", rType.NumIn(), rType.NumOut())
	// Call 方法不是啥都能调用的  Value 是一个泛设计 包括所有类型情形 所以调用某些特定方法前 最好做下类型判断 或者断言
	rtnVal := rValue.Call([]reflect.Value{})
	fmt.Println("return value is :", rtnVal)
}

func  inspectStruct(i interface{})  {
	rTyp := reflect.TypeOf(i)
	rVal := reflect.ValueOf(i)

	if rTyp.Kind() != reflect.Struct {
		log.Fatalln("need a struct ")
	}

	numField := rTyp.NumField()
    if numField > 0 {
    	for i:= 0 ; i< numField ; i++ {
    		fmt.Printf("the ith( %d ) filed is %s \n", i , rTyp.Field(i).Name)
    		fmt.Printf("  filed  tag is %v \n",  rTyp.Field(i).Tag)
    		fmt.Printf("  filed  json tag is %v \n",  rTyp.Field(i).Tag.Get("json"))
		}
	}

	// 消耗掉变量 占位用
	_ = rVal

}
func  callStructMethod(i interface{})  {
	rTyp := reflect.TypeOf(i)
	rVal := reflect.ValueOf(i)

	if rTyp.Kind() != reflect.Struct {
		log.Fatalln("need a struct ")
	}
	// 消耗掉变量 占位用
	//_ = rVal
	if _ ,ok := rVal.Interface().(Person); !ok {
		log.Fatalln("need a Person ")
	}
   // NOTE 按索引获取方法时  是ascii递增的 不是按照定义函数出现的顺序来的 ！！！
	m,_ := rTyp.MethodByName("EatAnything")
	var inParams []reflect.Value
	inParams = append(inParams, reflect.ValueOf("yes i am a string"))

	rtn := m.Func.Call([]reflect.Value{
		 rVal, // NOTE  第一个是函数的所有者！！！
		reflect.ValueOf("yes i am a string"),
		reflect.ValueOf(true),
		reflect.ValueOf(1),
	})

	 // rtn := m.Func.Call(inParams)
	fmt.Println("the function result is ", rtn)
}

func main() {
	i := 2
	reflectMe(i)

	p := Person{
		Name: "yiqing",
		Age: 18,
	}
	reflectMe(p)

	reflectMe(foo)

	p2 := models.Person{}
	reflectMe(p2)
	// 修改
	i2 := 3
	modify(&i2)
	fmt.Println("modified value is：" , i2)

	call(foo)

	// 结构体
	inspectStruct(p)
	callStructMethod(p)
}
