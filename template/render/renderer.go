package blogrenderer

import (
	"embed"
	"html/template"
	"io"

	blogpost "github.com/AmiyoKm/renderer/read"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostRenderer struct {
	templ    *template.Template
	mdParser *parser.Parser
}

func NewPostRenderer() (*PostRenderer, error) {

	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Renderer(w io.Writer, p blogpost.Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", p)
}
func (r *PostRenderer) RenderIndex(w io.Writer, posts []blogpost.Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

type PostViewModel struct {
	blogpost.Post
	HTMLBody template.HTML
}

func NewPostVM(p blogpost.Post, r *PostRenderer) PostViewModel {
	vm := PostViewModel{Post: p}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil))
	return vm
}
