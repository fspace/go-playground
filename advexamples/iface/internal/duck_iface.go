package internal

type Duck interface {
	Walk()
	Quack()
}

func DuckWalk(d Duck, dc ...Duck) {
	d.Walk()
	if len(dc) > 0 {
		for _, dx := range dc {
			dx.Walk()
		}
	}
}
func DuckQuack(d Duck, dc ...Duck) {
	d.Quack()
	if len(dc) > 0 {
		for _, dx := range dc {
			dx.Quack()
		}
	}
}
