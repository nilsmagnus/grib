package main

import (
	"flag"
	"fmt"
	"math"
	"os"

	"github.com/nilsmagnus/grib/griblib"
)

func optionsFromFlag() griblib.Options {
	filename := flag.String("file", "", "Grib filepath")
	exportType := flag.Int("export", griblib.ExportNone, "Export format. Valid types are 0 (none) 1 (json) ")
	maxNum := flag.Int("maxmsg", math.MaxInt32, "Maximum number of messages to parse.")
	discipline := flag.Int("discipline", -1, "Discipline. -1 means all disciplines")     // metereology ==0
	category := flag.Int("category", -1, "Category within discipline. -1 means all categories")   // temperature == 6

	flag.Parse()

	return griblib.Options{
		Filepath:                *filename,
		ExportType:              *exportType,
		MaximumNumberOfMessages: *maxNum,
		Discipline:              *discipline,
		Category:                *category,
	}
}

func main() {
	options := optionsFromFlag()

	fmt.Println(options)
	if options.Filepath == "" {
		fmt.Println("Missing 'file' option. ")
		flag.Usage()
		os.Exit(0)
	}

	gribFile, err := os.Open(options.Filepath)

	if err != nil {
		fmt.Printf("\nFile [%s] not found.\n", options.Filepath)
		os.Exit(1)
	}
	defer gribFile.Close()

	messages, err := griblib.ReadMessages(gribFile, options)

	if err != nil {
		fmt.Printf("Error reading all messages in gribfile: %s", err.Error())
		os.Exit(1)
	}

	filteredMessages := griblib.Filter(messages, options)

	griblib.Export(filteredMessages, options)
}
