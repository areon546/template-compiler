package main

import (
	"errors"
	"io/fs"
	"regexp"

	"github.com/areon546/go-files/files"
)

var ErrIncorrectHandler error = errors.New("special case: handler given incorrect file type")

// A template handler will perform a certain actions on a specific type of file.
// A handler returns ErrIncorrectHandler with information about the filename and the check requested by the handler.
type templateHandler func(string, fs.DirEntry)

// NOTE: Path needs to ends with a /
func indexHandler(path string, file fs.DirEntry) error {
	name := file.Name()
	_, err := checkHandlerMatch("index.html", name)
	if errors.As(err, ErrIncorrectHandler) {
		return err
	}

	// copy file to exact relative path in output directory
	pathToFile := "/" + path + name
	openFile := files.OpenFile(directoryRoots["content"] + pathToFile)
	print(openFile, pathToFile)

	fileToWriteTo := files.NewFile(directoryRoots["output"] + pathToFile)

	_, err = fileToWriteTo.Write(openFile.Contents())
	handle(err)

	print("Wrote to file", fileToWriteTo)
	// log(pathToFile)
	return nil
}

// ~~

// Will perform a regex check on the name of a file.
// Returns ErrIncorrectHandler, as well as extra information about the cause of the failure.
func checkHandlerMatch(regex, name string) (matched bool, err error) {
	matched, err = regexp.MatchString(regex, name)

	if !matched {
		errMatch := errors.New(format(": expected a match of: `%s` got: `%s`", regex, name))
		err = errors.Join(ErrIncorrectHandler, errMatch)
	}

	return
}
