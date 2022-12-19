package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName 	string `json:"Name"`
	Age      int
}

func main() {
	// 
	// data json
	var jsonString = `{"Name": "john wick", "Age": 27}`
	// konversi json ke byte
	var jsonData = []byte(jsonString)

	//  Decode JSON Ke Variabel Objek Struct

	// inisialisasi struct User
	var data User

	// decode json dan disimpan kedalam data
	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user :", data.FullName)
	fmt.Println("age  :", data.Age)

	// Decode JSON Ke map[string]interface{} & interface{}
	fmt.Println("===========================")

	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("user :", data1["Name"])
	fmt.Println("age  :", data1["Age"])

	var data2 interface{}
	json.Unmarshal(jsonData, &data2)
	
	var decodedData = data2.(map[string]interface{})
	fmt.Println("user :", decodedData["Name"])
	fmt.Println("age  :", decodedData["Age"])

	// Decode Array JSON Ke Array Objek
	fmt.Println("===========================")

	var jsonS = `[
		{"Name": "john wick", "Age": 27},
		{"Name": "ethan hunt", "Age": 32}
	]`
	
	var user []User
	
	var error = json.Unmarshal([]byte(jsonS), &user)
	if error != nil {
		fmt.Println(err.Error())
		return
	}
	
	fmt.Println("user 1:", user[0].FullName)
	fmt.Println("user 2:", user[1].FullName)

	// Encode Objek ke Json String
	fmt.Println("===========================")
	var object = []User{{"john wick", 27}, {"ethan hunt", 32}}
	// var asdf = User{"dicky",20}
	var jsonD, isError = json.Marshal(object)
	if isError != nil {
		fmt.Println(err.Error())
		return
	}
	
	var jsonStr = string(jsonD)
	fmt.Println(jsonStr)
}
