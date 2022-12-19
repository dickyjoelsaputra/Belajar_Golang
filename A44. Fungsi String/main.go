package main

import (
	"fmt"
	"strings"
)

func main() {
	// strings.Contains()
	var isExists = strings.Contains("john wick", "wick")
	fmt.Println(isExists)

	// strings.HasPrefix()
	var isPrefix1 = strings.HasPrefix("john wick", "jo")
	fmt.Println(isPrefix1) // true
	
	var isPrefix2 = strings.HasPrefix("john wick", "wi")
	fmt.Println(isPrefix2) // false

	// strings.HasSuffix()
	var isSuffix1 = strings.HasSuffix("john wick", "ic")
	fmt.Println(isSuffix1) // false
	
	var isSuffix2 = strings.HasSuffix("john wick", "ck")
	fmt.Println(isSuffix2) // true

	// strings.Count()
	var howMany = strings.Count("ethan hunt", "t")
	fmt.Println(howMany) // 2

	// strings.Index()
	var index1 = strings.Index("ethan hunt", "ha")
	fmt.Println(index1) // 2

	// strings.Replace()
	text := "banana"
	find := "a"
	replaceWith := "o"

	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1) // "bonana"

	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2) // "bonona"
	// replace semua string
	var newText3 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText3) // "bonono"

	// strings.Repeat()
	var str = strings.Repeat("na", 4)
	fmt.Println(str) // "nananana" 

	// strings.Split()
	var string1 = strings.Split("the dark knight", " ")
	fmt.Println(string1[1]) // output: ["the", "dark", "knight"]
	
	var string2 = strings.Split("batman", "")
	fmt.Println(string2) // output: ["b", "a", "t", "m", "a", "n"]

	// strings.Join()

	var data = []string{"banana", "papaya", "tomato"}
	var strs = strings.Join(data, "-")
	fmt.Println(strs) // "banana-papaya-tomato"

	// strings.ToLower()
	omg := strings.ToLower("aLAy")
	fmt.Println(omg) // "alay"

	// strings.ToUpper()
	sstr := strings.ToUpper("eat!")
	fmt.Println(sstr) // "EAT!"
}