package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
	
    "github.com/gorilla/mux"
)

// Define a struct to represent the data being posted
type PostData struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

// Handler function for the POST request
func createPostHandler(w http.ResponseWriter, r *http.Request) {
    var postData PostData

    // Decode the JSON request body into the PostData struct
    err := json.NewDecoder(r.Body).Decode(&postData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Do something with the data, for example, print it to the console
    fmt.Printf("Received: %+v\n", postData)

    // Send a response back to the client
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(postData)
}

func main() {
    r := mux.NewRouter()

    // Define the route for the POST request
    r.HandleFunc("/post", createPostHandler).Methods("POST")

    // Start the server
    fmt.Println("Server is running on port 8000...")
    log.Fatal(http.ListenAndServe(":8000", r))
}