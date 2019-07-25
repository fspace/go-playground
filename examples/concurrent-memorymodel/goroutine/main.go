package main

var a string

func f()  {
	print(a)
}

// 不可排序  hello和go f 无法确定谁更先执行
func hello(){
	a = "hello, world"
	go f()
}

// 不断运行该程序  有时候打印hello，world 有时候是空串！
func main() {
	/*
	for i:= 0; i<100; i++ {
		hello()
		time.Sleep(time.Second * 1)
		a = ""
		println("run times: ", i)
	}
	*/
	hello()
}
