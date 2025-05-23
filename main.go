package main

import (
	"flag"
	"os"

	"github.com/areon546/go-helpers/helpers"
)

// set defaults for now defaults, which would normally be set by arguments
var (
	templateDir string = "templates"
	contentDir  string = "content"
	outputDir   string = "docs"
)

func main() {
	// check args

	print("args: ", os.Args)

	// declare flags to use

	flag.StringVar(&templateDir, "t", templateDir, "Specify template directory. ")
	flag.StringVar(&contentDir, "c", contentDir, "Specify content directory. The content directory contains Markdown files which get inserted into templates. ")
	flag.StringVar(&outputDir, "o", outputDir, "Specify output directory. The output directory contains compiled html files. ")

	flag.Parse()

	helpers.Print(templateDir, contentDir, outputDir)

	CompileTemplates()
}
