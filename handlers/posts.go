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

  post, err := post.Get(product_id)
  if err != nil {
    http.Error(rw, err.Error(), http.StatusBadRequest)
    return
  }

  encoder := json.NewEncoder(rw)
  err = encoder.Encode(post)
  if err != nil {
    http.Error(rw, err.Error(), http.StatusBadRequest)
    return
  }
}

func (h *PostHandler) UpdatePost(rw http.ResponseWriter, req *http.Request) {
  vars := mux.Vars(req)
  post_id, _ := strconv.Atoi(vars["id"])

  var p *post.Post

  decoder := json.NewDecoder(req.Body)
  err := decoder.Decode(&p)
  if err != nil {
    http.Error(rw, err.Error(), http.StatusBadRequest)
    return
  }

  h.logger.Println(p)

  p.ID = post_id
  h.logger.Println(p)

  p, err = post.Update(p)
  h.logger.Println(p)
  if err != nil {
    http.Error(rw, err.Error(), http.StatusBadRequest)
    return
  }

  encoder := json.NewEncoder(rw)
  err = encoder.Encode(&p)

  if err != nil {
    http.Error(rw, "Unable to serialize post.", http.StatusBadRequest)
  }
}

func (h *PostHandler) Delete(rw http.ResponseWriter, req *http.Request) {
  vars := mux.Vars(req)
  post_id, _ := strconv.Atoi(vars["id"])

  post, err := post.Remove(post_id)
  if err != nil {
    http.Error(rw, err.Error(), http.StatusNotFound)
    return
  }

  encoder := json.NewEncoder(rw)
  encoder.Encode(&post)
}

func (h *PostHandler) HitLogginMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func (rw http.ResponseWriter, req *http.Request) {
    h.logger.Printf("Hit on PostHandler %s", req.Method)

    next.ServeHTTP(rw, req)
  })
}
