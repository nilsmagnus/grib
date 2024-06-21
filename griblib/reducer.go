package griblib

import (
	"bytes"
	"io"
	"log"
	"strings"
)

// Reduce the file in readseeker with the given options, omitting all other products and areas
func Reduce(readSeeker io.Reader, options Options, content chan []byte, end chan bool) {
	if options.Discipline == -1 {
		log.Println("No disciplines defined for reduce.")
		end <- true
		return
	}
	for {
		messageSection0Bytes := make([]byte, 16)
		if section0ByteCount, err := readSeeker.Read(messageSection0Bytes); section0ByteCount == 0 || err != nil {
			log.Println("Read 0 bytes, returning.")
			end <- true
			return
		}
		section0, err := ReadSection0(bytes.NewReader(messageSection0Bytes))

		if err != nil {
			if !isEOF(err) {
				log.Println("section0 read err: ", err.Error())
			}
			end <- true
			return

		}

		if section0.Discipline == uint8(options.Discipline) {
			messageContentBytes := make([]byte, section0.MessageLength-16)
			_, err = readSeeker.Read(messageContentBytes)
			if err != nil {
				log.Printf("read2 err: %v", err.Error())
				end <- true
			}
			content <- messageSection0Bytes
			content <- messageContentBytes
		} else {
			length := int64(section0.MessageLength) - 16
			bytesRead, err2 := readSeeker.Read(make([]byte, length))
			if int64(bytesRead) != length {
				log.Printf("bytesRead: %v, length: %v", bytesRead, length)
			}
			if err2 != nil {
				log.Printf("read3 err: %v", err2.Error())
			}
		}

	}
}

func isEOF(err error) bool {
	log.Printf("Error:: %s", err.Error())
	return strings.Contains(err.Error(), "EOF")
}
