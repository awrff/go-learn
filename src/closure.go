package main

import "fmt"

func cnt() func() {
	x := 0
	return func() {
		x += 1
		fmt.Printf("address of x: %p, value of x: %d\n", &x, x)
	}
}

// 外部引用函数参数局部变量
func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

func test01(base int) (func(int) int, func(int) int) {
	// 定义2个函数，并返回
	// 相加
	add := func(i int) int {
		base += i
		return base
	}
	// 相减
	sub := func(i int) int {
		base -= i
		return base
	}
	// 返回
	return add, sub
}

func closureTest() {
	// f1 := cnt()
	// f1()
	// f1()

	// f2 := cnt()
	// f2()

	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2))
	// 此时tmp1和tmp2不是一个实体了
	tmp2 := add(100)
	fmt.Println(tmp2(1), tmp2(2))

	// f1, f2 := test01(10)
	// // base一直是没有消
	// fmt.Println(f1(1), f2(2))
	// // 此时base是9
	// fmt.Println(f1(3), f2(4))
}
