package post

import (
	"testing"
)

func TestEmptyPostValidation(t *testing.T) {
	post := &Post{}
	err := post.Validate()

	if err == nil {
		t.Fatal(err)
	}
}

func TestFilledPostValidation(t *testing.T) {
	post := &Post{
		Title:       "title",
		Description: "Description",
		Author:      "author",
		Content:     "content",
		Tags:        []string{"a", "b"},
	}
	err := post.Validate()

	if err == nil {
		t.Fatal(err)
	}
}
