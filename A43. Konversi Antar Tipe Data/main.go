package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Konversi Menggunakan strconv
	var str = "124"
	var num, err = strconv.Atoi(str)

	if err != nil {
		fmt.Println(err.Error())
	}else{
		fmt.Println(num) // 124
	}

	// strconv.Itoa()
	var num1 = 124
	var str1 = strconv.Itoa(num1)
	
	fmt.Println(str1) // "124"

	// Type Assertions Pada Interface Kosong (interface{})

	var data = map[string]interface{}{
		"nama":    "john wick",
		"grade":   2,
		"height":  156.5,
		"isMale":  true,
		"hobbies": []string{"eating", "sleeping"},
	}

	fmt.Println(data["nama"].(string))
	fmt.Println(data["grade"].(int))
	fmt.Println(data["height"].(float64))
	fmt.Println(data["isMale"].(bool))
	fmt.Println(data["hobbies"].([]string))
	
}