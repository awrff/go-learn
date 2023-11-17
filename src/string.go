package main

import "fmt"

func printString() {
	var num = 10
	var s = "hello"
	var fmtString = "num=%d, s=%s"
	// fmt.Printf(fmtString, num, s)
	var fmtString2 = fmt.Sprintf(fmtString, num, s)
	fmt.Println("Sprintf: " + fmtString2)
}
