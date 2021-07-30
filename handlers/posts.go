package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/picolloo/go-playground/entities"
)


type PostHandler struct {
 logger *log.Logger
}

func (h *PostHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
  h.logger.Println("Hit on PostHandler Endpoint")

  enoder := json.NewEncoder(rw)
  err := enoder.Encode(entities.PostList)
  if err != nil {
    http.Error(rw, err.Error(), http.StatusBadRequest)
    return
  }
}

func NewPostHandler(l *log.Logger) *PostHandler {
  return &PostHandler{
    l,
  }
}
