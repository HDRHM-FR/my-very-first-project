package Arrays

import "fmt"

func Array() {

	//this section indicates the arrays
	a := [...]int{3, 4, 6, 5, 10, 1000}
	a[0] = 30
	fmt.Println(a)

	m := [8][2]int{
		{1, 3},
		{6, 2},
	}

	m[3][0] = 56

	fmt.Println(m)

	// section : slices
	// more info : the slices act like arrays but we can put index as many as we want

	b := []int{1, 2, 3, 4, 56, 7, 8, 9, 6}

	// section : the append function
	// more info : append function is used to add more index to our slice
	// Heads up : the append function can only be used in slices not in arrays !!!

	b = append(b, 903, 700)

	fmt.Println(b)

	// using the for loop to range an array or slice :

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// section : maps
	// info : maps are like arrays and slices but we can put 2 or more types of data types

	d := map[int]int{
		1: 2,
		2: 4,
		3: 1,
	}

	fmt.Println(d)

	c := map[string]int{
		"amaliate salam":    1,
		"amaliate khodafez": 2,
	}

	fmt.Println(c)

	j := user{
		1: {"user": "hadi rahimi"},
		2: {"pass": "1234"},
		3: {"request": "red lamburgini"},
	}

	for k, v := range j {
		fmt.Printf("key : %v\nvalue : %v\n\n", k, v)

	}

}

// we can also use costume types for matrix maps
// heads up : its used for "v" variable in Array function :

type user map[int]map[string]string
