package compile

import (
	"github.com/areon546/go-files/files"
	"github.com/areon546/go-helpers/helpers"
)

type content struct {
	markdown *files.File
}

func newContent(internalPath string) *content {
	file, err := files.OpenFile(ContentPath(internalPath))
	if err != nil {
		print("newContent err", err)
		helpers.Handle(err)
	}

	return &content{markdown: file}
}

func (c content) String() string {
	return c.getHTML()
}

func (c content) getHTML() string {
	markdownContents := c.markdown.Contents()
	return string(parseMarkdownToHtml(markdownContents))
}
