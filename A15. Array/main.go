package main

import "fmt"

func main() {
    var names [4]string
    names[0] = "trafalgar"
    names[1] = "d"
    names[2] = "water"
    names[3] = "law"
    fmt.Println(names[0], names[1], names[2], names[3])

    // tidak akan di eksekusi / error
    // names[4] = "dick"
    // fmt.Println(names[5]) // error
    
    // inisialisasi nilai aray di awal
    fruits := [4]string{"apple", "grape", "banana", "melon"}
    fmt.Println(fruits)
    fmt.Println(len(fruits))

    // array tanpa jumlah elemen
    numbers := [...]int{2, 3, 2, 4, 3}
    fmt.Println(numbers)

    // Array multi dimensi
    var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
    var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
    fmt.Println("numbers1", numbers1)
    fmt.Println("numbers2", numbers2)
    
    // Perulangan Elemen Array Menggunakan Keyword for
    for i := 0; i < len((fruits)); i++ {
        fmt.Println(i , fruits[i])
    }

    // Perulangan Elemen Array Menggunakan Keyword for - range
    for i, fruit := range fruits {
        fmt.Println("elemen :", i,"." , fruit)
    }

    // Perulangan Elemen Array Menggunakan Keyword for - range dengan var _
    for _, fruit := range fruits {
        fmt.Println("elemen :", fruit)
    }

    // Alokasi elemen array dengan keyword make
    var fruitss = make([]string, 2)
    fruitss[0] = "apple"
    fruitss[1] = "manggo"
}