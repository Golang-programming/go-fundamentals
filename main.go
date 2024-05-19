package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" || r.Method != "GET" {
		http.Error(w, "Not found", http.StatusNotFound)
	}

	fmt.Fprintf(w, "Hello! buddy")
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" || r.Method != "POST" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform error: %v", err)
		return
	}

	fmt.Fprintf(w, "Request successful")

	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", handleForm)
	http.HandleFunc("/hello", handleHello)

	fmt.Printf("Go server is running on PORT 8080\n")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
