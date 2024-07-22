package service

import (
	"gin_error/data"
)

func Register(uname string) error {
	str := "hello" + uname
	err := data.Register(str)
	if err != nil {
		return err
	}
	return nil
}
