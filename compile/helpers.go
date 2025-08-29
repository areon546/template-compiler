package compile

import (
	"fmt"

	"github.com/areon546/go-helpers/helpers"
)

func format(s string, a ...any) string {
	return helpers.Format(s, a...)
}

func handle(err error, msg string) {
	fmt.Println(msg)
	helpers.Handle(err)
}

func OutputPath(internalPathToFile string) string {
	return directoryRoots["output"] + "/" + internalPathToFile
}

func ContentPath(internalPathToFile string) string {
	return directoryRoots["content"] + "/" + internalPathToFile
}
