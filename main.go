package main

import (
	"flag"
	"fmt"
	"math"
	"os"

	"encoding/json"
	"github.com/nilsmagnus/grib/griblib"
	"io"
)

func optionsFromFlag() griblib.Options {
	filename := flag.String("file", "", "Grib filepath")
	reducedFile := flag.String("reducefile", "reduced.grib2", "Destination for reduced file.")
	operation := flag.String("operation", "parse", "Operation. Valid values: 'parse', 'reduce'.")
	exportType := flag.Int("export", griblib.ExportNone, "Export format. Valid types are 0 (none) 1(print discipline names) 2(print categories) 3(json) ")
	maxNum := flag.Int("maxmsg", math.MaxInt32, "Maximum number of messages to parse. Does not work in combination with filters.")
	discipline := flag.Int("discipline", -1, "Filters on Discipline. -1 means all disciplines")
	category := flag.Int("category", -1, "Filters on Category within discipline. -1 means all categories")
	dataExport := flag.Bool("dataExport", true, "Export data values.")
	latMin := flag.Int("latMin", 0, "Minimum latitude multiplied with 100000.")
	latMax := flag.Int("latMax", 36000000, "Maximum latitude multiplied with 100000.")
	longMin := flag.Int("longMin", -9000000, "Minimum longitude multiplied with 100000.")
	longMax := flag.Int("longMax", 9000000, "Maximum longitude multiplied with 100000.")

	flag.Parse()

	return griblib.Options{
		Operation:               string(*operation),
		Filepath:                string(*filename),
		ReduceFilePath:          string(*reducedFile),
		ExportType:              int(*exportType),
		MaximumNumberOfMessages: int(*maxNum),
		Discipline:              int(*discipline),
		Category:                int(*category),
		DataExport:              bool(*dataExport),
		GeoFilter: griblib.GeoFilter{
			MinLat:  int32(*latMin),
			MaxLat:  int32(*latMax),
			MinLong: int32(*longMin),
			MaxLong: int32(*longMax),
		},
	}
}

func main() {
	options := optionsFromFlag()

	if _, err := json.Marshal(options); err == nil {
		//fmt.Printf("Input parameters : %s \n", string(js))
	}

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

	switch options.Operation {
	case "parse":
		parse(gribFile, options)
	case "reduce":
		reduceToFile(gribFile, options)
	default:
		fmt.Printf("Operation '%s' not supported. Valid values are 'parse' and 'reduce'.", options.Operation)
		os.Exit(1)
	}
}

func reduceToFile(gribFile io.ReadSeeker, options griblib.Options) {
	if options.Discipline == -1 {
		fmt.Println("No discipline defined.")
		flag.Usage()
		os.Exit(0)
	}

	reduceFile, err := os.Create(options.ReduceFilePath)
	if err != nil {
		fmt.Printf("Error creating reduced reduceFile: %s", err.Error())
		os.Exit(1)
	}

	defer reduceFile.Close()

	end := make(chan bool)
	content := make(chan []byte)

	go griblib.Reduce(gribFile, options, content, end)

	for {
		select {
		case <-end:
			fmt.Printf("reduce done to file '%s'. \n", options.ReduceFilePath)
			return
		case bytesRead := <-content:
			reduceFile.Write(bytesRead)
		}
	}

}

func parse(gribFile io.Reader, options griblib.Options) {
	messages, err := griblib.ReadMessages(gribFile, options)

	if err != nil {
		fmt.Printf("Error reading all messages in gribfile: %s", err.Error())
		os.Exit(1)
	}

	filteredMessages := griblib.Filter(messages, options)

	griblib.Export(filteredMessages, options)
}
