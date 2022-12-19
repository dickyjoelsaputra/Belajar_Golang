package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Penerapan Regexp
	var text = "banana burger soup"
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	var res1 = regex.FindAllString(text, 2)
	fmt.Printf("%#v \n", res1)
	// []string{"banana", "burger"}

	var res2 = regex.FindAllString(text, -1)
	fmt.Printf("%#v \n", res2)
	// []string{"banana", "burger", "soup"}

	// Method MatchString()
	var isMatch = regex.MatchString(text)
	fmt.Println(isMatch)

	// FindString()
	var str = regex.FindString(text)
	fmt.Println(str)

	// Method FindStringIndex()
	var idx = regex.FindStringIndex(text)
	fmt.Println(idx)
	// [0, 6]
	
	var str1 = text[0:8]
	fmt.Println(str1)
	// "banana"


	// Method FindAllString()

	var str2 = regex.FindAllString(text, -1)
	fmt.Println(str2)
	// ["banana", "burger", "soup"]

	var str3 = regex.FindAllString(text, 2)
	fmt.Println(str3)
	// ["banana"]

	// Method ReplaceAllString()

	var strr = regex.ReplaceAllString(text, "potato")
	fmt.Println(strr)
	// "potato potato potato"
}
