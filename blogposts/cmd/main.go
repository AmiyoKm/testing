package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AmiyoKm/blogposts"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", posts)
}
