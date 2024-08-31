package service

import "errors"

func Response_demo() (string, error) {
	data := "sdsd"
	err := errors.New("write error")
	return data, err
}
