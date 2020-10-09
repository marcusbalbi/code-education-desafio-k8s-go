package main

import (
	"./balbi"
		"log"
		"fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, balbi.Greeting("Code.education Rocks!"))
}

func main() {
    http.HandleFunc("/", handler)
		fmt.Println("Servidor ouvindo na porta 8000")
		log.Fatal(http.ListenAndServe(":8000", nil))
}