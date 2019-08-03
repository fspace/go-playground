package main

import (
	"fmt"
	"github.com/pkg/errors"
	myerrors "github.com/pkg/errors"
)

// @see https://medium.com/hackernoon/golang-handling-errors-gracefully-8e27f1db729f
// @see https://github.com/henrmota/errors-handling-example
func main() {
	err := myerrors.Wrapf(errors.New("old error"), "error getting the result with id %d", 1)
	fmt.Println(err.Error())
}
