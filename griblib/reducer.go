package griblib

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func Reduce(readSeeker io.ReadSeeker, options Options, content chan []byte, end chan bool) {
	if options.Discipline == -1 {
		fmt.Println("No disciplines defined for reduce.")
		end <- true
	}
	for {
		section0Bytes := make([]byte, 16)
		readSeeker.Read(section0Bytes)
		section0, err := ReadSection0(bytes.NewReader(section0Bytes))

		if err != nil {
			if !isEOF(err) {
				fmt.Println("section0 read err: ", err.Error())
			}
			end <- true
			return

		}

		if section0.Indicator == 0 {
			end <- true
			return
		}

		if section0.Discipline == uint8(options.Discipline) {
			messsage := make([]byte, section0.MessageLength-16)
			_, err = readSeeker.Read(messsage)
			if err != nil {
				fmt.Printf("read2 err: ", err.Error())
				end <- true
			}
			content <- append(section0Bytes, messsage...)
		} else {
			readSeeker.Seek(int64(section0.MessageLength)-16, io.SeekCurrent)
		}

	}
}

func isEOF(err error) bool {
	fmt.Printf("Error:: %s", err.Error())
	return strings.Contains(err.Error(), "EOF")
}
