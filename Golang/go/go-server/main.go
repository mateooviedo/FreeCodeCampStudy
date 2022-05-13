package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
 * @param: { w } @type: { http.ResponseWriter }
 * @param: { r } @type: { http.Request }
 */
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "POST request sucessfully!")

	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprint(w, "\nName = ", name, "\n")
	fmt.Fprint(w, "Address = ", address)
}

/*
 * @param: { w } @type: { http.ResponseWriter }
 * @param: { r } @type: { http.Request }
 */
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8000\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
