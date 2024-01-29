package blogrenderer_test

import (
	"GoWithTests/blogposts"
	"GoWithTests/blogrenderer"
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestConvertHtml(t *testing.T) {
	var (
		aPost = blogposts.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
		bPost = blogposts.Post{
			Title:       "hello mars",
			Body:        "This is a post on mars",
			Description: "This is a description of mars",
			Tags:        []string{"go mars", "tdd mars"},
		}
		cPost = blogposts.Post{
			Title:       "hello jupiter",
			Body:        "This is a post on jupiter",
			Description: "This is a description of jupiter",
			Tags:        []string{"go jupiter", "tdd jupiter"},
		}
		mdPost = blogposts.Post{
			Title: "hello jupiter",
			Body: `# Index

-   lists all of the posts, with hyperlinks to view the specific post

## Tests

-   same question, how would you test that`,
			Description: "This is a description of jupiter",
			Tags:        []string{"go jupiter", "tdd jupiter"},
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := &bytes.Buffer{}

		if err := postRenderer.Render(buf, aPost); err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})

	t.Run("renders the index", func(t *testing.T) {
		posts := []blogposts.Post{aPost, bPost, cPost}
		buf := &bytes.Buffer{}

		if err := postRenderer.RenderIndex(buf, posts); err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})

	t.Run("Render markdown", func(t *testing.T) {
		postRenderer, err := blogrenderer.NewPostRenderer()
		if err != nil {
			t.Fatal(err)
		}
		buf := &bytes.Buffer{}

		if err := postRenderer.Render(buf, mdPost); err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})

}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogposts.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	b.ResetTimer()
	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
