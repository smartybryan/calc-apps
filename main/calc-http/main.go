package main

import (
	"log"
	"net/http"
	"os"

	"github.com/smartybryan/calc-apps/handlers"
)

func main() {
	logger := log.New(os.Stdout, "http: ", 0)
	router := handlers.NewHTTPHandler(logger)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
