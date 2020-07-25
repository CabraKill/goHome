package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"bytes"
	"encoding/json"
)

func main() {
	body,err := json.Marshal(map[string]string{"name": "goname"})
	res, err := http.Post("http://192.168.0.100:8080/post","application/json",bytes.NewBuffer(body))

	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(res.Body)

	res.Body.Close()

	fmt.Printf("Data received: %s\n", data)
}
