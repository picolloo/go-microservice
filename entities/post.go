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

func Create(post *Post) *Post {
  post.ID= findNextId()
  post.CreatedAt = time.Now().String()
  post.UpdatedAt = time.Now().String()
  postList = append(postList, post)

  return post
}

func GetAll() []*Post {
  return postList
}

func Get(id int) (*Post, error) {
  _, post, err := findPost(id)

  if err != nil {
    return nil, err
  }

  return post, nil
}

func Remove(id int) (*Post, error) {
  idx, post, err := findPost(id)

  if err != nil {
    return nil, err
  }

  postList = append(postList[:idx], postList[idx+1:]...)
  return post, nil
}

func findPost(id int) (int, *Post, error) {
  for idx, p := range(postList) {
    if p.ID == id {
      return idx, p, nil
    }
  }
  return -1, nil, fmt.Errorf("Post not found")
}

func findNextId() int {
  if len(postList) == 0 {
    return 1
  }
 return postList[len(postList) - 1].ID + 1
}
