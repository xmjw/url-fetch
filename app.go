package main

import (
	"fmt"
	"log"
	"net/http"
)

func Endpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request to /")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "<h1>Gonna be A-OK</h1>")
	return
}

func main() {
	fmt.Println("Booting HTTP Server")

	http.HandleFunc("/", Endpoint)

	err := http.ListenAndServe(fmt.Sprintf(":8080"), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
