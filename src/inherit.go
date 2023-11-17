package main

import "fmt"

type Animal struct {
	name string
}

type Dog struct {
	feet    int
	*Animal // 嵌套实现继承
}

func (d *Dog) move() {
	fmt.Println(d.name, "move")
}

func (d *Dog) bark() {
	fmt.Println(d.name, "bark")
}

func AnimalTest() {
	dog1 := Dog{
		feet: 4,
		Animal: &Animal{
			name: "xiaobai",
		},
	}

	fmt.Printf("%s has %d feet\n", dog1.name, dog1.feet)
	dog1.bark()
	dog1.move()
}
