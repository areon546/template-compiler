package main

import (
	"flag"
	"html/template"
	"log"
	"os"

	"github.com/areon546/go-helpers/helpers"
)

// set defaults for now defaults, which would normally be set by arguments
var (
	templateDir      string = "templates"
	contentDir       string = "content"
	outputDir        string = "docs"
	templateFileType string = "tpl"
)

func main() {
	run()
	// test()

	closeLoggers()
}

func test() {
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
		{{.}}
	</body>
</html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	data := Web{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}

	err = t.Execute(os.Stdout, data)
	check(err)
}

type Web struct {
	Title string
	Items []string
}

func (w Web) String() string {
	return w.Title
}

func run() {
	// check args

	print("args: ", os.Args)

	// declare flags to use

	flag.StringVar(&templateDir, "t", templateDir, "Specify template directory. ")
	flag.StringVar(&contentDir, "c", contentDir, "Specify content directory. The content directory contains Markdown files which get inserted into templates. ")
	flag.StringVar(&outputDir, "o", outputDir, "Specify output directory. The output directory contains compiled html files. ")
	flag.StringVar(&templateFileType, "s", templateFileType, "Specify the file type suffix for template files. ")

	flag.Parse()

	helpers.Print(templateDir, contentDir, outputDir)

	CompileTemplates()
}
