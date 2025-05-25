package main

import (
	"errors"
	"io/fs"

	"github.com/areon546/go-files/files"
	"github.com/areon546/go-helpers/helpers"
)

var (
	templateCases  map[string]handler = populateCaseHandlers()
	directoryRoots map[string]string  = map[string]string{"template": templateDir, "content": contentDir, "output": outputDir}
)

// ~~~~

func CompileTemplates() {
	// read contents of template and content directories

	// crawl the content files and directories
	path := "./" // records the folder ofset within the content, template, and output directories
	compileTemplatesRec(path)
}

// Crawls through the contents directory and compiles html files based on that.
func compileTemplatesRec(path string) {
	templates := files.ReadDirectory(templateDir + "/" + path)
	content := files.ReadDirectory(contentDir + "/" + path)

	print()
	print("path", path)
	print(templateDir, templates)

	print(contentDir, content)

	print()

	for _, dirEntry := range content {
		print()
		print("entry: ", dirEntry)

		if dirEntry.IsDir() { // if is a directory, then hanlde it recursively
			newPath := path + dirEntry.Name()
			handleSubdirectory(newPath, dirEntry)
		} else if !dirEntry.IsDir() { // if is a file, then check special cases and then handle it
			handleFile(path, dirEntry)
		}
	}
}

func handleSubdirectory(path string, directory fs.DirEntry) {
	name := directory.Name()

	print("path: ", path)
	print("directory ", name)
	print()
	compileTemplatesRec(path)
}

func handleFile(path string, file fs.DirEntry) {
	name := file.Name()

	print("path: ", path)
	print("file: ", name)

	// since it is a file, we want to:
	// check for special cases

	caseHandler, ok := templateCases[name]
	// if the key
	for _, handler := range templateCases {
		err := handler.handleFile(path, file)
		if !errors.Is(err, ErrIncorrectHandler) {
			helpers.HandleExcept(err, ErrIncorrectHandler)
		} else {
			continue
		}
	}

	// If the key refers to a special case:
	if ok {
		// run it's special case handler
		print("special case", name)
		err := caseHandler.handleFile(path, file)

		handle(err)
	} else {
		// compile it's template, and then write it to the output directory
		compileFile(path, file)
	}
}

func compileFile(path string, file fs.DirEntry) {
}
