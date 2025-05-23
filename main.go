package main

import (
	"github.com/areon546/go-files/files"
	"github.com/areon546/go-helpers/helpers"
)

type Library struct {
	Message string
}

func main() {
	print(("." + "asd"))

	readCSV()
}

// Reads the rankings csv file
func readCSV() {
	// This will read the csv rankings.csv and write them to the sqlite database.
	rankings := files.ReadCSV("../rankings")

	helpers.Print(rankings.GetContents())
}

func processTemplates() {
	// This will read the templates folder, execute them, and write them into the `docs` folder
}

// Opens a specified template based on the name, and return it's contents
func openTemplate(filename string) {

}
