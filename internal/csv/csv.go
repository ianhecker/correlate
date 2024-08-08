package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadFile(filepath string) ([][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("Could not find file: %s", filepath)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	data, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Unable to parse file: %s: %w", filepath, err)
	}

	return data, nil
}

func WriteFile(filepath string, matrix [][]string) error {
	file, err := os.Create(filepath)
	defer file.Close()

	if err != nil {
		return fmt.Errorf("Failed to open file: %w", err)
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	for _, row := range matrix {
		if err := w.Write(row); err != nil {
			return fmt.Errorf("Error writing to file: %w", err)
		}
	}
	return nil
}
