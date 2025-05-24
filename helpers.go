package main

import "github.com/areon546/go-helpers/helpers"

func print(a ...any) {
	helpers.Print(a...)
}

func handle(err error) {
	helpers.Handle(err)
}

func format(s string, a ...any) string {
	return helpers.Format(s, a...)
}
