package handler

import (
	"encoding/json"
	"fmt"
	"golangWeb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	var id = r.FormValue("id")
	// var id = r.URL.Query().Get("id") // Same syntax with
	var result []byte
	var err error
	// Checking Id is string or not
	idNumb, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "String is not allowed", http.StatusNotFound)
		return
	}
	var data = []entity.Student{
		{1, "Adri", 6},
		{2, "Tono", 1},
		{3, "Setyawan", 2},
		{4, "Fathur", 3},
		{5, "12", 6},
	}

	for _, each := range data {
		if each.ID == idNumb {
			result, err = json.Marshal(each)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(result)
			fmt.Fprintf(w, "Hasil : %d ", idNumb)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
	return
}

func StudentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data = []entity.Student{
		{1, "Adri", 6},
		{2, "Tono", 1},
		{3, "Setyawan", 2},
		{4, "Fathur", 3},
		{5, "12", 6},
	}
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

func StudentDisplayHandler(w http.ResponseWriter, r *http.Request) {
	var id = r.FormValue("id")
	// var id = r.URL.Query().Get("id") // Same syntax with
	var err error
	// Checking Id is string or not
	idNumb, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "String is not allowed", http.StatusNotFound)
		return
	}
	var data = []entity.Student{
		{ID: 1, Name: "Adri", Grade: 6},
		{ID: 2, Name: "Tono", Grade: 1},
		{ID: 3, Name: "Setyawan", Grade: 2},
		{ID: 4, Name: "Fathur", Grade: 3},
		{ID: 5, Name: "12", Grade: 6},
	}
	for _, each := range data {
		if each.ID == idNumb {
			log.Println(idNumb)
			tmp, err := template.ParseFiles(path.Join("views", "student.html"), path.Join("views", "layout.html"))

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			log.Println(data)

			err = tmp.Execute(w, each)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println(err)
				return
			}
		}
	}
}
