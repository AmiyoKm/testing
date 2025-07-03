package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	blogpost "github.com/AmiyoKm/renderer/read"
	blogrenderer "github.com/AmiyoKm/renderer/render"
	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogpost.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRenderer.Renderer(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
	t.Run("it renders an index of a post", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogpost.Post{{Title: "Hello world"}, {Title: "Hello world 2"}}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRenderer(b *testing.B) {
	var (
		aPost = blogpost.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		postRenderer.Renderer(io.Discard, aPost)
	}
}
