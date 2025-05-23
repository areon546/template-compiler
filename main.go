package main

import (
	"flag"
	"os"
)

// set defaults for now defaults, which would normally be set by arguments
var (
	templateDir string = "../templates"
	contentDir  string = "../content"
	outputDir   string = "../docs"
)

func main() {
	// check args

	print("args: ", os.Args)

	// declare flags to use
	flag.StringVar(&templateDir, "t", templateDir, "Specify template directory. Default is "+templateDir)

	flag.Parse()

	CompileTemplates(contentDir, templateDir)
}
