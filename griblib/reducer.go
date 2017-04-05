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
		return
	}
	for {
		messageSection0Bytes := make([]byte, 16)
		if section0ByteCount, err := readSeeker.Read(messageSection0Bytes); section0ByteCount == 0 || err != nil {
			fmt.Println("Read 0 bytes, returning.")
			end <- true
			return
		}
		section0, err := ReadSection0(bytes.NewReader(messageSection0Bytes))

		if err != nil {
			if !isEOF(err) {
				fmt.Println("section0 read err: ", err.Error())
			}
			end <- true
			return

		}

		if section0.Discipline == uint8(options.Discipline) {
			messageContentBytes := make([]byte, section0.MessageLength-16)
			_, err = readSeeker.Read(messageContentBytes)
			if err != nil {
				fmt.Printf("read2 err: ", err.Error())
				end <- true
			}
			content <- messageSection0Bytes
			content <- messageContentBytes
		} else {
			readSeeker.Seek(int64(section0.MessageLength)-16, io.SeekCurrent)
		}

	}
}

func isEOF(err error) bool {
	fmt.Printf("Error:: %s", err.Error())
	return strings.Contains(err.Error(), "EOF")
}
