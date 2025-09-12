package Error

import (
	"fmt"
)

func ER() {
	var number1, number2 int
	fmt.Printf("please enter the first number :\n")
	for {
		_, err1 := fmt.Scanln(&number1)
		if err1 != nil {
			fmt.Printf("error found !!!\n")
			fmt.Println("please try again : ")
			continue
		} else {
			break
		}
	}
	fmt.Println("enter second number : ")
	for {
		_, err2 := fmt.Scanln(&number2)
		if err2 != nil {
			fmt.Println("error found !!!")
			fmt.Println("please try again : ")
			continue
		} else {
			fmt.Printf("result : %v", number1+number2)
			break
		}

	}

}
