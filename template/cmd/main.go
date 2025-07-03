package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	blogpost "github.com/AmiyoKm/renderer/read"
	blogrenderer "github.com/AmiyoKm/renderer/render"
)

func NewBlogServer(posts []blogpost.Post, render *blogrenderer.PostRenderer) http.Handler {
	mux := http.NewServeMux()

	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		render.RenderIndex(w, posts)
	}

	postHandler := func(w http.ResponseWriter, r *http.Request) {
		title := r.PathValue("title")
		actualTitle := strings.ReplaceAll(title, "%20", " ")
		var post blogpost.Post

		for _, p := range posts {
			if strings.ToLower(p.Title) == actualTitle {
				post = p
			}
		}
		w.Header().Set("Content-Type", "text/html")
		render.Renderer(w, post)
	}

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/post/{title}", postHandler)
	return mux
}

func main() {
	posts, err := blogpost.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	render, err := blogrenderer.NewPostRenderer()
	if err != nil {
		log.Fatal(err)
	}

	server := NewBlogServer(posts, render)

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatal(err)
	}
}
