package main

/**
https://learnku.com/articles/39255
*/
import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/jawher/mow.cli"
	"github.com/olekukonko/tablewriter"
	"os"
	"reflect"
	"strconv"
	"unsafe"
)

func main() {
	os.Exit(realMain())
}

func realMain() (exitCode int) {
	app := cli.App("fuction-demo", "demo for function type")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	app.Command("bs", "basic syntax", cli.ActionCommand(basics))
	app.Command("si", "struct Info", cli.ActionCommand(structInfo))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ========================================= =======================================================
//创建两个个结构体
type Person1 struct {
	a bool
	b int64
	c int8
	d string
}
type Person2 struct {
	b int64
	c int8
	a bool
	d string
}

// ========================================= =======================================================
func basics() {
	//创建一个变量
	var i int8 = 10

	//建一个变量转化成Pointer 和 uintptr
	p := unsafe.Pointer(&i) //入参必须是指针类型的
	fmt.Println(p)          //是内存地址0xc0000182da
	u := uintptr(i)
	fmt.Println(u) //结果就是10

	//Pointer转换成uintptr
	//temp := uintptr(p)
	//uintptr转Pointer
	p = unsafe.Pointer(u)

	//获取指针大小
	u = unsafe.Sizeof(p) //传入指针，获取的是指针的大小
	fmt.Println(u)       // 打印u是：8
	//获取的是变量的大小
	u = unsafe.Sizeof(i)
	fmt.Println(u) //打印u是：1

	//接下来演示一下内存对齐,猜一猜下面l两个打印值是多少呢?
	person1 := Person1{a: true, b: 1, c: 1, d: "spw"}
	fmt.Println(unsafe.Sizeof(person1))
	person2 := Person2{b: 1, c: 1, a: true, d: "spw"}
	fmt.Println(unsafe.Sizeof(person2))
}

func structInfo() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Type", "IndexOfField", "Size"})
	/*
		for _, v := range data {
			table.Append(v)
		}
	*/
	person1 := Person1{a: true, b: 1, c: 1, d: "spw"}
	s := structs.New(person1)

	//table.Append(
	//	[]string{"B", "The Very very Bad Man", "288"})
	data := [][]string{}
	for idx, f := range s.Fields() {
		data = append(data, []string{
			f.Name(),
			f.Kind().String(),
			strconv.Itoa(idx),
			//strconv.Itoa(unsafe.Sizeof(f.Kind().))	,
			fmt.Sprint(unsafe.Sizeof(unsafe.Pointer(reflect.TypeOf(f.Value()).Size()))),
		})
	}

	table.AppendBulk(data)

	table.Render() // Send output
}
