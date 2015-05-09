package data

import (
	"encoding/binary"
	"fmt"
	"os"
    "errors"
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

func ReadMessage(f *os.File) (message Message, err error) {

    //
    // find header
    //
    var header Head
    found := false
    for offset := int64(0); true; offset++ {
        f.Seek(offset, 1)
        err = binary.Read(f, binary.BigEndian, &header)
        if err != nil {
            return message, err
        }

        if header.Indicator == 0x47524942 { // GRIB in ascii
            if header.Edition != 2 {
                return message, errors.New(fmt.Sprintf("Unknown edition %d", header.Edition))
            }
            found = true
            break
        }
    }

    if !found {
        return message, errors.New("Head not found")
    }


	for {
		position, err := f.Seek(0, 1)
		check(err)

		// pre-parse section head to decide which struct use
		var len uint32
		err = binary.Read(f, binary.BigEndian, &len)
		check(err)

		if len == 926365495 /* 7777 */ {
			return message, nil //errors.New("EOF")
		}

		var num uint8
		err = binary.Read(f, binary.BigEndian, &num)
		check(err)

		f.Seek(position, 0)

		switch num {
		case 1:
			message.Section1, err = ReadSection1(f, len)
		case 2:
			message.Section2, err = ReadSection2(f, len)
		case 3:
			message.Section3, err = ReadSection3(f, len)
		case 4:
			message.Section4, err = ReadSection4(f, len)
		case 5:
			message.Section5, err = ReadSection5(f, len)
		case 6:
			message.Section6, err = ReadSection6(f, len)
		case 7:
			message.Section7, err = ReadSection7(f, len, message.Section5.DataTemplate)
		default:
			panic(fmt.Sprint("Wrong section number ", num, " (Something bad with parser or files)"))
		}
		check(err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Head struct {
	Indicator  uint32
	Reserved1  uint8
	Reserved2  uint8
	Discipline uint8
	Edition    uint8
	ByteLength uint64
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
	SectionHead
	OriginatingCenter         uint16
	OriginatingSubCenter      uint16
	MasterTablesVersion       uint8
	LocalTablesVersion        uint8
	ReferenceTimeSignificance uint8
	ReferenceTime             Time
	ProductionStatus          uint8
	Type                      uint8
}

func ReadSection1(f *os.File, len uint32) (section Section1, err error) {
	return section, binary.Read(f, binary.BigEndian, &section)
}

type Section2 struct {
	SectionHead
	LocalUse []uint8
}

func ReadSection2(f *os.File, len uint32) (section Section2, err error) {
	section.LocalUse = make([]uint8, len-5)
	return section, read(f, &section.ByteLength, &section.Number, &section.LocalUse)
}

type Section3 struct {
	SectionHead
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

func ReadSection3(f *os.File, len uint32) (section Section3, err error) {

	err = read(f, &section.ByteLength, &section.Number, &section.Source, &section.DataPointCount, &section.PointCountOctets, &section.PointCountInterpretation, &section.TemplateNumber)
	if err != nil {
		return section, err
	}

	section.Definition, err = ReadGrid(f, section.TemplateNumber)
	return section, err
}

type Section4 struct {
	SectionHead
	CoordinatesCount                uint16
	ProductDefinitionTemplateNumber uint16
	ProductDefinitionTemplate       Product0 // FIXME
	Coordinates                     []byte
}

func (s Section4) String() string {
	return fmt.Sprint(s.ProductDefinitionTemplate.String())
}

func ReadSection4(f *os.File, length uint32) (section Section4, err error) {
	err = read(f, &section.ByteLength, &section.Number, &section.CoordinatesCount, &section.ProductDefinitionTemplateNumber, &section.ProductDefinitionTemplate)
	if err != nil {
		return section, err
	}

	if section.ProductDefinitionTemplateNumber != 0 {
		// TODO
		panic(fmt.Sprint("Product definition template number", section.ProductDefinitionTemplateNumber, "not implemented yet"))
	}

	section.Coordinates = make([]byte, length-9-uint32(binary.Size(section.ProductDefinitionTemplate)))

	return section, read(f, &section.Coordinates)
}

type Section5 struct {
	SectionHead
	PointsNumber       uint32
	DataTemplateNumber uint16
	DataTemplate       Data3 // FIXME
}

func ReadSection5(f *os.File, length uint32) (section Section5, err error) {
	err = read(f, &section.ByteLength, &section.Number, &section.PointsNumber, &section.DataTemplateNumber, &section.DataTemplate)
	if err != nil {
		return section, err
	}

	if section.DataTemplateNumber != 3 {
		// TODO
		panic(fmt.Sprint("Data template number", section.DataTemplateNumber, "not implemented yet"))
	}
	//f.Seek(int64(length - 11), 1);

	//fmt.Println(section.DataTemplate)

	return section, nil
}

type Section6 struct {
	SectionHead
	BitmapIndicator uint8
	Bitmap          []byte
}

func ReadSection6(f *os.File, length uint32) (section Section6, err error) {
	section.Bitmap = make([]byte, length-6)

	return section, read(f, &section.ByteLength, &section.Number, &section.BitmapIndicator, &section.Bitmap)
}

type Section7 struct {
	SectionHead
	RawData []byte
    Data []int64
}

func ReadSection7(f *os.File, length uint32, template Data3) (section Section7, err error) {
	//dataLength := length - 5
	//section.RawData = make([]byte, dataLength)

	err = read(f, &section.ByteLength, &section.Number) // , &section.RawData)
	if err != nil {
		return section, err
	}

	//fmt.Println("l", dataLength, template.Bits)
	//template.Bits

	section.Data = ParseData3(f, length-5, &template)

	return section, err
}

func read(f *os.File, data ...interface{}) (err error) {
	for _, what := range data {
		err = binary.Read(f, binary.BigEndian, what)
		if err != nil {
			return err
		}
	}
	return nil
}
