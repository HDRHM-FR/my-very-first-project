package Goroutine

import (
	"fmt"
	"time"
)

// section : go routine
// note : it is a simple mathmatic written with go routine

func Routine() {
	var (
		a = 15
		b = 9
		c = a + b
		d = 8 - 3
	)
	go func() {
		time.Sleep(29 * time.Millisecond)
		fmt.Printf("first value :%v\n", c)
		fmt.Printf("second value :%v\n", d)
	}()

	time.Sleep(30 * time.Millisecond)
	fmt.Println(c * d)
}
