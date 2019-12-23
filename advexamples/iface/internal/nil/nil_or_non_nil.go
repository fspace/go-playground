package nil

type TestStruct struct{}

func NilOrNot(v interface{}) {
	if v == nil {
		println("nil")
	} else {
		println("non-nil")
	}
}
