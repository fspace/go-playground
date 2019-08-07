package mypkg

import "fmt"

var pkgInternalVar = "this is a package scope var"

var PkgPublicVar = "this is a pkg public var"

func ReadSomething() string {
	return pkgInternalVar
}

func WriteSomething(in string) error {
	pkgInternalVar = in
	return nil
}

func ExecSomething(params ...interface{}) {
	fmt.Println(params...)
}
