package main

import (
	// "GoWithTests/blogposts"
	// "fmt"
	"github.com/gomarkdown/markdown/parser"
)

func main() {
	// p := blogposts.Post{Body: "thats a body"}
	// ps := []blogposts.Post{p}
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	mdParser := parser.NewWithExtensions(extensions)
	mdParser.Parse([]byte{'a'})
	mdParser = parser.NewWithExtensions(extensions)
	mdParser.Parse([]byte{'b'})
	// parser.Parse([]byte(p.Body))
	// parser.Parse([]byte(p.Body))

	// fmt.Println([]byte(p.Body))
	// fmt.Println([]byte(ps[0].Body))

	// markdown.ToHTML([]byte(p.Body), parser, nil)
	// markdown.ToHTML([]byte(ps[0].Body), parser, nil)
}
