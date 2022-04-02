package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)

	log.Println("Statring website on port:8081")

	err := http.ListenAndServe(":8081", mux)
	log.Fatal(err)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Saya baru coba golang website"))
}
