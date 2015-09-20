package csvfilechecker

import (
	"encoding/csv"
	"fmt"
	"os"
)

//Read function that takes in file location and returns statistics
func Read(filename string, delimiter rune) (statistics string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	r := csv.NewReader(file)
	r.Comma = delimiter

	records, err := r.ReadAll()
	if err != nil {
		return "", err
	}

	result := "splits\n"

	resultMap := make(map[int]int)

	//get the length of each record and record it in Map
	for record := range records {
		resultMap[len(records[record])]++
	}

	//convert resultMap to statistics
	for r, v := range resultMap {
		result += fmt.Sprintf("%v -> %v lines", r, v)
	}

	return result, nil
}
