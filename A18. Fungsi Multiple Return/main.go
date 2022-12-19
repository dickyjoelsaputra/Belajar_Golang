package main

import (
	"fmt"
	"math"
)

func main() {
	luas , keliling := calculate(5)
	fmt.Printf("t: %.2f , %.2f" , luas, keliling)

}

// func calculate(data float64) (float64,float64)  {
// 	// hitung luas
// 	var area = math.Pi * math.Pow(data / 2 , 2)
// 	// hitung keliling
// 	var circumference  = math.Pi * data

// 	return area , circumference
// }

// func calculate bisa dibuat lebih sederhana

func calculate(data float64) (area float64, circumference float64)  {
	// hitung luas
	area = math.Pi * math.Pow(data / 2 , 2)
	// hitung keliling
	circumference = math.Pi * data

	return
}