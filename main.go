package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8000"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println(err)
	}
}
