package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Success (Response from Backend 1)")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Backend 1 listening on port 9001")
	log.Fatal(http.ListenAndServe(":9001", nil))
}
