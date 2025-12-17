package data_parser

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

func ReadCSV(filePath string) ([]string, error) {
	var data []string
	file, err := os.Open(filePath)
	if err != nil {
		return data, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return data, err
	}

	for _, record := range records {
		data = append(data, record...)
	}
	return data, nil
}

func ReadInt64(filePath string) ([]int64, error) {
	var data []int64
	file, err := os.Open(filePath)
	if err != nil {
		return data, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return data, err
	}

	for _, record := range records {
		for _, value := range record {
			i, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return data, err
			}
			data = append(data, i)
		}
	}
	return data, nil
}