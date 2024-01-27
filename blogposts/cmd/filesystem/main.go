package main

import (
	"blogposts"
	"fmt"
	"os"
)

func main() {
	fmt.Println(blogposts.NewPostsFromFs(os.DirFS("./md_files")))
}
