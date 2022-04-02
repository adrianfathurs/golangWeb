package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type student struct {
	ID    int
	Name  string
	Grade int
}

var data = []student{
	{1, "Adri", 6},
	{2, "Tono", 1},
	{3, "Setyawan", 2},
	{4, "Fathur", 3},
	{5, "12", 6},
}

func main() {
	mux := http.NewServeMux()
	// Declaration routing with func
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/student", studentHandler)
	mux.HandleFunc("/students", studentsHandler)
	// Declaration routing with closure
	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Profile"))
	})
	log.Println("Starting website on port:3000")

	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}

func studentHandler(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")
	var result []byte
	var err error
	// Checking Id is string or not
	idNumb, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "String is not allowed", http.StatusNotFound)
		return
	}
	for _, each := range data {
		if each.ID == idNumb {
			result, err = json.Marshal(each)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(result)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
	return
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
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

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	log.Println(r.URL)
	w.Write([]byte("Welcome To Home"))
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	w.Write([]byte("Hello Saya baru coba golang website"))
}
