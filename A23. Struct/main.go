package main

import "fmt"

// deklarasi struct
type student struct {
    name string
    grade int
}

// Embedded struct adalah mekanisme untuk menempelkan sebuah struct sebagai properti struct lain. Agar lebih mudah dipahami

type orang struct {
    name string
    age  int
}

type murid struct {
    grade int
    orang
}

func main() {
    // inisialisai object struct
	var s1 student
    s1.name = "john wick"
    s1.grade = 2
    // inisialisai object struct
    s2 := student{"dicky",200}
    // inisialisai object struct
    s3 := student{name: "jason"}

    fmt.Println("name  :", s1.name , "grade :", s1.grade)
    fmt.Println(s2)
    fmt.Println(s3)

    fmt.Println("============================")

    // struct pointer
    var s11 = student{name: "wick", grade: 2}

    var s22 *student = &s11
    fmt.Println("student 1, name :", s11.name)
    fmt.Println("student 4, name :", s22.name)

    s22.name = "ethan"
    fmt.Println("student 1, name :", s11.name)
    fmt.Println("student 4, name :", s22.name)

    fmt.Println("============================")
    // Embedded struct
    people := orang{name: "dicky",age: 23}
    m1 := murid{grade: 20, orang: people} 
    fmt.Println(m1)
    fmt.Println(m1.name)
    fmt.Println("============================")
    // Anonymous Struct
    var people2 = struct{
        grade int
        orang
    }{}

    people2.orang.age = 20
    people2.orang.name = "dickyjoelsss"
    people2.grade = 100

    fmt.Println(people2)
    fmt.Println(people2.orang.name)

    fmt.Println("============================")
    // Kombinasi Slice & Struct
    allStudents := []student{
        {name: "daniel",grade: 20},
        {name: "lenni",grade: 23},
        {name: "maya",grade: 40},
    }

    for _, student := range allStudents {
        fmt.Println(student.name, " punya nilai sebesar : " , student.grade)
    }

    // Nested Struct
    // student struct {
    //     person struct {
    //         name string
    //         age  int
    //     }
    //     grade   int
    //     hobbies []string
    // }

    // Penggunaan Alias
    type Ayam = orang
    
    burung := Ayam{name: "bango",}
    fmt.Println(burung)
}
