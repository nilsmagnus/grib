package main

import (
	"flag"
	"fmt"
	"math"
	"os"

	"encoding/json"
	"github.com/nilsmagnus/grib/griblib"
)

func optionsFromFlag() griblib.Options {
	filename := flag.String("file", "", "Grib filepath")
	exportType := flag.Int("export", griblib.ExportNone, "Export format. Valid types are 0 (none) 1(print discipline names) 2 (json) ")
	maxNum := flag.Int("maxmsg", math.MaxInt32, "Maximum number of messages to parse. Does not work in combination with filters.")
	discipline := flag.Int("discipline", -1, "Filters on Discipline. -1 means all disciplines")
	category := flag.Int("category", -1, "Filters on Category within discipline. -1 means all categories")
	dataExport := flag.Bool("dataExport", true, "Export data values.")
	latMin := flag.Int("latMin", 0, "Minimum latitude.")
	latMax := flag.Int("latMax", 36000000, "Maximum latitude.")
	longMin := flag.Int("longMin", -9000000, "Minimum longitude.")
	longMax := flag.Int("longMax", 9000000, "Maximum longitude.")

	flag.Parse()

	return griblib.Options{
		Filepath:                string(*filename),
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

	if js, err := json.Marshal(options); err == nil {
		fmt.Printf("Input parameters : %s", string(js))
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

	messages, err := griblib.ReadMessages(gribFile, options)

	if err != nil {
		fmt.Printf("Error reading all messages in gribfile: %s", err.Error())
		os.Exit(1)
	}

	filteredMessages := griblib.Filter(messages, options)

	griblib.Export(filteredMessages, options)
}
