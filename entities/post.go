package post

import (
	"fmt"
	"time"

	"github.com/go-playground/validator"
	"github.com/imdario/mergo"
)

type Post struct {
	ID          int      `json:"id"`
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Content     string   `json:"content" validate:"required"`
	Tags        []string `json:"tags"`
	Author      string   `json:"-" validate:"required"`
	CreatedAt   string   `json:"-"`
	UpdatedAt   string   `json:"-"`
}

func (p *Post) Validate() error {
	validate := validator.New()
	err := validate.Struct(p)
	return err
}

var postList = []*Post{
	{
		ID:          1,
		Title:       "First blog post",
		Description: "Let me see if this works",
		Content:     "lorem ipsum dolur anmet",
		Tags:        []string{"Go", "Docker"},
		Author:      "Lucas Picollo",
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	},
}

func Create(post *Post) *Post {
	post.ID = findNextId()
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

func Update(p *Post) (*Post, error) {
	_, post, err := findPost(p.ID)

	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v", post)
	mergo.Merge(&post, p)
	fmt.Printf("%+v", post)
	return post, nil
}

func findPost(id int) (int, *Post, error) {
	for idx, p := range postList {
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
	return postList[len(postList)-1].ID + 1
}
