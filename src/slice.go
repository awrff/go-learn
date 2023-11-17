package main

import "fmt"

func sliceTest() {
	arr := [...]int{5, 2, 1, 7, 8}

	sli1 := arr[0:3]
	sli2 := arr[2:4]

	fmt.Println(sli1)
	fmt.Println(sli2)

}

func sliceModify() {
	arr := [...]int{5, 2, 1, 7, 8}
	sli1 := arr[0:3]
	sli1[0] = 10
	fmt.Println(sli1, arr)
	sli2 := sli1
	sli2[1] = 20
	fmt.Println(sli2, sli1, arr)
}

func sliceCopy() {
	arr := [...]int{5, 2, 1, 7, 8}
	sli1 := arr[0:5]
	sli2 := make([]int, 3, 3)
	copy(sli2, sli1)
	fmt.Println(sli1, sli2)
}

func stringSlice() {
	str1 := "hello, world"
	str2 := "你好，世界"
	sli1 := []byte(str1)
	sli2 := []rune(str2)

	fmt.Println(string(sli1), string(sli2))

	sli1[0] = 'k'
	sli2[0] = '我'

	fmt.Println(string(sli1), string(sli2))
}
