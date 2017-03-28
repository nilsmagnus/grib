package data

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Message struct {
	Section1 Section1
	Section2 Section2
	Section3 Section3
	Section4 Section4
	Section5 Section5
	Section6 Section6
	Section7 Section7
}

const (
	GRIB                   = 0x47524942
	SUPPORTED_GRIB_EDITION = 2
)

// ReadMessages reads all message from gribFile
func ReadMessages(gribFile io.Reader) (messages []Message, err error) {

	for {
		message, messageErr := ReadMessage(gribFile)
		if messageErr != nil {
			if messageErr.Error() == "EOF" {
				return
			} else {
				fmt.Println("Error when parsing a message, ", messageErr.Error())
				return messages, err
			}
		} else {
			messages = append(messages, message)
			return messages, err
		}
	}
}

func ReadMessage(gribFile io.Reader) (message Message, err error) {

	section0, headError := ReadSection0(gribFile)

	if headError != nil {
		return message, headError
	}

	messageBytes := make([]byte, section0.MessageLength)

	numBytes, readError := gribFile.Read(messageBytes)

	if readError != nil {
		fmt.Println("Error reading message")
		return message, readError
	}

	if numBytes != int(section0.MessageLength) {
		fmt.Println("Did not read full message")
	}

	return

}

func ReadMessage2(gribFile io.Reader) (message Message, err error) {

	messageHead, headError := ReadSection0(gribFile)

	if headError != nil {
		return message, headError
	}

	fmt.Println("", messageHead)

	for {

		// pre-parse section head to decide which struct use
		var sectionHead SectionHead
		readError := binary.Read(gribFile, binary.BigEndian, &sectionHead)
		if readError != nil {
			return message, readError
		}

		fmt.Println("", sectionHead)
		switch sectionHead.Number {

		case 1:
			message.Section1, err = ReadSection1(gribFile, sectionHead.ByteLength)
		case 2:
			message.Section2, err = ReadSection2(gribFile, sectionHead.ByteLength)
		case 3:
			message.Section3, err = ReadSection3(gribFile, sectionHead.ByteLength)
		case 4:
			message.Section4, err = ReadSection4(gribFile, sectionHead.ByteLength)
		case 5:
			message.Section5, err = ReadSection5(gribFile, sectionHead.ByteLength)
		case 6:
			message.Section6, err = ReadSection6(gribFile, sectionHead.ByteLength)
		case 7:
			message.Section7, err = ReadSection7(gribFile, sectionHead.ByteLength, message.Section5.DataTemplate)
		case 8:
			return message, fmt.Errorf("EOF")
		default:

			err = fmt.Errorf("Unknown section number %d  (Something bad with parser or files)", sectionHead.Number)
		}
		if err != nil {
			return message, err
		}
	}
}

func panicIfNotNil(e error) {
	if e != nil {
		panic(e)
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

	if section0.Indicator == GRIB {
		if section0.Edition != 2 {
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

func (s SectionHead) SectionNumber() uint8 {
	return s.Number
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

func ReadSection1(f io.Reader, len uint32) (section Section1, err error) {
	return section, binary.Read(f, binary.BigEndian, &section)
}

type Section2 struct {
	LocalUse []uint8
}

func ReadSection2(f io.Reader, len uint32) (section Section2, err error) {
	section.LocalUse = make([]uint8, len-5)
	return section, read(f, &section.LocalUse)
}

type Section3 struct {
	Source                   uint8
	DataPointCount           uint32
	PointCountOctets         uint8
	PointCountInterpretation uint8
	TemplateNumber           uint16
	Definition               Grid
}

func (s Section3) String() string {
	return fmt.Sprint("Point count: ", s.DataPointCount, " Definition: ", ReadGridDefinitionTemplateNumber(int(s.TemplateNumber)))
}

func ReadSection3(f io.Reader, len uint32) (section Section3, err error) {

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

func ReadSection4(f io.Reader, length uint32) (section Section4, err error) {
	err = read(f, &section.CoordinatesCount, &section.ProductDefinitionTemplateNumber, &section.ProductDefinitionTemplate)
	if err != nil {
		return section, err
	}

	if section.ProductDefinitionTemplateNumber != 0 {
		return section, fmt.Errorf("Product definition template number %d not implemented yet", section.ProductDefinitionTemplateNumber)
	}

	section.Coordinates = make([]byte, length-9-uint32(binary.Size(section.ProductDefinitionTemplate)))

	return section, read(f, &section.Coordinates)
}

type Section5 struct {
	PointsNumber       uint32
	DataTemplateNumber uint16
	DataTemplate       Data3 // FIXME
}

func ReadSection5(f io.Reader, length uint32) (section Section5, err error) {
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

func ReadSection6(f io.Reader, length uint32) (section Section6, err error) {
	section.Bitmap = make([]byte, length-6)

	return section, read(f, &section.BitmapIndicator, &section.Bitmap)
}

type Section7 struct {
	Data []int64
}

func ReadSection7(f io.Reader, length uint32, template Data3) (section Section7, err error) {
	section.Data = ParseData3(f, length-5, &template) // 5 is the length of (octet 1-5)
	return section, err
}

func read(f io.Reader, data ...interface{}) (err error) {
	for _, what := range data {
		err = binary.Read(f, binary.BigEndian, what)
		if err != nil {
			return err
		}
	}
	return nil
}
