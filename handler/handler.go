package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	data := map[string]interface{}{
		"title":   "Hello I'm Golang",
		"message": "Let's GO With Adrian Fathur Setyawan",
	}
	tmp, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Keep Calm, this website maintanance", 404)
		return
	}
	err = tmp.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Keep Calm, this website maintanance2", 404)
		return
	}

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	w.Write([]byte("Hello Saya baru coba golang website"))
}
