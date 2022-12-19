package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type student struct {
	ID    string
	Name  string
	Grade int
}

var data = []student{
	{"E001", "ethan", 21},
	{"W001", "wick", 22},
	{"B001", "bourne", 23},
	{"B002", "bond", 23},
}

func main() {
	http.HandleFunc("/users", users)
    http.HandleFunc("/user", user)

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}

func users(w http.ResponseWriter, r *http.Request) {
	// menentukan tipe response, yaitu sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// cek jika method nya GET
	if r.Method == "GET" {
		// konfersi data ke json
        var result, err = json.Marshal(data)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Write(result)
        return
    }

    http.Error(w, "", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request)  {
	// menentukan tipe response, yaitu sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET"{
		id := r.FormValue("id")
		var result []byte
		var err error

		for _, each := range data {
			if each.ID == id {
				result, err = json.Marshal(each)

				if err != nil{
					http.Error(w,err.Error(),http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}
		http.Error(w, "User not found", http.StatusNotFound)
        return
	}
	http.Error(w, "", http.StatusBadRequest)
}