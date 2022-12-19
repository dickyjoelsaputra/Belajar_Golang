package main

import (
	"fmt"
	"strings"
)

func main() {
	nama := []string{"dicky", "joel", "saputra"}

	var dataContainsO = filter(nama, func(each string) bool {
		return strings.Contains(each, "o")
	})

    fmt.Println(dataContainsO)

    var dataLenght5 = filter(nama, func(each string) bool {
        return len(each) == 5
    })

    fmt.Println(dataLenght5)
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
