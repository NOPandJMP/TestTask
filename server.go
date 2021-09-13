package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/api/v1/workspace", WorkSpace)
	fmt.Println("Server started on localhost:8000")
	//Starting the server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
