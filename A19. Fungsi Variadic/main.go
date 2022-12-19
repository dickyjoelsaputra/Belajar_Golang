package main

import (
	"fmt"
	"strings"
)

func main() {
	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
    fmt.Println("rata rata nya adalah :" , avg)

	// dengan data slice 
	numbers := []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	fmt.Println(calculate(numbers...))

	hobbies := []string{"gaming" , "fuckass" , "genjer"}
	name := "dicky"
	fmt.Println(yourHobbies(name,hobbies...))
}

// vunc calculate 
func calculate(numbers ...int) float64 {
    var total int = 0
    for _ , number := range numbers {
        total += number
    }

    var avg = float64(total) / float64(len(numbers))
    return avg
}

func yourHobbies(name string, hobbies ...string)(string,string){
	hobbiesAsString := strings.Join(hobbies, ", ")
	return name, hobbiesAsString
}