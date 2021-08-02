package post

import (
	"fmt"
	"time"
)

type Post struct {
  ID int `json:"id"`
  Title string `json:"title"`
  Description string `json:"description"`
  Content string `json:"content"`
  Tags []string `json:"tags"`
  Author string `json:"-"`
  CreatedAt string `json:"-"`
  UpdatedAt string `json:"-"`
}

var postList = []*Post{
  {
    ID: 1,
    Title: "First blog post",
    Description: "Let me see if this works",
    Content: "lorem ipsum dolur anmet",
    Tags: []string{ "Go", "Docker" },
    Author: "Lucas Picollo",
    CreatedAt: time.Now().String(),
    UpdatedAt: time.Now().String(),
  },
}

func GetAll() []*Post {
  return postList
}

func Get(id int) (*Post, error) {
  for _, p := range(postList) {
    if p.ID == id {
      return p, nil
    }
  }
  return nil, fmt.Errorf("Post not found")
}
