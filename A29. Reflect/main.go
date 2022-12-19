package main

import (
	"fmt"
	"reflect"
)

// Pengaksesan Informasi Property Variabel Objek
type student struct {
    Name  string
    Grade int
}

func (s *student) getPropertyInfo() {
	// menampum objek s student dan value nya
    var reflectV = reflect.ValueOf(s)

	// mengecek apakah objek pointer
    if reflectV.Kind() == reflect.Ptr {
		// melakukan pengambil object reflect aslinya
        reflectV = reflectV.Elem()
    }

	// berisi object s student
    var reflectType = reflectV.Type()

	// perulangan untuk property public lewat method Field()
    for i := 0; i < reflectV.NumField(); i++ {
        fmt.Println("key       :", reflectType.Field(i).Name) // nama property
		fmt.Println("key       :", reflectType.Field(i).Name) // tipe data property
        fmt.Println("nilai     :", reflectV.Field(i).Interface()) // mengambiul nilai properti dalam bentuk interface
        fmt.Println("")
    }
}

func (s *student)SetName(name string)  {
	s.Name = name
}

func main() {
	// Reflection adalah teknik untuk inspeksi variabel, mengambil informasi dari variabel tersebut atau bahkan memanipulasinya. Cakupan informasi yang bisa didapatkan lewat reflection sangat luas, seperti melihat struktur variabel, tipe, nilai pointer, dan banyak lagi.

	// Go menyediakan package reflect, berisikan banyak sekali fungsi untuk keperluan reflection. Pada chapter ini, kita akan belajar tentang dasar penggunaan package tersebut.

	// Dari banyak fungsi yang tersedia di dalam package tersebut, ada 2 fungsi yang paling penting untuk diketahui, yaitu reflect.ValueOf() dan reflect.TypeOf().

	// Fungsi reflect.ValueOf() akan mengembalikan objek dalam tipe reflect.Value, yang berisikan informasi yang berhubungan dengan nilai pada variabel yang dicari
	// Sedangkan reflect.TypeOf() mengembalikan objek dalam tipe reflect.Type. Objek tersebut berisikan informasi yang berhubungan dengan tipe data variabel yang dicari

	number := 10
	nomer := reflect.ValueOf(number)
	fmt.Println(nomer)

	if nomer.Kind() == reflect.Int {
		fmt.Println("tipe variabel :", reflect.TypeOf(number)) // int
		fmt.Println("tipe variabel :", nomer.Kind()) // int
		fmt.Println("nilai variabel :", nomer.Type()) // int
		fmt.Println("nilai variabel :", nomer.Int()) // 10
	}

	fmt.Println("=============================")
	// Pengaksesan Nilai Dalam Bentuk interface{}

	var nilai = 23
	var reflectValue = reflect.ValueOf(nilai)

	fmt.Println("tipe  variabel :", reflectValue.Type())
	fmt.Println("nilai variabel :", reflectValue.Interface())

	var angka = reflectValue.Interface().(int)
	fmt.Println(angka)

	fmt.Println("=============================")
	// Pengaksesan Informasi Property Variabel Objek

	var s1 = &student{Name: "wick", Grade: 2}
    s1.getPropertyInfo()

	fmt.Println("=============================")
	// Pengaksesan Informasi Method Variabel Objek
	var s2 = &student{"dicky",20}
	fmt.Println(s2.Name)

	var reflectVal = reflect.ValueOf(s2)
    var method = reflectVal.MethodByName("SetName")
    method.Call([]reflect.Value{
        reflect.ValueOf("wick"),
    })

    fmt.Println("nama :", s2.Name)

	var name = "dicky"
	// var rv = reflect.ValueOf(name)
	s2.SetName(name)

	fmt.Println(s2)
}
