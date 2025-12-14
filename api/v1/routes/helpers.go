package dataparser

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

// Helpers provides utility functions for data parsing
type Helpers struct{}

func NewHelpers() *Helpers {
	return &Helpers{}
}

// ReadCSVFile reads a CSV file and returns a map of rows to their corresponding values
func (h *Helpers) ReadCSVFile(filePath string) (map[string][]map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	data := map[string][]map[string]string{}
	for _, record := range records {
		dataKey := strings.Join(record[:len(record)-1], ",")
		dataValue := strings.Join(record[len(record)-1:], ",")
		if _, ok := data[dataKey]; !ok {
			data[dataKey] = []map[string]string{}
		}
		data[dataKey] = append(data[dataKey], map[string]string{"value": dataValue})
	}

	return data, nil
}

// ValidateCSV reads a CSV file and validates its structure
func (h *Helpers) ValidateCSV(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	columns := make(map[int][]string)
	for _, record := range records {
		for i, columnValue := range record {
			if _, ok := columns[i]; !ok {
				columns[i] = []string{}
			}
			columns[i] = append(columns[i], columnValue)
		}
	}

	for columnIndex := range columns {
		if len(columns[columnIndex]) != len(records) {
			return fmt.Errorf("column %d has different length", columnIndex)
		}
	}

	return nil
}

// ToStringArray converts a string slice to a string array
func ToStringArray(s []string) []string {
	result := make([]string, len(s))
	copy(result, s)
	return result
}

// IsEmptyString checks if a string is empty
func IsEmptyString(s string) bool {
	return strings.TrimSpace(s) == ""
}

// IsEmptyStringArray checks if a string array is empty
func IsEmptyStringArray(s []string) bool {
	for _, value := range s {
		if !IsEmptyString(value) {
			return false
		}
	}
	return true
}

// GetFileExtension returns the file extension from a file path
func GetFileExtension(filePath string) string {
	return filePath[strings.LastIndexByte(filePath, '.')+1:]
}

// IsCSVFile checks if a file is a CSV file
func IsCSVFile(filePath string) bool {
	return GetFileExtension(filePath) == "csv"
}

// GetFileContent returns the content of a file
func GetFileContent(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var buffer bytes.Buffer
	reader := bufio.NewReader(file)
	for {
		line, isPrefix, err := reader.ReadLine()
		if err != nil || !isPrefix && err == io.EOF {
			break
		}
		buffer.WriteString(string(line) + "\n")
	}
	return buffer.String(), nil
}

// GetFileChecksum returns the checksum of a file
func GetFileChecksum(filePath string) (string, error) {
	hasher := sha256.New()
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(hasher, file)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}