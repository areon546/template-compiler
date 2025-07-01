package compile

import (
	"errors"
	"html/template"
	"io/fs"
	"regexp"
	"strings"

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
	path = checkPath(path)

	name := file.Name()
	match, err := checkHandlerMatch(handler.regex, name)
	debugCaseHandler("handler match", file, handler.regex, match)

	if err != nil {
		return err
	}

	return handler.handle(path, name)
}

func populateCaseHandlers() (handlerMap map[string]handler) {
	handlerMap = map[string]handler{}

	handlerMap["index.html"] = *HandleIndex()
	handlerMap["markdown"] = *HandleMarkdown()
	handlerMap["skipTemplate"] = *HandleTemplateFile()
	handlerMap["static"] = *HandleStaticFile()

	return handlerMap
}

func AddCaseHandler(key string, newHandler handler) {
	templateCases[key] = newHandler
}

func CreateOutputFile(internalPathToFile string) *files.File {
	pathToWriteTo := OutputPath(internalPathToFile)

	return files.NewFile(pathToWriteTo)
}

// ~~~~~~~~~~~~~~~~~~~~~ Handlers

func HandleIndex() *handler {
	return NewHandler(indexHandler, "index.html")
}

func indexHandler(path, name string) error {
	// copy file to exact relative path in output directory
	pathToFile := "/" + path + name
	openFile := files.OpenFile(directoryRoots["content"] + pathToFile)
	debugCaseHandler("index handler: ", openFile, pathToFile)

	outputFile := CreateOutputFile(pathToFile)
	_, err := outputFile.Write(openFile.Contents())
	return err
}

func HandleMarkdown() *handler {
	return NewHandler(markdownHandler, "[.]*\\.md|markdown") // [.]*\.md initially
	// [.]*							- match any number of any characters
	// \.  						 	- match a '.'
	// (?=md|markdown)	- lookahead, match either 'md' or 'markdown' - doesn't work properly so this idea is being set off for later
}

// Reads the markdown file and converts it's content to HTML content in memory.
// Create a file at the output directory with the same internal path and
// Then populates the respective template with the content.
func markdownHandler(path, name string) (err error) {
	debugCaseHandler("\nmarkdown handling ~~~~~~~~~~~~~~~~~~~~~~")
	defer debugCaseHandler("\nend of markdown handling ~~~~~~~~~~~~~~~~~\n")

	// open contents
	internalPathMD := path + name
	contentFile := newContent(internalPathMD)

	internalOutPath, err := replaceMDExtensionWith(internalPathMD, "html")
	if err != nil {
		return err
	}

	// creating new output file
	fileToWriteTo := CreateOutputFile(internalOutPath)

	// parse template
	templateName := LookupTemplate(path)
	debugCaseHandler("template name", templateName)

	err = insertIntoTemplate(templateName, fileToWriteTo, *contentFile)

	print(path, fileToWriteTo, internalOutPath)
	return err
}

func HandleTemplateFile() *handler {
	return NewHandler(ignoreTemplateHandler, "template."+templateFileType)
}

// Made for the case of having the content and template directory as the same folder
func ignoreTemplateHandler(path, name string) (err error) {
	debugPrint(path, name, "being skipped")

	return nil
}

func HandleStaticFile() *handler {
	allowed := func() string {
		acceptedSuffixes := make([]string, 0)

		acceptedSuffixes = append(acceptedSuffixes, "css")
		acceptedSuffixes = append(acceptedSuffixes, "jpg|jpeg|png|webp")
		acceptedSuffixes = append(acceptedSuffixes, "js")

		return strings.Join(acceptedSuffixes, "|")
	}

	anyFileName := "[.]*\\."
	regex := anyFileName + allowed()
	return NewHandler(copyOverFile, regex)
}

// Made for the case of having all of your files in the content directory for ease of access
func copyOverFile(path, name string) (err error) {
	// TODO: make it copy over any file

	debugCaseHandler("\n copying over file", path, name)
	defer debugCaseHandler("\n finished copying over file", path, name)

	debugPrint("path: ", path, ", name: ", name)

	internalPath := path + name
	debugPrint("Path to file", internalPath)

	// read content file
	contentFile := files.OpenFile(ContentPath(internalPath))
	fileContents := contentFile.Contents()

	outputFile := files.NewFile(OutputPath(internalPath))
	outputFile.Append(fileContents)
	outputFile.Close()

	return nil
}

// ~~

func insertIntoTemplate(templateName string, outputFile *files.File, content content) (err error) {
	debugCaseHandler("\n			attempting to insert into: ", outputFile.String())
	defer debugCaseHandler("			after template execution\n")
	debugCaseHandler(templateName, outputFile, content)

	tpl, err := template.ParseFiles(templateName)
	debugCaseHandler("template parsed", err)
	if err != nil {
		return err
	}

	outputFile.ClearFile()
	debugCaseHandler("file cleared")

	debugCaseHandler("before template execution")
	debugCaseHandler("html", content.getHTML())
	err = tpl.Execute(outputFile, template.HTML(content.getHTML()))
	debugCaseHandler(outputFile, outputFile.Contents())
	defer debugCaseHandler("after exe")

	return
}

// Will perform a regex check on the name of a file.
// Returns ErrIncorrectHandler, as well as extra information about the cause of the failure.
func checkHandlerMatch(regex, name string) (matched bool, err error) {
	debugCaseHandler("regex", regex, "name", name)
	matched, err = regexp.MatchString(regex, name)

	if !matched {
		errMatch := errors.New(format(": expected a match of: `%s` got: `%s`", regex, name))
		err = errors.Join(ErrIncorrectHandler, errMatch)
	}

	return
}

func LookupTemplate(path string) string {
	dirs, _ := files.SplitDirectories(path)
	outPath := templateDir + "/"
	debugPrint("LookupTemplate", dirs)

	// dirs is a [] of strings, you can loop through that in a simple way
	for range dirs {

		// join directories
		// check for template
		// check next

		testPath := outPath + strings.Join(dirs, "/") + "/" + "template." + templateFileType

		// test if this exists
		exists, _ := files.FileExists(testPath)

		debugPrint("LookupTemplate", path, testPath, exists)
		if exists {
			return testPath
		} else {
			dirs = dirs[0 : len(dirs)-1]
		}

	}

	return templateDir + "/" + "template." + templateFileType
}
