package main

import (
	"strings"

	"github.com/areon546/go-files/files"
)

type Content struct {
	markdown files.TextFile
}

func (c Content) generateContent() string {
	markdown := c.markdown.Contents()

	return strings.Join(markdown, "/n")
}

func (c Content) String() string {
	return c.generateContent()
}
