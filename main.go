package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/picolloo/go-playground/handlers"
)

func main() {

  logger := log.New(os.Stdout, "product-api", log.LstdFlags)
  bodyHandler := handlers.NewBody(logger)

  mux := http.NewServeMux()
  mux.Handle("/", bodyHandler)

  server := &http.Server{
    Addr: ":3000",
    Handler: mux,
    IdleTimeout: 120 * time.Second,
    ReadTimeout: 1 * time.Second,
    WriteTimeout: 1 * time.Second,
  }


  server.ListenAndServe()
}
