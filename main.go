package main

import (
	"golangWeb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// Declaration routing with func
	mux.HandleFunc("/", handler.RootHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/student", handler.StudentHandler)
	mux.HandleFunc("/studentDisplay", handler.StudentDisplayHandler)
	mux.HandleFunc("/students", handler.StudentsHandler)
	// Declaration routing with closure
	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Profile"))
	})
	log.Println("Starting website on port:3000")

	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
