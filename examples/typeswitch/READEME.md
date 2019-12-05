https://blog.csdn.net/u012291393/article/details/79244424

~~~go

package main

import (
    "fmt"
    "reflect"
)

type Animal interface {
    shout() string
}

type Dog struct {
    name string
}

func (self Dog) shout() string {
    return fmt.Sprintf("wang wang")
}

type Cat struct {
    name string
}

func (self Cat) shout() string {
    return fmt.Sprintf("miao miao")
}

type Tiger struct {
    name string
}

func (self Tiger) shout() string {
    return fmt.Sprintf("hou hou")
}

func main() {
    // var animal Animal = Tiger{}
    // var animal Animal  // 验证 case nil
    // var animal Animal = Wolf{} // 验证 default
    var animal Animal = Dog{}

    switch a := animal.(type) {
    case nil: // a的类型是 Animal
        fmt.Println("nil", a)
    case Dog, Cat: // a的类型是 Animal
        fmt.Println(a) // 输出 {}
        // fmt.Println(a.name) 这里会报错，因为 Animal 类型没有成员name
    case Tiger: // a的类型是 Tiger
        fmt.Println(a.shout(), a.name) // 这里可以直接取出 name 成员
    default: // a的类型是 Animal
        fmt.Println("default", reflect.TypeOf(a), a)
    }
}

~~~
在上述代码中，我们可以看到a := animal.(type)语句隐式地为每个case子句声明了一个变量a。

变量a类型的判定规则如下:

    如果case后面跟着一个类型，那么变量a在这个case子句中就是这个类型。例如在case Tiger子句中a的类型就是Tiger
    如果case后面跟着多个类型，那么变量a的类型就是接口变量animal的类型，例如在case Dog, Cat子句中a的类型就是Animal
    如果case后面跟着nil，那么变量a的类型就是接口变量animal的类型Animal，通常这种子句用来判断未赋值的接口变量
    default子句中变量a的类型是接口变量animal的类型

为了更好地理解上述规则，我们可以用if语句和类型断言来重写这个switch语句，如下所示：
~~~go
v := animal   // animal 只会被求值一次
    if v == nil { // case nil 子句
        a := v
        fmt.Println("nil", a)
    } else if a, isTiger := v.(Tiger); isTiger { // case Tiger 子句
        fmt.Println(a.shout(), a.name)
    } else {
        _, isDog := v.(Dog)
        _, isCat := v.(Cat)
        if isDog || isCat { // case Dog, Cat 子句
            a := v
            fmt.Println(a)
            // fmt.Println(a.name)
        } else { // default 子句
            a := v
            fmt.Println("default", reflect.TypeOf(a), a)
        }
    }

~~~
