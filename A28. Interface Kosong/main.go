package main

import (
	"fmt"
	"strings"
)

// Casting Variabel Interface Kosong Ke Objek Pointer
type person struct {
    name string
    age  int
}

type buah struct {
	nama string
	ukuran int
	busuk bool
}

func main() {
	// Interface kosong atau empty interface yang dinotasikan dengan interface{} atau any (per go v1.18), merupakan tipe data yang sangat spesial. Variabel bertipe ini bisa menampung segala jenis data, bahkan array, pointer, apapun. Tipe data dengan konsep ini biasa disebut dengan dynamic typing

	// var secret interface{}
	var secret any

	secret = "ethan hunt"
	fmt.Println(secret)

	secret = []string{"apple", "manggo", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	// yang lain 
	// var data map[string]interface{}
	var data map[string]any

	data = map[string]interface{}{
		"name":      "ethan hunt",
		"grade":     2,
		"breakfast": []string{"apple", "manggo", "banana"},
	}
	
	// dari sini terlihat bahwa interface{} merupakan tipe data , bukan objek (berbeda dengan interface method)

	fmt.Println(data)


	// CASTING VAR INTEFACE{}
	var rahasia any

	rahasia = 2
	var number = rahasia.(int) * 10
	fmt.Println(rahasia, "multiplied by 10 is :", number)

	rahasia = []string{"apple","mango","banana"}
	var gruits = strings.Join(rahasia.([]string), ",")
	fmt.Println(gruits, "is my favorite fruit")


	// Casting Variabel Interface Kosong Ke Objek Pointer Struct
	var orang any = &person{name: "dicky",age: 23}
	var name = orang.(*person).name

	fmt.Println(name)
	fmt.Println(orang)


	// Kombinasi Slice, map, dan interface{}

	var person = []map[string]interface{}{
		{"name": "Wick", "age": 23},
		{"name": "Ethan", "age": 23},
		{"name": "Bourne", "age": 22},
	}

	for _, each := range person {
		fmt.Println(each["name"], " age is ", each["age"])
		
	}

	// Dengan memanfaatkan slice dan interface{}, kita bisa membuat data array yang isinya adalah bisa apa saja. Silakan perhatikan contoh berikut.

	fruits := []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10},
		[]string{"manggo", "pineapple", "papaya"},
		"orange",
	}
	
	for _, each := range fruits {
		fmt.Println(each)
	}

	// tester

	var Semangka = buah{"semangka",20,true}
	var semongka any = &buah{"semongko",23,false}

	fmt.Println(Semangka)
	fmt.Println(semongka)
	// fmt.Println(semongka.(buah).nama)
	fmt.Println(semongka.(*buah).nama)
	

	var numberA int = 4
  	var numberB *int = &numberA
	var numberC *int = numberB

	fmt.Println(numberA)
	fmt.Println(numberB)
	fmt.Println(*numberC)

}
