package main

import (
	"fmt"
	httphandlers "go-training/random-api/http_handlers"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	httphandlers.NewRandomIntHandler(router)
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening in 8081")
	server.ListenAndServe()
}
