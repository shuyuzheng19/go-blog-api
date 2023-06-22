package myerr

import (
	"gin-demo/common"
)

func MessageError(err error, message string) {
	if err != nil {
		panic(common.Fail(message))
	}
}

func ResultError(err error, result common.Result) {
	if err != nil {
		panic(result)
	}
}

func PanicError(result common.Result) {
	panic(result)
}

func IError(err error, code int) {
	if err != nil {
		panic(common.IError(code, err.Error()))
	}
}
