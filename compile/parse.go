package compile

import (
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func parseMarkdownToHtml(markdownContent []byte) (htmlContent []byte) {
	parserExtensions := parser.CommonExtensions
	parser := parser.NewWithExtensions(parserExtensions)

	htmlFlags := html.CommonFlags
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	// markdown -> html
	htmlContent = markdown.ToHTML(markdownContent, parser, renderer)

	return
}

func replaceMDExtensionWith(name, extension string) (newName string, err error) {
	newName, cut := strings.CutSuffix(name, "md")
	if !cut {
		return "", ErrIncorrectFileType
	} else {
		newName += "html"
	}

	return newName, nil
}
