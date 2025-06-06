package main

import (
	"errors"
	"io/fs"

	"github.com/areon546/go-files/files"
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
	debug("path", path)
	debug(templateDir, templates)
	debug(contentDir, content)

	for _, dirEntry := range content {
		debug("entry: ", dirEntry)

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
	debug("path: ", path, "directory ", name)

	compileTemplatesRec(path)
}

func handleFile(path string, file fs.DirEntry) {
	name := file.Name()

	debug("path: ", path, "file: ", name)

	// if the key, perform the related action
	for key, handler := range templateCases {
		debug("trying ", key)
		err := handler.handleFile(path, file)

		incorrectHandler := errors.Is(err, ErrIncorrectHandler)
		if incorrectHandler {
			print(path, name, "unsuccessfully used ", key)
		} else if !incorrectHandler {
			handle(err, "incorrect handlers")

			print(path, name, "successfuly used ", key)
			print()
			break
		}
	}
}
