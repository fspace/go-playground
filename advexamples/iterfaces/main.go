package main

import (
	"bytes"
	"fmt"
	"github.com/jawher/mow.cli"
	"github.com/kr/pretty"
	"io"
	"os"
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
	app.Command("tc", "type conversion", cli.ActionCommand(typeConversions))
	app.Command("ei", "empty interfaces", cli.ActionCommand(emptyInterface))
	app.Command("ts", "type switches", cli.ActionCommand(typeSwitches))

	app.Command("eip", "empty Interface Param", cli.ActionCommand(emptyInterfaceParam))
	app.Command("di", "defining Interfaces", cli.ActionCommand(definingInterfaces))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================
type Writer interface {
	Write([]byte) (int, error)
}
type Closer interface {
	Close() error
}
type WriterCloser interface {
	Writer
	Closer
}
type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Close() error {
	//panic("implement me")
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}

	}
	return nil
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	//panic("implement me")
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

//var _ Writer = &BufferedWriterCloser{}
//var _ Closer = &BufferedWriterCloser{}

type ConsoleWriter struct {
}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type Incrementer interface {
	Increment() int
}
type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

// ----------------------------------------------
// AGENDA

// Basics
func basics() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("hello Go!"))

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i <= 10; i++ {
		fmt.Println(inc.Increment())
	}

	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello Youtube listener!, this is a test!"))
	wc.Close()
}

func composingInterfaces() {
	emptyInterface()
	typeSwitches()
}
func emptyInterface() {
	var myObj interface{} = NewBufferedWriterCloser()
	if wc, ok := myObj.(WriterCloser); ok {
		wc.Write([]byte("hello youtube ussers !"))
		wc.Close()
	}
	r, ok := myObj.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("conversion failed!")
	}
}
func typeSwitches() {
	var i interface{} = 0
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("I don't know what i is ")

	}
}
func typeConversions() {
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("hi yes this is a test data!"))
	wc.Close()

	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed!")
	}
}

func implWithValues() {

}
func implWithPointers() {

}

func bestPractices() {

}

// ---------------

func myFunc(a interface{}) {
	fmt.Println(a)
}

func emptyInterfaceParam() {
	/**\
	By defining a function that takes in an interface{}, we essentially give ourselves the flexibility to pass in
	anything we want. It’s a Go programmers way of saying, this function takes in something, but I don’t necessarily
	care about its type.
	*/
	var my_age int
	my_age = 25

	myFunc(my_age)
}

// ----------------------------------------------------------------------------------

type Guitarist interface {
	// PlayGuitar prints out "Playing Guitar"
	// to the terminal
	PlayGuitar()
}

type BaseGuitarist struct {
	Name string
}

type AcousticGuitarist struct {
	Name string
}

func (b BaseGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the Bass Guitar\n", b.Name)
}

func (b AcousticGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the Acoustic Guitar\n", b.Name)
}

// ----------------------------------------------------------------------------------
func definingInterfaces() {
	var player BaseGuitarist
	player.Name = "Paul"
	player.PlayGuitar()

	var player2 AcousticGuitarist
	player2.Name = "Ringo"
	player2.PlayGuitar()

	var guitarists []Guitarist
	guitarists = append(guitarists, player)
	guitarists = append(guitarists, player2)
	pretty.Print(guitarists)
}
