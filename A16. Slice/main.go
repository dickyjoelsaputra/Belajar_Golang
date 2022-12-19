package main

import "fmt"

func main() {

	// perbedaan slice dan array
	var fruitsA = []string{"apple", "grape"}   // slice
	var fruitsB = [2]string{"banana", "melon"} // array
	var fruitsC = [...]string{"papaya", "grape"} //array

	fruitsD := fruitsC[0:1]

	fmt.Println(fruitsA,fruitsB,fruitsC)
	fmt.Println(fruitsD)
	fmt.Println(fruitsC)
// Hubungan Slice Dengan Array & Operasi Slice
	var fruits = []string{"apple", "grape", "banana", "melon"}
	var newFruits = fruits[0:2]

	fmt.Println(newFruits)
}