package blogpost_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogpost "github.com/AmiyoKm/renderer/read"
)

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("always fail")
}

func TestNewBlogPost(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
	)
	t.Run("opens a directory and reads the files", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, _ := blogpost.NewPostsFromFS(fs)

		if len(posts) != len(fs) {
			t.Fatalf("got %d, want %d", len(posts), len(fs))
		}

		assertPost(t, posts[0], blogpost.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
			Body: `Hello
World`,
		})
	})

	t.Run("returns an error correctly", func(t *testing.T) {
		fs := StubFailingFS{}

		_, err := blogpost.NewPostsFromFS(fs)
		if err == nil {
			t.Error("expected error but didn't execute ", err)
		}
	})
}

func assertPost(t *testing.T, got blogpost.Post, want blogpost.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
