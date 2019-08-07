package main

const MAX_Xx = 100                    // integer
const MAX_Xx2 = 20.99                 // float
const EVENT_USER_ADDED = "user-added" // string

const STATUS_ON = true // boolean
const STATUS_OFF = false

// const RGBS  = [...]int{255} // wrong！

const PiApprox = 3.14

var PiInt int = PiApprox  // 3, converted to integer
var Pi float64 = PiApprox // is a float

type MyString string

const Greeting = "Hello!"

var s1 string = Greeting   // is a string
var s2 MyString = Greeting // string is converted to MyString

func main() {

}
