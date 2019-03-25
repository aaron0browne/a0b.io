package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/api", apiHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("coming soon")); err != nil {
		log.Printf("Write error: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
