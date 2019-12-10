package lib

type Animal struct {
	Name   string `required max: "100" `
	Origin string
}

type Bird struct {
	Animal   // 内嵌 组合模式
	SpeedKPH float32
	CanFly   bool
}
