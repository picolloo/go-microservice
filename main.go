package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	gohanlders "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/picolloo/go-playground/handlers"
)

func main() {

	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	postHandler := handlers.NewPostHandler(logger)

	router := mux.NewRouter()

	router.Use(postHandler.HitLogginMiddleware)
	router.HandleFunc("/posts", postHandler.AddPost).Methods("POST")
	router.HandleFunc("/posts", postHandler.GetPosts).Methods("GET")
	router.HandleFunc("/posts/{id:[0-9]+}", postHandler.GetPost).Methods("GET")
	router.HandleFunc("/posts/{id:[0-9]+}", postHandler.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id:[0-9]+}", postHandler.Delete).Methods("DELETE")

	ch := gohanlders.CORS(gohanlders.AllowedOrigins([]string{"localhost:3000"}))

	server := &http.Server{
		Addr:         ":3000",
		Handler:      ch(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
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
	signal := <-sigChan
	logger.Println("Received terminate, graceful shutdown", signal)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
