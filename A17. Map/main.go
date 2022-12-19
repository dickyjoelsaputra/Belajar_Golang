package main

import "fmt"

func main() {

	var chicken map[string]int
	// chicken = map[string]int{}
	chicken = map[string]int{"januari" : 1000 , "februari" : 40 , "maret" : 120 , "april" : 234}
	// key string , value nya int

	// chicken["januari"] = 50
	// chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("februari", chicken["februari"])         // februari 40


	// for - range pada map
	for i, v := range chicken {
		fmt.Println(i , " adalah " , v , " , ")	
	}

	// delete map
	delete(chicken,"februari")
	fmt.Println(chicken)

	// Deteksi Keberadaan Item Dengan Key Tertentu
	var value, isExist = chicken["april"]

	if isExist {
		fmt.Println(value , isExist)
	} else {
		fmt.Println("item is not exists")
	}

	// Kombinasi Slice & Map
	chicks := []map[string]string{
		{"name": "chicken blue",   "gender": "male"},
		{"name": "chicken red",    "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	for i , v := range chicks {
		fmt.Println(i , "." , v["name"] , "=" , v["gender"])
	}
}