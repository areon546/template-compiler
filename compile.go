package main

import (
	"io/fs"

	"github.com/areon546/go-files/files"
	"github.com/areon546/go-helpers/helpers"
)

func CompileTemplates(contentDir, templateDir string) {
	// read contents of template and content directories
	templates := files.ReadDirectory(templateDir)
	content := files.ReadDirectory(contentDir)

	print(templateDir)
	print(templates)

	print(contentDir)
	print(content)

	// crawl the content files and directories
	compileTemplatesRec(".", content, templates)
}

// Crawls through the contents directory and compiles html files based on that.
func compileTemplatesRec(path string, content, templates []fs.DirEntry) {
	for _, dirEntry := range content {
		print(dirEntry)

		if dirEntry.IsDir() {
			// if is a directory, then hanlde it recursively
			handleSubdirectory("", content, templates, dirEntry) // TODO: implement path
		} else if !dirEntry.IsDir() {
			// if is a file, then handle it
			handleFile("", content, templates, dirEntry)
		}
	}
}

func handleSubdirectory(path string, content, templates []fs.DirEntry, directory fs.DirEntry) {
	newContent := content // TODO: make it read the specific directory we want to look at
	newTemplates := templates

	print(directory, newContent, newTemplates)

	// compileTemplatesRec(newContent, newTemplates)
}

func handleFile(path string, content, templates []fs.DirEntry, file fs.DirEntry) {
	// since it is a file, we want to:
	// check for special cases
	//
	// generate template
	//
	// write html file to output directory
}

func print(a ...any) {
	helpers.Print(a...)
}

//

func processTemplates() {
	// This will read the templates folder, execute them, and write them into the `docs` folder
}

// Opens a specified template based on the name, and return it's contents
func openTemplate(filename string) {
}
