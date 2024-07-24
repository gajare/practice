package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Define a struct to pass data to the template
type PageData struct {
	Title   string
	Message string
}

// Handler function to render the HTML view
func renderViewHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:   "Welcome to My Website",
		Message: "This is a message from the Go server.",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}

func main() {
	r := mux.NewRouter()

	// Define the route to render the view
	r.HandleFunc("/", renderViewHandler).Methods("GET")

	// Start the server
	log.Println("Server is running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
