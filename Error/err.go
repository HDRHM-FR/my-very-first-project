package Error

import (
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

	// section : "errors" package
	// note : this method is more efficiant

	_, err3 := sum()
	if err3 == nil {
		fmt.Println("")
	} else {
		return
	}
}

func sum() (int, error) {
	var a, b int
	_, err := fmt.Scanln(&a)
	if err != nil {
		panic("please enter a valid number!!")
	}

	_, errr := fmt.Scanln(&b)
	if errr != nil {
		panic("please enter a valid number!!")
	} else {
		fmt.Printf("result :%v\n", a+b)
	}
	return a + b, nil

}
