package main

import (
	http2 "eateries-in-kgp/pkg/http"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./cmd/web/")))
	http.HandleFunc("/getRestaurants/", http2.GetRestaurants)
	fmt.Println("Started server at http://localhost:8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}