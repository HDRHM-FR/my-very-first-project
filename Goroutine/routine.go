package Goroutine

import (
	"fmt"
	"time"
)

// section : go routine
// note : it is a simple mathmatic written with go routine
// HEADS UP : i added the error handling to underestand it more

func Routine() {
	go func() {
		time.Sleep(29 * time.Millisecond)
		_, err := Err(s, 4, 6, 2)
		if err != nil {
			panic("")
		} else {
			return
		}
	}()
	time.Sleep(30 * time.Millisecond)
}

func Err(a, b, c, d float64) (float64, error) {
	var (
		e = a + b
		f = c + d
	)
	for {
		_, err1 := fmt.Printf("first value :%v\n", e)
		if err1 != nil {
			fmt.Printf("please enter a valid number :%v\n ", err1)
			continue
		}
		_, err2 := fmt.Printf("second value :%v\n", f)
		if err2 != nil {
			fmt.Printf("please enter a valid number :%v\n ", err2)
			break
		} else {
			fmt.Printf("result :%v\n", f*e)
		}
	}
	return 0, nil

}
