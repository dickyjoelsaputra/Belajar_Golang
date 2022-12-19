package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// seedeing random
	rand.Seed(10)
	fmt.Println("seed ke 10 : " , rand.Int() )

	// unique seed menggunakan unix nano
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("seed random nya :" , rand.Int())

	// Angka Random Index Tertentu
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("index random nya :" , rand.Intn(10)) //sampai angka 10

	// Random Tipe Data String
	fmt.Println(randomString(20))
}

// Random Tipe Data String
func randomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b := make([]rune, length)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}