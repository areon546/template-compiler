package main

import "github.com/areon546/go-files/log"

var (
	debugger log.LogOutput = log.NewFileLogger("attemptLog.lg")
	output   log.LogOutput = log.NewPrintLogger()
)

func debug(a ...any) {
	debugger.Output(a...)
}

func print(a ...any) {
	debug(a...)
	output.Output(a...)
}
