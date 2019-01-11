package gribtest

import (
	"bufio"
	"io"
	"strconv"
)

func readCsvAsSlice(file io.Reader) ([]float64, error) {
	result := []float64{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return result, err
		}
		result = append(result, f)
	}
	return result, nil
}
