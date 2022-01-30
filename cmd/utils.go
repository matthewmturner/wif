package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func formatTime(seconds float32) string {
	return fmt.Sprintf("%f milliseconds", seconds*1_000)
}
