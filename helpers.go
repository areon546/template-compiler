package main

import (
	"github.com/areon546/go-helpers/helpers"
)

func format(s string, a ...any) string {
	return helpers.Format(s, a...)
}

func handle(err error, msg string) {
	helpers.Handle(err)
}

func checkPath(path string) string {
	if path[len(path)-1] != '/' {
		path += "/"
	}

	return path
}
