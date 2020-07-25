package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gookit/color"

	databaseDriver "github.com/goHome/devices/driver"
	model "github.com/goHome/devices/model"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	optionsSendOk(&w, r)

	/*if r.Method == "POST" {
		postHandler(&w, r)
		return
	} else if r.Method == "GET" {
		response := databaseDriver.DataBase.ToString()
		w.Write([]byte(response))
	}*/

	switch method := r.Method; method {
	case "GET":
		//response := databaseDriver.DataBase.ToString()
		response := "ola"
		w.Write([]byte(response))
		return
	case "POST":
		postHandler(w, r)
		return
	default:
		response := "Method not implemented ;-;\nSorry, try again."
		w.Write([]byte(response))
	}

}

func setupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func optionsSendOk(w *http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		(*w).WriteHeader(http.StatusOK)
		return
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	optionsSendOk(&w, r)

	var device model.DeviceModel

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&device)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//log.Printf("Request: %+v.\n", device)
	color.Green.Printf("Request: %+v.\n", device)
	databaseDriver.DataBase.AddDevice(device)
	databaseDriver.DataBase.Show()

	//w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := databaseDriver.DataBase.ToString()
	w.Write([]byte(response))

}

const (
	port = 8081
)

func main() {
	//.NewDataBaseJson()

	//fmt.Printf("................Server on................\n")
	color.Cyan.Println("................................Server on................................")
	//http.Handle("/", http.FileServer(http.Dir("./page")))
	http.HandleFunc("/", mainHandler)

	http.HandleFunc("/post", postHandler)
	if err := http.ListenAndServe(":"+strconv.FormatInt(port, 10), nil); err != nil {
		log.Fatal(err)
	}
}
