package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/picolloo/go-playground/handlers"
)

func main() {

  logger := log.New(os.Stdout, "product-api", log.LstdFlags)
  postHandler := handlers.NewPostHandler(logger)

  mux := http.NewServeMux()
  mux.Handle("/", postHandler)

  server := &http.Server{
    Addr: ":3000",
    Handler: mux,
    IdleTimeout: 120 * time.Second,
    ReadTimeout: 1 * time.Second,
    WriteTimeout: 1 * time.Second,
  }

  go func() {
    err := server.ListenAndServe()
    if err != nil {
      log.Fatal(err.Error())
    }
  }()

  sigChan := make(chan os.Signal)
  signal.Notify(sigChan, os.Interrupt, os.Kill)
  signal := <- sigChan
  logger.Println("Received terminate, graceful shutdown", signal)

  ctx, _ := context.WithTimeout(context.Background(), 30 * time.Second)
  server.Shutdown(ctx)
}
