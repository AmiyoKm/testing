package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	blogpost "github.com/AmiyoKm/renderer/read"
	blogrenderer "github.com/AmiyoKm/renderer/render"
)

func TestBlogServer(t *testing.T) {
	posts := []blogpost.Post{
		{Title: "Post 1", Description: "This is a post"},
		{Title: "Post 2", Description: "This is another post"},
	}
	renderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	server := httptest.NewServer(NewBlogServer(posts, renderer))
	defer server.Close()

	t.Run("it returns the index page", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, server.URL, nil)
		response, _ := http.DefaultClient.Do(request)

		if response.StatusCode != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, response.StatusCode)
		}
	})

	t.Run("it returns a specific post", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, server.URL+"/post/Post-1", nil)
		response, _ := http.DefaultClient.Do(request)

		if response.StatusCode != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, response.StatusCode)
		}
	})
}
