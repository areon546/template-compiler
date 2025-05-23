package main

import (
	"io/fs"

	"github.com/areon546/go-files/files"
	"github.com/areon546/go-helpers/helpers"
)

func CompileTemplates(contentDir, templateDir string) {
	// read contents of template and content directories

	// crawl the content files and directories
	compileTemplatesRec(contentDir, templateDir)
}

// Crawls through the contents directory and compiles html files based on that.
func compileTemplatesRec(contentDir, templateDir string) {
	templates := files.ReadDirectory(templateDir)
	content := files.ReadDirectory(contentDir)

	print(templateDir)
	print(templates)

	print(contentDir)
	print(content)

	for _, dirEntry := range content {
		print("entry: ", dirEntry)

		if dirEntry.IsDir() {
			// if is a directory, then hanlde it recursively
			handleSubdirectory(contentDir, templateDir, dirEntry)
		} else if !dirEntry.IsDir() {
			// if is a file, then handle it
			handleFile(content, templates, dirEntry)
		}
	}
}

func handleSubdirectory(contentDir, templateDir string, directory fs.DirEntry) {
	name := directory.Name()
	newContent := contentDir
	newTemplates := templateDir

	print("directory ", name, newContent, newTemplates)
	// compileTemplatesRec(newContent, newTemplates)
}

func handleFile(content, templates []fs.DirEntry, file fs.DirEntry) {
	print("file ", file.Name())
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
