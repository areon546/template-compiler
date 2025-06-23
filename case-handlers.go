package main

import (
	"errors"
	"html/template"
	"io/fs"
	"regexp"

	"github.com/areon546/go-files/files"
)

var (
	ErrIncorrectHandler  error = errors.New("special case: handler given incorrect file type")
	ErrIncorrectFileType error = ErrIncorrectHandler
)

// A template handler will perform a certain actions on a specific type of file.
// A handler returns ErrIncorrectHandler with information about the filename and the check requested by the handler.
type TemplateHandler func(path, name string) error

type handler struct {
	handle TemplateHandler
	regex  string
}

func NewHandler(handlerFunction TemplateHandler, regex string) *handler {
	return &handler{handle: handlerFunction, regex: regex}
}

// NOTE: Path needs to ends with a /
func (handler handler) handleFile(path string, file fs.DirEntry) (err error) {
	// check if path ends with /
	if path[len(path)-1] != '/' {
		path += "/"
	}

	name := file.Name()
	match, err := checkHandlerMatch(handler.regex, name)
	debug("handler match", file, handler.regex, match)

	if err != nil {
		return err
	}

	return handler.handle(path, name)
}

func populateCaseHandlers() (handlerMap map[string]handler) {
	handlerMap = map[string]handler{}

	handlerMap["index.html"] = *NewHandler(indexHandler, "index.html")

	handlerMap["markdown"] = *NewHandler(markdownHandler, "[.]*\\.md") // [.]*\.md initially

	handlerMap["skipTemplate"] = *NewHandler(templateHandler, "template."+templateFileType)

	return handlerMap
}

func AddCaseHandler(key string, newHandler handler) {
	templateCases[key] = newHandler
}

func writeToOutputPath(out *files.File, content []byte) (err error) {
	_, err = out.Write(content)
	if err != nil {
		debug("Wrote to file: ", out)
	}

	return
}

func CreateOutputFile(internalPathToFile string) (out *files.File) {
	pathToWriteTo := directoryRoots["output"] + internalPathToFile

	out = files.NewFile(pathToWriteTo)

	return
}

// ~~~~~~~~~~~~~~~~~~~~~ Handlers

func indexHandler(path, name string) error {
	// copy file to exact relative path in output directory
	pathToFile := "/" + path + name
	openFile := files.OpenFile(directoryRoots["content"] + pathToFile)
	debug("index handler: ", openFile, pathToFile)

	outputFile := CreateOutputFile(pathToFile)
	_, err := outputFile.Write(openFile.Contents())
	return err
}

// Reads the markdown file and converts it's content to HTML content in memory.
// Create a file at the output directory with the same internal path and
// Then populates the respective template with the content.
func markdownHandler(path, name string) (err error) {
	debug("\nmarkdown handling ~~~~~~~~~~~~~~~~~~~~~~")
	defer debug("\nend of markdown handling ~~~~~~~~~~~~~~~~~\n")

	// open contents
	contentFile := newContent(path, name)

	newName, err := replaceMDExtensionWith(name, "html")
	if err != nil {
		return err
	}

	// creating new output file
	pathToFile := "/" + path + newName
	fileToWriteTo := CreateOutputFile(pathToFile)

	// parse template
	templateName := templateDir + "/" + path + "template." + templateFileType
	debug("template name", templateName)

	err = insertIntoTemplate(templateName, fileToWriteTo, *contentFile)
	return err
}

func templateHandler(path, name string) (err error) {
	print(path, name, "being skipped")

	return nil
}

// ~~

func insertIntoTemplate(templateName string, outputFile *files.File, content content) (err error) {
	debug("\n			attempting to insert into: ", outputFile.String())
	defer debug("			after template execution\n")
	debug(templateName, outputFile, content)

	tpl, err := template.ParseFiles(templateName)
	debug("template parsed", err)
	if err != nil {
		return err
	}

	outputFile.ClearFile()
	debug("file cleared")

	debug("before template execution")
	debug("html", content.getHTML())
	err = tpl.Execute(outputFile, template.HTML(content.getHTML()))
	debug(outputFile, outputFile.Contents())
	defer debug("after exe")

	return
}

// Will perform a regex check on the name of a file.
// Returns ErrIncorrectHandler, as well as extra information about the cause of the failure.
func checkHandlerMatch(regex, name string) (matched bool, err error) {
	debug("regex", regex, "name", name)
	matched, err = regexp.MatchString(regex, name)

	if !matched {
		errMatch := errors.New(format(": expected a match of: `%s` got: `%s`", regex, name))
		err = errors.Join(ErrIncorrectHandler, errMatch)
	}

	return
}
