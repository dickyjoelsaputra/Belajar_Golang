package main

import "fmt"

func main() {

    // Variabel Beserta Tipe Data dan Keyword Var
    var firstName string = "john"
    var lastName string
    lastName = "wick"

    // Variabel Tanpa Tipe Data

    nama := "Dicky Ganteng"

    // Multi Variable
    var first, second, third string
    first, second, third = "satu", "dua", "tiga"
    var fourth, fifth, sixth string = "empat", "lima", "enam"
    seventh, eight, ninth := "tujuh", "delapan", "sembilan"
    one , isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

    // Variable _
    _ ="belajar Golang"
    _ = "Golang itu mudah"
    names, _ := "john", "wick"


    // Variable dengan keyword new
    Dicky := new(string)
    Dicky = &nama
    berak := *Dicky

    fmt.Println(firstName,lastName,nama,first,second,third,fourth,fifth,sixth,seventh,eight,ninth,one,isFriday,twoPointTwo,say,names,Dicky,berak)
}