package main

import (
	"GoWithTests/clockface"
	"os"
	"time"
)

func main() {
	clockface.SVGWriter(os.Stdout, time.Now())
}
