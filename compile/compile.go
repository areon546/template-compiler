package compile

import (
	"errors"
	"fmt"
	"io/fs"
	"template-compiler/compile/dirs"
	"template-compiler/compile/options"

	"github.com/areon546/go-files/files"
)

var (
	// program options
	opt options.Options

	// misc
	templateCases  map[string]handler
	directoryRoots map[string]string
	filesCompiled  = false
)

// ~~~~

func CompileTemplates(options options.Options) {
	opt = options
	fmt.Println(opt)

	templateCases = populateCaseHandlers(opt)
	directoryRoots = map[string]string{"template": opt.Template(), "content": opt.Content(), "output": opt.Output()}

	err := dirs.RemoveOutputDirectory(opt) // Purges output directory to remove artifacts.
	if errors.Is(err, dirs.ErrOutIsContOrTemp) {
		print("The program would otherwise delete the output directory if it went any further. ")
		print(err)
		return
	}

	err = dirs.MakeRelevantDirectories(opt)
	if err != nil {
		print("Some of the directories you referenced do not exist. ")
		print(err)
		return
	}

	// read contents of template and content directories

	// crawl the content files and directories
	path := "./" // records the folder ofset within the content, template, and output directories
	compileTemplatesRec(opt, path)

	if !filesCompiled {
		print("No files compiled. Please try again. ")
	}
}

// Crawls through the contents directory and compiles html files based on that.
func compileTemplatesRec(opt options.Options, path string) {
	templates := files.ReadDirectory(opt.Template() + "/" + path)
	content := files.ReadDirectory(ContentPath(path))

	debugPrint("path", path)
	debugPrint("Templates: ", opt.Template(), templates)
	debugPrint("Content  : ", opt.Content(), content)

	for _, dirEntry := range content {
		debug("entry: ", dirEntry)

		if dirEntry.IsDir() { // if is a directory, then hanlde it recursively
			newPath := path + dirEntry.Name()
			handleSubdirectory(opt, newPath, dirEntry)
		} else if !dirEntry.IsDir() { // if is a file, then check special cases and then handle it
			handleFile(opt, path, dirEntry)
		}

		debugNL()
	}
}

func handleSubdirectory(opt options.Options, path string, directory fs.DirEntry) {
	name := directory.Name()
	path = dirs.CleanPath(path) // NOTE: for some reason, check path makes the logs much shorter, look into
	debug("path: ", path, "directory ", name)

	compileTemplatesRec(opt, path)
}

func handleFile(opt options.Options, path string, file fs.DirEntry) {
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
			handle(err)

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
