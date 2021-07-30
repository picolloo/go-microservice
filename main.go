package main

import (
	"log"
	"net/http"
	"os"

	"github.com/picolloo/go-playground/handlers"
)

func main() {

  logger := log.New(os.Stdout, "product-api", log.LstdFlags)
  bodyHandler := handlers.NewBody(logger)

  mux := http.NewServeMux()
  mux.Handle("/", bodyHandler)

  http.ListenAndServe(":3000", mux)
}
