package main

import (
	"flag"
	"log"
	"math"
	"os"

	"github.com/nilsmagnus/grib/griblib"
	"io"
)

func optionsFromFlag() griblib.Options {
	filename := flag.String("file", "", "Grib filepath")
	reducedFile := flag.String("reducefile", "reduced.grib2", "Destination for reduced file.")
	operation := flag.String("operation", "parse", "Operation. Valid values: 'parse', 'reduce'.")
	exportType := flag.Int("export", griblib.ExportNone, "Export format. Valid types are 0 (none) 1(print discipline names) 2(print categories) 3(json) 4(png - experimental) ")
	maxNum := flag.Int("maxmsg", math.MaxInt32, "Maximum number of messages to parse. Does not work in combination with filters.")
	discipline := flag.Int("discipline", -1, "Filters on Discipline. -1 means all disciplines")
	category := flag.Int("category", -1, "Filters on Category within discipline. -1 means all categories")
	dataExport := flag.Bool("dataExport", true, "Export data values.")
	surface := flag.Int("surfacetype", 255, "Surface type (1== ground/sea level)")
	latMin := flag.Int("latMin", griblib.LatitudeSouth, "Minimum latitude multiplied with 100000.")
	latMax := flag.Int("latMax", griblib.LatitudeNorth, "Maximum latitude multiplied with 100000.")
	longMin := flag.Int("longMin", griblib.LongitudeStart, "Minimum longitude multiplied with 100000.")
	longMax := flag.Int("longMax", griblib.LongitudeEnd, "Maximum longitude multiplied with 100000.")

	flag.Parse()

	return griblib.Options{
		Operation:               *operation,
		Filepath:                *filename,
		ReduceFilePath:          *reducedFile,
		ExportType:              *exportType,
		MaximumNumberOfMessages: *maxNum,
		Discipline:              *discipline,
		Category:                *category,
		DataExport:              *dataExport,
		Surface: griblib.Surface{
			Type: uint8(*surface),
		},
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

	log.Printf("Input parameters : %#v \n", options)

	if options.Filepath == "" {
		log.Println("Missing 'file' option. ")
		flag.Usage()
		os.Exit(0)
	}

	gribFile, err := os.Open(options.Filepath)

	if err != nil {
		log.Printf("\nFile [%s] not found.\n", options.Filepath)
		os.Exit(1)
	}
	defer gribFile.Close()

	switch options.Operation {
	case "parse":
		parse(gribFile, options)
	case "reduce":
		reduceToFile(gribFile, options)
	default:
		log.Printf("Operation '%s' not supported. Valid values are 'parse' and 'reduce'.", options.Operation)
		os.Exit(1)
	}
}

func reduceToFile(gribFile io.Reader, options griblib.Options) {
	if options.Discipline == -1 {
		log.Println("No discipline defined.")
		flag.Usage()
		os.Exit(0)
	}

	reduceFile, err := os.Create(options.ReduceFilePath)
	if err != nil {
		log.Printf("Error creating reduced reduceFile: %s", err.Error())
		os.Exit(1)
	}

	defer reduceFile.Close()

	end := make(chan bool)
	content := make(chan []byte)

	go griblib.Reduce(gribFile, options, content, end)

	for {
		select {
		case <-end:
			log.Printf("reduce done to file '%s'. \n", options.ReduceFilePath)
			return
		case bytesRead := <-content:
			reduceFile.Write(bytesRead)
		}
	}

}

func parse(gribFile io.Reader, options griblib.Options) {
	messages, err := griblib.ReadMessages(gribFile)

	if err != nil {
		log.Printf("Error reading all messages in gribfile: %s", err.Error())
		os.Exit(1)
	}

	filteredMessages := griblib.Filter(messages, options)

	griblib.Export(filteredMessages, options)
}
