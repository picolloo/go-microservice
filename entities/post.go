package entities

import "time"

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

var PostList = []*Post{
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

