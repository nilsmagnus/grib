package griblib

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func Reduce(readSeeker io.ReadSeeker, options Options, content chan []byte, end chan bool) {

	for {
		section0Bytes := make([]byte, 16)
		readSeeker.Read(section0Bytes)
		section0, err := ReadSection0(bytes.NewReader(section0Bytes))

		if err != nil {
			fmt.Println("got eof:", err.Error())
			end <- isEOF(err)
			return
		}

		if section0.Discipline == uint8(options.Discipline) {
			fmt.Printf("Found discipline %d \n", options.Discipline)
			readSeeker.Seek(-16, io.SeekCurrent)
			messsageBytes := make([]byte, section0.MessageLength)
			readSeeker.Read(messsageBytes)
			content <- messsageBytes
		} else {
			fmt.Printf("Skipping discipline %d \n", options.Discipline)
			readSeeker.Seek(16+int64(section0.MessageLength), io.SeekCurrent)
		}

	}
}

func isEOF(err error) bool {
	return strings.Contains(err.Error(), "EOF")
}
