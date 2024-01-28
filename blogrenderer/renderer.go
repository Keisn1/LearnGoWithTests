package blogrenderer

import (
	"GoWithTests/blogposts"
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed templates/*
	postTemplates embed.FS
)

func Render(w io.Writer, p blogposts.Post) error {

	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := templ.Execute(w, p); err != nil {
		return err
	}

	return nil
}
