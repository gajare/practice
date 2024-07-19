package main

import (
	"errors"
	"fmt"
)

func main() {
	err := f1()
	fmt.Println("err :", err)
	err2 := f2()
	fmt.Println("err2 :", err2)
}

func f1() error {
	err := errors.New("func 1 error")
	return err
}
func f2() error {

	err := f1()
	if err != nil {
		return errors.Join(err, errors.New("error from func 2"))
	}
	return nil
}
