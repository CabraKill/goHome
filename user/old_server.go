// +build ignore

package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	/*if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}*/

	fmt.Printf("Request type: %s\n", r.Method)

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Ih, rapaz. Você pediu demais de mim nesse request :/")
		//http.ServeFile(w, r, "/index.html")
	case "POST":
		fmt.Println(r.URL)
		fmt.Fprintf(w, "Hi, client.")

	default:
		fmt.Fprintf(w, "Ih, rapaz. Você pediu demais de mim nesse request :/")
	}
}

func post_handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		log.Fatal(err)
		return
	}
	/*//fmt.Fprintf(w, "POST request successful")
	fmt.Printf("POST %s.\n", r.FormValue("name"))
	//w.WriteHeader(http.StatusOK)

	//w.Write([]byte("Hi,Client!"))
	fmt.Fprint(w, "Hi, client.")
	*/

	fmt.Printf("POST %s.\n", r.FormValue("name"))
	
	
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("Hi, client."))

}

func main() {
	fmt.Printf("................Server on................\n")
	http.Handle("/", http.FileServer(http.Dir("./page")))
	http.HandleFunc("/post", post_handler)
	//log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("page/"))))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
