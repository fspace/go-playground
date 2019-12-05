package main

import "fmt"

/**
- golang 避坑指南(1)interface 之坑多多
*/
type Talkable interface {
	TalkEnglish(string)
	TalkChinese(string)
}

type Student1 struct {
	Talkable // 内嵌接口  按需实现
	Name     string
	Age      int
}

func (s *Student1) TalkEnglish(s1 string) {
	fmt.Printf("I'm %s,%d years old,%s", s.Name, s.Age, s1)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic : ", r)
		}
	}()
	a := Student1{Name: "aaa", Age: 12}
	a.TalkEnglish("nice to meet you\n")

	a.TalkChinese("汉语呀 你会说么")
}
