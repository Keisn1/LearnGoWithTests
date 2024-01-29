package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	titleSeparator       string = "Title: "
	descriptionSeparator        = "Description: "
	tagsSeparator               = "Tags: "
	bodySeparator               = "---"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	return Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
		Tags:        readTagsLine(readMetaLine(tagsSeparator)),
		Body:        readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()
	buffer := &bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(buffer, scanner.Text())
	}

	return strings.TrimSuffix(buffer.String(), "\n")
}

func readTagsLine(tagsLine string) []string {
	var tags []string
	for _, tag := range strings.Split(tagsLine, ",") {
		tags = append(tags, strings.TrimSpace(tag))
	}
	return tags
}

func (p Post) Slug() string {
	return strings.ToLower(strings.ReplaceAll(p.Title, " ", "-"))
}
