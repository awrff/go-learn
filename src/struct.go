package main

import (
	"encoding/json"
	"fmt"
)

const (
	male = iota
	female
)

type Student struct {
	name   string
	age    int
	stuNo  string
	gender int
	string // 院系名，string类型匿名字段
}

func (s Student) printInfo() {
	fmt.Printf("student info: %#v\n", s)
}

func (s *Student) setAge(age int) {
	s.age = age
}

type StudentPublic struct {
	Name   string `json:"na"`
	Age    int
	StuNo  string
	Gender int
	string // 院系名，string类型匿名字段
}

func jsonTest() {
	// struct to json
	stu1 := StudentPublic{
		Name:   "李华",
		Age:    21,
		StuNo:  "8888888",
		Gender: male,
		string: "计算机学院",
	}

	data, err := json.Marshal(stu1)

	if err != nil {
		fmt.Println("struct to json fail!", err)
		return
	} else {
		fmt.Printf("json: %s, type: %T\n", data, data)
	}

	// json to struct
	stuStr := `{"na":"jack","age":22,"stuNo":"999999","gender":0,"string":"计算机学院"}`
	fmt.Printf("json type: %T\n", stuStr)
	stu2 := &StudentPublic{}
	err = json.Unmarshal([]byte(stuStr), stu2)

	if err != nil {
		fmt.Println("json to struct fail!", err)
		return
	} else {
		fmt.Printf("stu2: %#v\n", stu2)
	}
}

func structTest() {
	stu1 := Student{
		name:   "李华",
		age:    21,
		stuNo:  "8888888",
		gender: male,
		string: "计算机学院",
	}
	stu2 := Student{
		"张三",
		24,
		"2222222",
		female,
		"计算机学院",
	}
	// stu1.name = "李可"
	// stu1.age = 22
	// stu1.stuNo = "123456"
	// stu1.gender = male

	fmt.Printf("stu1: %v\n", stu1)
	fmt.Printf("stu1: %+v\n", stu1)
	fmt.Printf("stu1: %#v\n", stu1)

	fmt.Printf("stu2: %v\n", stu2)
	fmt.Printf("stu2: %+v\n", stu2)
	fmt.Printf("stu2: %#v\n", stu2)
}

func structPointer() {
	var stu1 = new(Student)
	stu1.name = "jack"
	stu1.age = 30
	fmt.Printf("stu1 address: %p\n", stu1)
	fmt.Printf("stu1: %#v\n", stu1)
}

func newStudent(name string, age int, stuNo string, gender int, schoolName string) *Student {
	return &Student{
		name,
		age,
		stuNo,
		gender,
		schoolName,
	}
}

func structConstructor() {
	stu1 := newStudent("李四", 19, "4444444", male, "医学院")
	fmt.Printf("stu1: %#v\n", stu1)
}

func methodTest() {
	stu1 := newStudent("王五", 26, "333333", male, "环境学院")
	stu1.printInfo()

	stu1.setAge(36)
	stu1.printInfo()

	fmt.Println("所属学院：", stu1.string)
}

type Person struct {
	name    string
	int     // age
	*string // address
}

func anonymousTest() {
	address := "华南理工大学"
	p1 := Person{
		name:   "jack",
		int:    22,
		string: &address,
	}

	fmt.Printf("name: %s, age: %d, address: %s\n", p1.name, p1.int, *p1.string)
}
