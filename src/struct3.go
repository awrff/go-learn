package main

import (
	"encoding/json"
	"fmt"
)

type Animal2 struct {
	Name string `json:"animal_name"`
}

type Dog2 struct {
	Feet     int `json:"dog_feet"`
	*Animal2 `json:"animal"`
}

func multiStructJson() {
	dog1 := Dog2{
		Feet: 4,
		Animal2: &Animal2{
			Name: "xiaobai",
		},
	}

	data, err := json.Marshal(dog1)
	if err != nil {
		fmt.Println("struct to json fail!", err)
	} else {
		fmt.Printf("json result: %s\n", data)
	}

	// str := []byte(`{"Feet":4,"Name":"dabai"}`)
	str := []byte(`{"dog_feet":4,"animal":{"animal_name":"dabai"}}`)
	dog2 := &Dog2{}
	err = json.Unmarshal(str, dog2)

	if err != nil {
		fmt.Println("json to struct fail!", err)
	} else {
		fmt.Printf("struct result: %#v\n%#v\n", dog2, dog2.Animal2)
	}
}
