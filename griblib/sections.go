package griblib

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strings"
	"unsafe"
)

type Message struct {
	Section0 Section0
	Section1 Section1
	Section2 Section2
	Section3 Section3
	Section4 Section4
	Section5 Section5
	Section6 Section6
	Section7 Section7
}

type Options struct {
	Discipline              int
	DataExport              bool
	Category                int
	Filepath                string
	ExportType              int
	MaximumNumberOfMessages int
	GeoFilter               GeoFilter
}

const (
	Grib                 = 0x47524942
	EndSectionLength     = 926365495
	SupportedGribEdition = 2
)

// ReadMessages reads all message from gribFile
func ReadMessages(gribFile io.Reader, options Options) (messages []Message, err error) {

	for {
		message, messageErr := ReadMessage(gribFile)
		if messageErr != nil {
			if strings.Contains(messageErr.Error(), "EOF") {
				return messages, nil
			} else {
				fmt.Println("Error when parsing a message, ", messageErr.Error())
				return messages, err
			}
		} else {
			messages = append(messages, message)
			if len(messages) >= int(options.MaximumNumberOfMessages) {
				return messages, nil
			}
		}
	}
}

func ReadMessage(gribFile io.Reader) (message Message, err error) {

	section0, headError := ReadSection0(gribFile)

	if headError != nil {
		return message, headError
	}

	fmt.Sprintln("section 0 length is ", unsafe.Sizeof(section0))
	messageBytes := make([]byte, section0.MessageLength-16)

	numBytes, readError := gribFile.Read(messageBytes)

	if readError != nil {
		fmt.Println("Error reading message")
		return message, readError
	}

	if numBytes != int(section0.MessageLength-16) {
		fmt.Println("Did not read full message")
	}

	return readMessage(bytes.NewReader(messageBytes), section0)

}

func readMessage(gribFile io.Reader, section0 Section0) (message Message, err error) {

	message.Section0 = section0
	for {

		// pre-parse section head to decide which struct use
		sectionHead, headErr := ReadSectionHead(gribFile)
		if headErr != nil {
			fmt.Println("Error reading header", headErr.Error())
			return message, headErr
		}

		switch sectionHead.Number {

		case 1:
			message.Section1, err = ReadSection1(gribFile)
		case 2:
			message.Section2, err = ReadSection2(gribFile, sectionHead.ContentLength())
		case 3:
			message.Section3, err = ReadSection3(gribFile)
		case 4:
			message.Section4, err = ReadSection4(gribFile)
		case 5:
			message.Section5, err = ReadSection5(gribFile)
		case 6:
			message.Section6, err = ReadSection6(gribFile, sectionHead.ContentLength())
		case 7:
			message.Section7, err = ReadSection7(gribFile, sectionHead.ContentLength(), message.Section5.DataTemplate)
		case 8:
			// end-section, return
			return message, nil
		default:
			err = fmt.Errorf("Unknown section number %d  (Something bad with parser or files)", sectionHead.Number)
		}
		if err != nil {
			return message, err
		}
	}
}

type Section0 struct {
	Indicator     uint32
	Reserved      uint16
	Discipline    uint8
	Edition       uint8
	MessageLength uint64
}

func ReadSection0(reader io.Reader) (section0 Section0, err error) {
	err = binary.Read(reader, binary.BigEndian, &section0)
	if err != nil {
		return section0, err
	}

	if section0.Indicator == Grib {
		if section0.Edition != SupportedGribEdition {
			return section0, fmt.Errorf("Unsupported  grib edition %d", section0.Edition)
		}
	}

	return

}

type Section interface {
	SectionNumber() uint8
	String() string
}

type SectionHead struct {
	ByteLength uint32
	Number     uint8
}

func ReadSectionHead(section io.Reader) (head SectionHead, err error) {
	var length uint32
	err = binary.Read(section, binary.BigEndian, &length)
	if err != nil {
		return head, fmt.Errorf("Read of Length failed: %s", err.Error())
	}
	if length == EndSectionLength {
		return SectionHead{
			ByteLength: 4,
			Number:     8,
		}, nil
	}
	var sectionNumber uint8
	err = binary.Read(section, binary.BigEndian, &sectionNumber)
	if err != nil {
		return head, err
	}

	return SectionHead{
		ByteLength: length,
		Number:     sectionNumber,
	}, nil
}

