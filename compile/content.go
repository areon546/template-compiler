package compile

import (
	"github.com/areon546/go-files/files"
)

type content struct {
	markdown *files.File
}

func newContent(path, name string) *content {
	file := files.OpenFile(contentDir + "/" + path + name)
	return &content{markdown: file}
}

func (c content) String() string {
	return c.getHTML()
}

func (c content) getHTML() string {
	markdownContents := c.markdown.Contents()
	return string(parseMarkdownToHtml(markdownContents))
}
