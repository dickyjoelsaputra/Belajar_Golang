package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var baseURL = "http://localhost:8080"

type student struct {
    ID    string
    Name  string
    Grade int
}

func main() {
// 	var users, err = fetchUsers()
//     if err != nil {
//         fmt.Println("Error!", err.Error())
//         return
//     }

//     for _, each := range users {
//         fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
//     }

    var user1, err = fetchUser("E001")
    if err != nil {
        fmt.Println("Error!", err.Error())
        return
    }

    fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", user1.ID, user1.Name, user1.Grade)
}

func fetchUsers() ([]student, error) {
    var err error
    //  menghasilkan instance http.Client. Objek ini nantinya diperlukan untuk eksekusi request.
    var client = &http.Client{}
    var data []student

    // http.NewRequest() digunakan untuk membuat request baru. Fungsi tersebut memiliki 3 parameter yang wajib diisi.
    // menghasilkan instance bertipe http.Request
    request, err := http.NewRequest("GET", baseURL+"/users", nil)
    if err != nil {
        return nil, err
    }

    response, err := client.Do(request)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    err = json.NewDecoder(response.Body).Decode(&data)
    if err != nil {
        return nil, err
    }

    return data,nil
}

func fetchUser(ID string) (student, error) {
    var err error
    var client = &http.Client{}
    var data student

    var param = url.Values{}
    param.Set("id", ID)
    var payload = bytes.NewBufferString(param.Encode())

    request, err := http.NewRequest("POST", baseURL+"/user", payload)
    if err != nil {
        return data, err
    }
    request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    response, err := client.Do(request)
    if err != nil {
        return data, err
    }
    defer response.Body.Close()

    err = json.NewDecoder(response.Body).Decode(&data)
    if err != nil {
        return data, err
    }

    return data, nil
}