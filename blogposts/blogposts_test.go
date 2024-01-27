package blogposts_test

import (
	"GoWithTests/blogposts"
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFileSystem struct{}

func (s StubFileSystem) Open(name string) (fs.File, error) {
	return nil, errors.New("Oh no I always fail")

}

func TestFailingFileSystem(t *testing.T) {
	_, err := blogposts.NewPostsFromFs(StubFileSystem{})
	assertError(t, err, errors.New("Oh no I always fail"))
}

func TestReadFile(t *testing.T) {
	var fs = fstest.MapFS{
		"hello.md": {Data: []byte(`Title: Post 1
Description: Description 1
Tags: tag1, tag2
---
Body 1
`)},
		"hola.md": {Data: []byte(`Title: Post 2
Description: Description 2
Tags: tag3, tag4
---
`)},
		"salut.md": {Data: []byte(`Title: Post 3
Description: Description 3
Tags: tag5
---
Body 2
Body 3
`)},
	}

	gotPosts, err := blogposts.NewPostsFromFs(fs)
	assertNoError(t, err)

	wantPosts := []blogposts.Post{
		{
			Title:       `Post 1`,
			Description: "Description 1",
			Tags:        []string{"tag1", "tag2"},
			Body:        "Body 1",
		},
		{
			Title:       `Post 2`,
			Description: "Description 2",
			Tags:        []string{"tag3", "tag4"},
			Body:        "",
		},
		{
			Title:       `Post 3`,
			Description: "Description 3",
			Tags:        []string{"tag5"},
			Body: `Body 2
Body 3`,
		},
	}
	assertLength(t, wantPosts, gotPosts)
	assertPosts(t, wantPosts, gotPosts)
}

func assertLength(t *testing.T, wantPosts, gotPosts []blogposts.Post) {
	t.Helper()
	if len(wantPosts) != len(gotPosts) {
		t.Errorf("Length WantPosts %d not equal lenght gotPosts %d", len(wantPosts), len(gotPosts))
	}
}

func assertPosts(t *testing.T, wantPosts, gotPosts []blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(wantPosts, gotPosts) {
		t.Errorf("%v not equal gotPosts %v", wantPosts, gotPosts)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("Unwanted Error occured: %v", err)
	}
}

func assertError(t *testing.T, err, want error) {
	t.Helper()

	if err == want {
		t.Errorf("Didn't get error: %v, instead %v", want, err)
	}
}
