package main

import (
	"github.com/areon546/go-files/log"
)

var (
	debugger log.LogOutput = log.NewFileLogger("attemptLog.lg")
	output   log.LogOutput = log.NewPrintLogger()
)

func debug(a ...any) {
	debugger.Output(a...)

	// fmt.Println(a...)

	debugger.Close()
}

func print(a ...any) {
	debugPrint(a...)
	output.Output(a...)
}

func closeLoggers() {
	debugger.Close()
	output.Close()
}

// ~~~~~~ Debug Functions

func prepend(el any, a ...any) []any {
	preffix := []any{el}

	return append(preffix, a...)
}

func debugCaseHandler(a ...any) {
	prfx := "case handler: "
	a = prepend(prfx, a...)
	debug(a...)
}

func debugPrint(a ...any) {
	prfx := "print: "
	a = prepend(prfx, a...)
	debug(a...)
}

func debugNL() {
	debug("")
}
