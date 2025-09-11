package Error

import (
	"fmt"
)

func ER() {
	var number1, number2 int
	for {
		fmt.Printf("please enter the first number :\n")
		_, err1 := fmt.Scanln(&number1)
		if err1 != nil {
			fmt.Printf("error found !!!\n")
			continue
		} else {
			break
		}
	}

	for {
		fmt.Println("enter second number : ")
		_, err2 := fmt.Scanln(&number2)
		if err2 != nil {
			fmt.Println("error found !!!")
			continue
		} else {
			fmt.Printf("result : %v", number1+number2)
			break
		}
	}

}
