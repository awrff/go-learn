package main

import "fmt"

func myFunc(args ...int) {
	fmt.Println(args)
	args[0] = 100
	fmt.Println(args)
}

func myFuncTest() {
	arr := [3]int{1, 2, 3}
	sli := []int{4, 5, 6}

	myFunc(arr[:]...)
	myFunc(sli...)
	myFunc(7, 8, 9)
}
