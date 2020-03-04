package ch6

import "errors"

type ValidationError error

var (
	errNoUsername       = ValidationError(errors.New("You must supply a username"))
	errNoEmail          = ValidationError(errors.New("You must supply an email"))
	errNoPassword       = ValidationError(errors.New("You must supply a password"))
	errPasswordTooShort = ValidationError(errors.New("Your password is too short"))
	errUsernameExists   = ValidationError(errors.New("That username is taken"))
	errEmailExists      = ValidationError(errors.New("That email address has an account"))
)

func IsValidationError(err error) bool {
	// ioutil 包中也有类似技巧 用来检测错误类型是否是本包相关的错误
	_, ok := err.(ValidationError)
	return ok
}
