package main

import (
	"flag"
	"os"
	"template-compiler/compile"
	"template-compiler/compile/options"

	"github.com/areon546/go-helpers/helpers"
)

// set defaults for now defaults, which would normally be set by arguments
var (
	templateDir      string = "templates"
	contentDir       string = "content"
	outputDir        string = "docs"
	templateFileType string = "tpl"
	logFileName      string = "compilation.log"
)

func main() {
	run()
	// test()

	compile.CloseLoggers()
}

func run() {
	// check args

	helpers.Print("args: ", os.Args)

	// declare flags to use

	flag.StringVar(&templateDir, "t", templateDir, "Specify template directory. ")
	flag.StringVar(&contentDir, "c", contentDir, "Specify content directory. The content directory contains Markdown files which get inserted into templates. ")
	flag.StringVar(&outputDir, "o", outputDir, "Specify output directory. The output directory contains compiled html files. ")
	flag.StringVar(&templateFileType, "s", templateFileType, "Specify the file type suffix for template files. ")
	flag.StringVar(&logFileName, "l", logFileName, "Specify name of the logfile to write to. ")

	flag.Parse()

	helpers.Print(templateDir, contentDir, outputDir)

	opt := options.NewOptions(templateDir, contentDir, outputDir, templateFileType, logFileName)
	compile.CompileTemplates(opt)
}
