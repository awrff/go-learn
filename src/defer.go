package main

import "fmt"

func test() {
	var run = func() {}
	defer run()
}

func deferTest() {
	// x, y := 10, 20
	// defer fmt.Println(x, y)
	// x += 10
	// y += 10
	// fmt.Println(x, y)
	test()
	defer fmt.Println("end")
}

// defer func() {
// 	x += 1
// 	y += 1
// 	fmt.Println(x, y)
// }()
