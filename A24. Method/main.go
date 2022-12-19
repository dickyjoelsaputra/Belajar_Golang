package main

import (
	"fmt"
	"strings"
)

type student struct {
	name  string
	grade int
}

func main() {
	// Method adalah fungsi yang menempel pada type (bisa struct atau tipe data lainnya). Method bisa diakses lewat variabel objek.
	var s1 = student{"john wick", 21}
    s1.sayHello()

    var name = s1.getNameAt(1)
    fmt.Println("nama panggilan :", name)

	fmt.Println("============================")
	// Method Pointer
	s1.changeName1("jason bourne")
    fmt.Println("s1 after changeName1", s1.name)
    // john wick

    s1.changeName2("ethan hunt")
    fmt.Println("s1 after changeName2", s1.name)
    // ethan hunt
}

func (s student) sayHello() {
	fmt.Println("halo", s.name)
}

func (s student) getNameAt(i int) string {
    return strings.Split(s.name, " ")[i-1]
}

// Method Pointer


func (s student) changeName1(name string) {
    fmt.Println("---> on changeName1, name changed to", name)
    s.name = name
}

func (s *student) changeName2(name string) {
    fmt.Println("---> on changeName2, name changed to", name)
    s.name = name
}