package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func mapTest() {
	mp1 := map[int]string{}
	mp2 := make(map[int]string, 8)

	mp1[0] = "你好"
	mp1[1] = ","
	mp2[0] = "世界"

	fmt.Println(mp1[0], mp1[1], mp2[0])

	fmt.Print("mp1里是否有key 0: ")
	v1, ok := mp1[0]
	if ok {
		fmt.Println("yes", v1)
	} else {
		fmt.Println("no")
	}

	fmt.Print("mp2里是否有key 1: ")
	v2, ok := mp2[1]
	if ok {
		fmt.Println("yes", v2)
	} else {
		fmt.Println("no")
	}
}

func printMap() {
	mp1 := map[int]string{}
	mp2 := make(map[int]string, 8)

	mp1[0] = "你好"
	mp1[1] = ","
	mp2[0] = "世界"

	fmt.Println("mp1: ")
	for k, v := range mp1 {
		fmt.Println(k, v)
	}

	fmt.Println("mp2: (only key)")
	for k := range mp2 {
		fmt.Println(k)
	}
}

func printSortedMap() {
	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}

	keys := make([]string, 200)
	for k, _ := range scoreMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, scoreMap[k])
	}
}