func (s SectionHead) SectionNumber() uint8 {
	return s.Number
}

func (s SectionHead) ContentLength() int {
	return int(s.ByteLength) - binary.Size(s)
}

func (s SectionHead) String() string {
	return fmt.Sprint("Section ", s.Number)
}

type Time struct {
	Year   uint16 // year
	Month  uint8  // month + 1
	Day    uint8  // day
	Hour   uint8  // hour
	Minute uint8  // minute
	Second uint8  // second
}

type Section1 struct {
	OriginatingCenter         uint16
	OriginatingSubCenter      uint16
	MasterTablesVersion       uint8
	LocalTablesVersion        uint8
	ReferenceTimeSignificance uint8
	ReferenceTime             Time
	ProductionStatus          uint8
	Type                      uint8
}

func ReadSection1(f io.Reader) (section Section1, err error) {
	return section, binary.Read(f, binary.BigEndian, &section)
}

type Section2 struct {
	LocalUse []uint8
}

func ReadSection2(f io.Reader, len int) (section Section2, err error) {
	section.LocalUse = make([]uint8, len)
	return section, read(f, &section.LocalUse)
}

type Section3 struct {
	Source                   uint8
	DataPointCount           uint32
	PointCountOctets         uint8
	PointCountInterpretation uint8
	TemplateNumber           uint16
	Definition               interface{}
}

func (s Section3) String() string {
	return fmt.Sprint("Point count: ", s.DataPointCount, " Definition: ", ReadGridDefinitionTemplateNumber(int(s.TemplateNumber)))
}

func ReadSection3(f io.Reader) (section Section3, err error) {

	err = read(f, &section.Source, &section.DataPointCount, &section.PointCountOctets, &section.PointCountInterpretation, &section.TemplateNumber)
	if err != nil {
		return section, err
	}

	section.Definition, err = ReadGrid(f, section.TemplateNumber)
	return section, err
}

type Section4 struct {
	CoordinatesCount                uint16
	ProductDefinitionTemplateNumber uint16
	ProductDefinitionTemplate       Product0 // FIXME
	Coordinates                     []byte
}

func (s Section4) String() string {
	return fmt.Sprint(s.ProductDefinitionTemplate.String())
}

func ReadSection4(f io.Reader) (section Section4, err error) {
	err = read(f, &section.CoordinatesCount, &section.ProductDefinitionTemplateNumber, &section.ProductDefinitionTemplate)
	if err != nil {
		return section, err
	}

	if section.ProductDefinitionTemplateNumber != 0 {
		return section, fmt.Errorf("Category definition template number %d not implemented yet", section.ProductDefinitionTemplateNumber)
	}

	section.Coordinates = make([]byte, section.CoordinatesCount)

	return section, read(f, &section.Coordinates)
}

type Section5 struct {
	PointsNumber       uint32
	DataTemplateNumber uint16
	DataTemplate       Data3 // FIXME
}

func ReadSection5(f io.Reader) (section Section5, err error) {
	err = read(f, &section.PointsNumber, &section.DataTemplateNumber, &section.DataTemplate)
	if err != nil {
		return section, err
	}

	if section.DataTemplateNumber != 3 {
		return section, fmt.Errorf("Template number not supported: %d", section.DataTemplateNumber)
	}
	//f.Seek(int64(length - 11), 1);

	//fmt.Println(section.DataTemplate)

	return section, nil
}

type Section6 struct {
	BitmapIndicator uint8
	Bitmap          []byte
}

func ReadSection6(f io.Reader, length int) (section Section6, err error) {
	section.Bitmap = make([]byte, length-1)

	return section, read(f, &section.BitmapIndicator, &section.Bitmap)
}

type Section7 struct {
	Data []int64
}

func ReadSection7(f io.Reader, length int, template Data3) (section Section7, err error) {
	section.Data = ParseData3(f, length, &template) // 5 is the length of (octet 1-5)
	return section, err
}

func read(reader io.Reader, data ...interface{}) (err error) {
	for _, what := range data {
		err = binary.Read(reader, binary.BigEndian, what)
		if err != nil {
			return err
		}
	}
	return nil
}
