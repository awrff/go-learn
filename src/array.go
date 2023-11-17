package main

import "fmt"

func arrayTest() {
	a := [...][3]int{{1, 2, 3}, {3, 5, 7}}
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%d ", v2)
		}
		fmt.Println()
	}
}

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArr2d(arr *[2][3]int) {
	for i1, v1 := range arr {
		fmt.Printf("%d: ", i1)
		for _, v2 := range v1 {
			fmt.Printf("%d ", v2)
		}
		fmt.Println()
	}
}
