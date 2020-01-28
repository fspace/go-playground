package ch2

import "fmt"

// ====================================================================
type ErrInvalidStatusCode int

func (code ErrInvalidStatusCode) Error() string {
	return fmt.Sprintf("Expected 200 ok, but got %d", code)
}

// ====================================================================

func DoSomeWebRequest() {
	var (
		err error
	)
	myStatusCode := 201 // 此处可以调用第三方http请求
	if myStatusCode != 200 {
		err = ErrInvalidStatusCode(myStatusCode)
	}
	if err != nil {
		fmt.Printf("We get an error: %#v", err)
	}
}
