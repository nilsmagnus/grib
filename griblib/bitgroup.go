package griblib

import "fmt"

type bitGroupParameter struct {
	Reference uint64
	Width     uint64
	Length    uint64
}

func (bitGroup *bitGroupParameter) zeroGroup() []int64 {
	return make([]int64, bitGroup.Length)
}

func (bitGroup *bitGroupParameter) readData(bitReader *BitReader) ([]int64, error) {
	var tmp []int64
	var err error
	if bitGroup.Width != 0 {
		tmp, err = bitReader.readIntsBlock(int(bitGroup.Width), int64(bitGroup.Length), false)
		if err != nil {
			fmt.Printf("ERROR %s\n", err.Error())
		}
	} else {
		tmp = bitGroup.zeroGroup()
	}

	return tmp, err
}

//
//  Test to see if the group widths and lengths are consistent with number of
//  values, and length of section 7.
//
func checkLengths(bitGroups []bitGroupParameter, dataLength int) error {
	totBit := 0
	totLen := 0

	for _, param := range bitGroups {
		totBit += int(param.Width) * int(param.Length)
		totLen += int(param.Length)
	}

	if totBit/8 > int(dataLength) {
		return fmt.Errorf("Checksum err %d - %d", dataLength, totBit/8)
	}

	return nil
}

//
//  Extract Each Group's reference value
//
func (template *Data2) extractGroupReferences(bitReader *BitReader) ([]uint64, error) {
	numberOfGroups := int64(template.NG)
	return bitReader.readUintsBlock(int(template.Bits), numberOfGroups, true)
}

//
//  Extract Each Group's bit width
//
func (template *Data2) extractGroupBitWidths(bitReader *BitReader) ([]uint64, error) {
	numberOfGroups := int64(template.NG)
	widths, err := bitReader.readUintsBlock(int(template.GroupWidthsBits), numberOfGroups, true)
	if err != nil {
		return widths, err
	}

	for j := range widths {
		widths[j] += uint64(template.GroupWidths)
	}

	return widths, nil
}

//
//  Extract Each Group's length (number of values in each group)
//
func (template *Data2) extractGroupLengths(bitReader *BitReader) ([]uint64, error) {
	numberOfGroups := int64(template.NG)
	lengths, err := bitReader.readUintsBlock(int(template.GroupScaledLengthsBits), numberOfGroups, true)
	if err != nil {
		return lengths, err
	}

	for j := range lengths {
		lengths[j] = (lengths[j] * uint64(template.GroupLengthIncrement)) + uint64(template.GroupLengthsReference)
	}
	lengths[numberOfGroups-1] = uint64(template.GroupLastLength)
	return lengths, nil
}

func (template *Data2) extractBitGroupParameters(bitReader *BitReader) ([]bitGroupParameter, error) {
	result := []bitGroupParameter{}
	//
	//  Extract Each Group's reference value
	//
	references, err := template.extractGroupReferences(bitReader)
	if err != nil {
		return result, err
	}

	//
	//  Extract Each Group's bit width
	//
	widths, err := template.extractGroupBitWidths(bitReader)
	if err != nil {
		return result, err
	}

	//
	//  Extract Each Group's length (number of values in each group)
	//
	lengths, err := template.extractGroupLengths(bitReader)
	if err != nil {
		return result, err
	}

	for index := range references {
		result = append(result, bitGroupParameter{
			Reference: references[index],
			Width:     widths[index],
			Length:    lengths[index],
		})
	}

	bitReader.resetOffset()

	return result, nil
}
