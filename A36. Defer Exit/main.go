package main

import (
	"fmt"
	"os"
)

func main() {
	// defer
	defer fmt.Println("halo")
	fmt.Println("selamat datang")

	// exit
	// defer tidak dijalankan setelah IIFE dan print
	func(){
		defer fmt.Println("halo")
		os.Exit(1)
		fmt.Println("selamat datang")
	}()
}

func orderSomeFood(menu string) {
    defer fmt.Println("Terimakasih, silakan tunggu")
    if menu == "pizza" {
        fmt.Print("Pilihan tepat!", " ")
        fmt.Print("Pizza ditempat kami paling enak!", "\n")
        return
    }

    fmt.Println("Pesanan anda:", menu)
}
