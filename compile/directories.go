package compile

import (
	"errors"
	"os"
	"strings"

	"github.com/areon546/go-files/files"
)

func MakeOutputDirectories(outPath string) error {
	if dirExists(outPath) {
		return nil
	} else {

		dirs, _ := files.SplitDirectories(outPath)

		dirs = files.CleanUpDirs(dirs)
		finalDirPath := strings.Join(dirs, "/") + "/"

		print("Creating directory:", finalDirPath)
		return files.MakeDirectory(finalDirPath)
	}
}

func makeRelevantDirectories() (err error) {
	_, err = os.Stat(contentDir)
	_, err2 := os.Stat(templateDir)
	err = errors.Join(err, err2)

	return err
}

func dirExists(dir string) bool {
	exists, _ := files.DirExists(dir)
	return exists
}
