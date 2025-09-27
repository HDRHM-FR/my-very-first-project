package Bank

import (
	"fmt"
)

func Pass() {
	var Password int
	fmt.Printf("please enter your password : \n")
	for {
		fmt.Scanln(&Password)
		switch Password {
		case 1380:
			fmt.Printf("\nhello and welcome :)\n")
			return
		default:
			fmt.Printf("wrong password please try again : \n")
		}

	}

}

func Bal() float64 {
	a := 10000.0
	fmt.Printf("balance : %v\n", a)
	return a

}

func Welcome() {

	var (
		Amo    float64
		choice int
		Opt    int
	)

	Pass()
	b := Bal()
	for {
		fmt.Println("\nplease choose one of the options below")
		fmt.Println("1) to whidraw : ")
		fmt.Println("2) to deposite : ")
		fmt.Println("3) to exit : ")
		fmt.Scanln(&choice)

		if choice == 1 {
			fmt.Printf("withdraw amount : \n")
			fmt.Scanln(&Amo)
			if Amo < b {
				fmt.Printf("take your money please :%v\n", Amo)
				fmt.Printf("\nnew balance : %v\n\n", b-Amo)
				b -= Amo
				fmt.Printf("would you like to continue ?\n1)yes\n2)no\n")
				fmt.Scanln(&Opt)
				if Opt == 1 {
					continue
				} else if Opt == 2 {
					fmt.Println("thanks for using our sevices , goodbye")
					break
				}
			} else if Amo > b {
				fmt.Println("\n\nnot enough money please try again")
				continue
			} else if Amo == b {
				fmt.Printf("withdraw failed : the account must have some balance !!!")
			}

		} else if choice == 2 {
			fmt.Printf("deposite amount :\n")
			fmt.Scanln(&Amo)
			fmt.Printf("new balance :%v\n", b+Amo)
			b += Amo
			fmt.Printf("would you like to continue ?\n1)yes\n2)no\n")
			fmt.Scanln(&Opt)
			if Opt == 1 {
				continue
			} else if Opt == 2 {
				fmt.Println("thanks for using our sevices , goodbye")
				break
			}

		} else if choice == 3 {
			fmt.Println("thaks for using our services :)")
			break

		}

	}

}
