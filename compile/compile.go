package compile

import (
	"errors"
	"io/fs"

	"github.com/areon546/go-files/files"
)

var (
	// program options
	templateDir      string
	contentDir       string
	outputDir        string
	templateFileType string
	logFileName      string

	// misc
	templateCases  map[string]handler
	directoryRoots map[string]string
	filesCompiled  = false
)

// ~~~~

func CompileTemplates(templateDirectory, contentDirectory, outputDirectory, templateSuff, logFName string) {
	templateDir = templateDirectory
	contentDir = contentDirectory
	outputDir = outputDirectory
	templateFileType = templateSuff
	logFileName = logFName

	templateCases = populateCaseHandlers()
	directoryRoots = map[string]string{"template": templateDir, "content": contentDir, "output": outputDir}

	// read contents of template and content directories

	// crawl the content files and directories
	path := "./" // records the folder ofset within the content, template, and output directories
	compileTemplatesRec(path)

	if !filesCompiled {
		print("No files compiled. Please try again. ")
	}
}

// Crawls through the contents directory and compiles html files based on that.
func compileTemplatesRec(path string) {
	templates := files.ReadDirectory(templateDir + "/" + path)
	content := files.ReadDirectory(ContentPath(path))

	debugPrint("path", path)
	debugPrint("Templates: ", templateDir, templates)
	debugPrint("Content  : ", contentDir, content)

	for _, dirEntry := range content {
		debug("entry: ", dirEntry)

		if dirEntry.IsDir() { // if is a directory, then hanlde it recursively
			newPath := path + dirEntry.Name()
			handleSubdirectory(newPath, dirEntry)
		} else if !dirEntry.IsDir() { // if is a file, then check special cases and then handle it
			handleFile(path, dirEntry)
		}

		debugNL()
	}
}

func handleSubdirectory(path string, directory fs.DirEntry) {
	name := directory.Name()
	path = checkPath(path) // NOTE: for some reason, check path makes the logs much shorter, look into
	debug("path: ", path, "directory ", name)

	compileTemplatesRec(path)
}

func handleFile(path string, file fs.DirEntry) {
	name := file.Name()
	handlerUsed := false
	fileName := path + name

	debug("path: ", path, "file: ", name)

	// if the key, perform the related action
	for key, handler := range templateCases {
		debug("trying ", key)
		err := handler.handleFile(path, file)

		incorrectHandler := errors.Is(err, ErrIncorrectHandler)
		if !incorrectHandler {
			handle(err, "incorrect handlers")

			print(fileName, "successfuly used ", key)
			handlerUsed = true
			filesCompiled = true
			break

		} else {
			debugPrint(fileName, "unsuccessfully used ", key)
		}
	}

	if !handlerUsed {
		print(fileName, "unsuccessfully compiled ")
	}
}
