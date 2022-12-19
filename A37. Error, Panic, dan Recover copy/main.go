package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//  fungsi strconv.Atoi(). Fungsi tersebut digunakan untuk konversi data string menjadi numerik
	var input string
	fmt.Print("Type some number: ")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err != nil {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
		} else {
		fmt.Println(number, "is number")
	}

	// if err == nil {
    //     fmt.Println(number, "is number")
    // } else {
    //     fmt.Println(input, "is not number")
    //     fmt.Println(err.Error())
    // }


	// Membuat Custom Error
	var name string
	fmt.Print("Type your name: ")
    fmt.Scanln(&name)

    if valid, err := validate(name); valid {
        fmt.Println("halo", name)
    } else {
		panic(err.Error())
		// fmt.Println("end")
        // fmt.Println(err.Error())
    }

	// Pengunaan Panic
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}