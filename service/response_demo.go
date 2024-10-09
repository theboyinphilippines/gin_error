package service

import (
	"gin_error/res"
)

func Response_demo() (string, *res.Error) {
	data := "sdsd"
	//err := errors.New("write error")
	return data, res.SqlError
}
