package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Read a balance file
func GetFloatFromFile(FileName string) (float64, error) {
	data, err := os.ReadFile(FileName)

	if err != nil {
		return 1000, errors.New("Failed to read file.")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored value.")
	}

	return value, nil
}

// Write a balance to file
func WriteValueToFile(filename string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueText), 0644)
}
