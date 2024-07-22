package data

import "github.com/pkg/errors"

func Register(str string) error {
	cause := errors.New("用户已存在")
	err := errors.Wrap(cause, "data failed")
	return err
}
