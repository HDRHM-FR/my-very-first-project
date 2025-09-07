package Quadraticequation

import (
	"fmt"
	"math"
)

type variables struct {
	a float64
	b float64
	c float64
}

func Math() {
	entry := variables{}
	fmt.Println("please enter the first number :")
	fmt.Scanln(&entry.a)
	fmt.Println("please enter the second number :")
	fmt.Scanln(&entry.b)
	fmt.Println("please enter the third number :")
	fmt.Scanln(&entry.c)

	delta := (entry.b * entry.b) - 4*(entry.a*entry.c)
	fmt.Printf("delta : %v\n\n", delta)
	form1 := ((-entry.b) + (math.Sqrt(delta))) / (2 * entry.a)
	form2 := ((-entry.b) - (math.Sqrt(delta))) / (2 * entry.a)
	form3 := (-entry.b) / (2 * entry.a)
	if delta > 0 {
		fmt.Println("answer : ")
		fmt.Printf("x1 : %v\n\nx2 : %v\n\n", form1, form2)
	} else if delta == 0 {
		fmt.Printf("x : %v\n\n", form3)
	} else if delta < 0 {
		fmt.Println("undefined")
		return
	}
}
