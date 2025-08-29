package dirs

import (
	"errors"
	"os"
	"reflect"
	"strings"
	"template-compiler/compile/options"

	"github.com/areon546/go-files/files"
)

var ErrOutIsContOrTemp = errors.New("template-compiler: Cannot have output path be the same as content or template path")

func MakeOutputDirectories(opt options.Options, intPath string) error {
	outPath := CleanPath(opt.Output()) + intPath
	if DirExists(outPath) {
		return nil
	} else {

		dirs, _ := files.SplitDirectories(outPath)

		dirs = files.CleanUpDirs(dirs)
		finalDirPath := strings.Join(dirs, "/") + "/"

		print("Creating directory:", finalDirPath)
		return files.MakeDirectory(finalDirPath)
	}
}

func MakeRelevantDirectories(opt options.Options) (err error) {
	_, err = os.Stat(opt.Content())
	_, err2 := os.Stat(opt.Template())
	err = errors.Join(err, err2)

	return err
}

func RemoveOutputDirectory(opt options.Options) error {
	outIsContOrTemp := reflect.DeepEqual(opt.Output(), opt.Content()) || reflect.DeepEqual(opt.Output(), opt.Template())
	if outIsContOrTemp {
		return ErrOutIsContOrTemp
	}

	files.RemoveAllWithinDirectory(opt.Output())
	return nil
}

func DirExists(dir string) bool {
	exists, _ := files.DirExists(dir)
	return exists
}
