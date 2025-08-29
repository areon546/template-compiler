package compile

import (
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"regexp"
	"strings"
	"template-compiler/compile/dirs"
	"template-compiler/compile/options"

	"github.com/areon546/go-files/files"
)

var (
	ErrIncorrectHandler  error = errors.New("special case: handler given incorrect file type")
	ErrIncorrectFileType error = ErrIncorrectHandler
)

// A template handler will perform a certain actions on a specific type of file.
// A handler returns ErrIncorrectHandler with information about the filename and the check requested by the handler.
type TemplateHandler func(opt options.Options, path, name string) error

type handler struct {
	handle TemplateHandler
	regex  string
	opt    options.Options
}

func NewHandler(handlerFunction TemplateHandler, regex string, opt options.Options) *handler {
	return &handler{handle: handlerFunction, regex: regex, opt: opt}
}

// NOTE: Path needs to ends with a /
func (handler handler) handleFile(intPath string, file fs.DirEntry) (err error) {
	// check if path ends with /
	intPath = dirs.CleanPath(intPath)

	name := file.Name()
	match, err := checkHandlerMatch(handler.regex, name)
	debugCaseHandler("handler match", file, handler.regex, match)

	if err != nil {
		return err
	}

	// Create missing directories.
	err = dirs.MakeOutputDirectories(handler.opt, intPath)
	if err != nil {
		return err
	}

	return handler.handle(handler.opt, intPath, name)
}

func populateCaseHandlers(opt options.Options) (handlerMap map[string]handler) {
	handlerMap = map[string]handler{}

	handlerMap["html"] = *HandleHTML(opt)
	handlerMap["markdown"] = *HandleMarkdown(opt)
	handlerMap["skipTemplate"] = *HandleTemplateFile(opt)
	handlerMap["static"] = *HandleStaticFile(opt)
	// handlerMap["nothing"] = *HandleNothing(opt)

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

func HandleNothing(opt options.Options) *handler {
	return NewHandler(doNothing, "", opt)
}

func doNothing(opt options.Options, path, name string) error {
	return nil
}

func HandleHTML(opt options.Options) *handler {
	return NewHandler(copyOverFile, "[.]*\\.html", opt)
}

func HandleStaticFile(opt options.Options) *handler {
	allowed := func() string {
		acceptedSuffixes := make([]string, 0)

		acceptedSuffixes = append(acceptedSuffixes, "css")
		acceptedSuffixes = append(acceptedSuffixes, "jpg|jpeg|png|webp")
		acceptedSuffixes = append(acceptedSuffixes, "js")

		return strings.Join(acceptedSuffixes, "|")
	}

	anyFileName := "[.]*\\."
	regex := anyFileName + allowed()
	return NewHandler(copyOverFile, regex, opt)
}

// Made for the case of having all of your files in the content directory for ease of access
func copyOverFile(opt options.Options, path, name string) error {
	// TODO: make it copy over any file

	debugCaseHandler("\n copying over file", path, name)
	defer debugCaseHandler("\n finished copying over file", path, name)

	internalPath := path + name
	fmt.Println("Internal Path to ", name, "Path", path)
	openFile, err := files.OpenFile(ContentPath(internalPath))
	if err != nil {
		return err
	}
	debugPrint("path: ", path, ", name: ", name)
	debugPrint("Path to file", internalPath)

	outputFile := CreateOutputFile(internalPath)
	outputFile.ClearFile()

	_, err = outputFile.Write(openFile.Contents())
	return err
}

func HandleMarkdown(opt options.Options) *handler {
	return NewHandler(markdownHandler, "[.]*\\.md|markdown", opt) // [.]*\.md initially
	// [.]*							- match any number of any characters
	// \.  						 	- match a '.'
	// (?=md|markdown)	- lookahead, match either 'md' or 'markdown' - doesn't work properly so this idea is being set off for later
}

// Reads the markdown file and converts it's content to HTML content in memory.
// Create a file at the output directory with the same internal path and
// Then populates the respective template with the content.
func markdownHandler(opt options.Options, path, name string) (err error) {
	debugCaseHandler("\nmarkdown handling ~~~~~~~~~~~~~~~~~~~~~~")
	defer debugCaseHandler("\nend of markdown handling ~~~~~~~~~~~~~~~~~\n")

	// open contents
	internalPathMD := path + name
	contentFile := newContent(internalPathMD)

	internalOutPath, err := replaceExtensionWith(internalPathMD, "md", "html")
	if err != nil {
		return err
	}

	// creating new output file
	fileToWriteTo := CreateOutputFile(internalOutPath)

	// parse template
	templateName := LookupTemplate(opt, path)
	debugCaseHandler("template name", templateName)

	err = insertIntoTemplate(templateName, fileToWriteTo, *contentFile)
	return err
}

func HandleTemplateFile(opt options.Options) *handler {
	return NewHandler(ignoreTemplateHandler, "template."+opt.TemplateSuffix(), opt)
}

// Made for the case of having the content and template directory as the same folder
func ignoreTemplateHandler(opt options.Options, path, name string) (err error) {
	debugPrint(path, name, "being skipped")

	return nil
}

// ~~

func insertIntoTemplate(templateName string, outputFile *files.File, content content) (err error) {
	debugCaseHandler("attempting to insert into: ", outputFile.String())
	defer debugCaseHandler("after template execution\n")
	debugCaseHandler(templateName, outputFile)

	// load template
	tpl, err := template.ParseFiles(templateName)
	debugCaseHandler("template parsed", err)
	if err != nil {
		return err
	}

	// clean outputFile
	err = outputFile.ClearFile()
	debugCaseHandler("file cleared", err)

	err = tpl.Execute(outputFile, template.HTML(content.getHTML()))
	debugCaseHandler("Template inserted", err)

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

func LookupTemplate(opt options.Options, path string) string {
	dirs, _ := files.SplitDirectories(path)
	outPath := opt.Template() + "/"
	debugPrint("LookupTemplate", dirs)

	// dirs is a [] of strings, you can loop through that in a simple way
	for range dirs {

		// join directories
		// check for template
		// check next

		testPath := outPath + strings.Join(dirs, "/") + "/" + "template." + opt.TemplateSuffix()

		// test if this exists
		exists, _ := files.FileExists(testPath)

		debugPrint("LookupTemplate", path, testPath, exists)
		if exists {
			return testPath
		} else {
			dirs = dirs[0 : len(dirs)-1]
		}

	}

	return opt.Template() + "/" + "template." + opt.TemplateSuffix()
}
