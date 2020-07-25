// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type request struct {
	Command int
}

func setupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func post_handler(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w,r)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	var t request

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Request: %+v.\n", t)

	//w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := fmt.Sprintf("{\"command\": %d ,\"succces\": 1}", t.Command)
	w.Write([]byte(response))

}

func main() {
	fmt.Printf("................Server on................\n")
	http.Handle("/", http.FileServer(http.Dir("./page")))
	http.HandleFunc("/post", post_handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
