package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"example.com/simple-web-server/data"
)

func handleHi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from go server"))
}

func hanldeTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("./templates/index.tmpl")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal serve error"))
		return
	}

	html.Execute(w, data.GetAllMuseums())

}

func handleAddMuseum(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println(r.Method)

	var museum data.Museum

	err := json.NewDecoder(r.Body).Decode(&museum)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data.AddMuseum(museum)

}

func main() {
	http.HandleFunc("/hi", handleHi)

	fs := http.FileServer(http.Dir("./public"))

	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.HandleFunc("/", hanldeTemplate)
	http.HandleFunc("/add-museum", handleAddMuseum)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
