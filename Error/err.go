package Error

import (
	"errors"
	"fmt"
)

// section : Error
func ER() {
	// var (
	// 	number1 int
	// 	number2 int
	// )
	// fmt.Printf("please enter the first number :\n")
	// for {
	// 	_, err1 := fmt.Scanln(&number1)
	// 	if err1 != nil {
	// 		fmt.Println("error found !!!\nplease enter a number : ")
	// 		continue
	// 	} else {
	// 		break
	// 	}
	// }
	// fmt.Println("enter second number : ")
	// for {
	// 	_, err2 := fmt.Scanln(&number2)
	// 	if err2 != nil {
	// 		fmt.Printf("\nerror found !!!\nplease enter a number : ")
	// 		continue
	// 	} else {
	// 		break
	// 	}

	// }

	// fmt.Printf("result : %v", number1+number2)

	// section : errors package
	// note : this method is more efficiant

	_, err3 := sum()
	if err3 != nil {
		panic(err3.Error())

	} else {
		fmt.Print("result : ")
		fmt.Println(sum())
	}
}

func sum() (int, error) {
	var a, b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	return a + b, errors.New("sum couldnt perform, please enter a valid number")
}
