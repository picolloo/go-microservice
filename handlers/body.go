package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Body struct {
  logger *log.Logger
}

func (b *Body) ServeHTTP(res http.ResponseWriter, req *http.Request) {
   b.logger.Println("Body handler")

    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
      http.Error(res, err.Error(), http.StatusBadRequest)
      return
    }
    res.Write(body)
}


func NewBody(logger *log.Logger) *Body {
  return &Body{
    logger,
  }
}
