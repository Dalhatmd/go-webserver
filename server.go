package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	fmt.Println("Starting server at port 80")

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
