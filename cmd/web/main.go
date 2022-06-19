package main

import (
	"fmt"
	"net/http"

	"github.com/duo97/go-webapp/pkg/handlers"
)

const portNumber = ":8000"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println(err)
	}
}
