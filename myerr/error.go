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

/**
failed to encode args[11]: unable to encode models.FileMd5{Md5:"bdc8829d2a3851fdf78345c9c70c947a", Url:"https://127.0.0.1/static/f

                                                                                                                                                                                                             files/4b5aafad-fb49-460e-9faa-86ddbf006030.conf"} into text format for text (OID 25): cannot find encode plan

*/
