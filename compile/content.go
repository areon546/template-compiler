package compile

import (
	"github.com/areon546/go-files/files"
)

type content struct {
	markdown *files.File
}

func newContent(internalPath string) *content {
	file := files.OpenFile(ContentPath(internalPath))
	return &content{markdown: file}
}

func (c content) String() string {
	return c.getHTML()
}

func (c content) getHTML() string {
	markdownContents := c.markdown.Contents()
	return string(parseMarkdownToHtml(markdownContents))
}
