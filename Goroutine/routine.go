package Goroutine

import (
	"errors"
	"fmt"
	"time"
)

// section : go routine
// note : it is a simple mathmatic written with go routine
// HEADS UP : i added the error handling to underestand it more

func Routine() {
	go func() {
		time.Sleep(1 * time.Millisecond)
		_, err := CheckInputs(2, 6, 4, 3)
		if err != nil {
			panic(err)
		} else {
			return
		}
	}()
	time.Sleep(30 * time.Millisecond)
}

func CheckInputs(a, b, c, d float64) (float64, error) {
	var (
		e      = a + b
		f      = c + d
		result = e * f
	)

	fmt.Printf("first value :%v\n", e)
	fmt.Printf("second value :%v\n", f)

	if result < 10 {
		return 0, errors.New("the multiplication is less than 10")
	} else {
		fmt.Printf("result :%v", result)
	}

	return result, nil
}
