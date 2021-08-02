package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/picolloo/go-playground/entities"
)


type PostHandler struct {
 logger *log.Logger
}

func NewPostHandler(l *log.Logger) *PostHandler {
  return &PostHandler{
    l,
  }
}

func (h *PostHandler) AddPost(rw http.ResponseWriter, req *http.Request) {
  h.logger.Println("Hit on PostHandler POST")

  var p *post.Post
  decoder := json.NewDecoder(req.Body)
  err := decoder.Decode(&p)

  if err != nil {
    http.Error(rw, err.Error(), http.StatusBadRequest)
    return
  }

  p = post.Create(p)
  encoder := json.NewEncoder(rw)
  err = encoder.Encode(&p)
  if err != nil {
    http.Error(rw, err.Error(), http.StatusBadRequest)
    return
  }
}

func (h *PostHandler) GetPosts(rw http.ResponseWriter, req *http.Request) {
  h.logger.Println("Hit on PostHandler GetPosts")

  encoder := json.NewEncoder(rw)
  err := encoder.Encode(post.GetAll())
  if err != nil {
    http.Error(rw, err.Error(), http.StatusBadRequest)
    return
  }
}

func (h *PostHandler) GetPost(rw http.ResponseWriter, req *http.Request) {
  vars := mux.Vars(req)
  product_id, _ := strconv.Atoi(vars["id"])

  h.logger.Printf("Hit on PostHandler POST with ID: %d", product_id)

  encoder := json.NewEncoder(rw)

  post, err := post.Get(product_id)
  if err != nil {
    http.Error(rw, err.Error(), http.StatusBadRequest)
    return
  }

  err = encoder.Encode(post)
  if err != nil {
    http.Error(rw, err.Error(), http.StatusBadRequest)
    return
  }
}

func (h *PostHandler) UpdatePost(rw http.ResponseWriter, req *http.Request) {
  vars := mux.Vars(req)
  post_id, _ := strconv.Atoi(vars["id"])

  h.logger.Printf("Hit on PostHandler PUT with ID: %d", post_id)

  post, err := post.Get(post_id)
  if err != nil {
    http.Error(rw, "Invalid post ID", http.StatusBadRequest)
    return
  }

  decoder := json.NewDecoder(req.Body)
  err = decoder.Decode(&post)
  if err != nil {
    http.Error(rw, "Invalid payload", http.StatusBadRequest)
    return
  }

  encoder := json.NewEncoder(rw)
  err = encoder.Encode(&post)

  if err != nil {
    http.Error(rw, "Unable to serialize post.", http.StatusBadRequest)
  }
}

func (h *PostHandler) Delete(rw http.ResponseWriter, req *http.Request) {
  vars := mux.Vars(req)
  post_id, _ := strconv.Atoi(vars["id"])

  h.logger.Printf("Hit on PostHandler DELETE with ID: %d", post_id)

  post, err := post.Remove(post_id)
  if err != nil {
    http.Error(rw, err.Error(), http.StatusNotFound)
    return
  }

  encoder := json.NewEncoder(rw)
  encoder.Encode(&post)
}
