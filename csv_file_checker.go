package csvfilechecker

import (
	"encoding/csv"
	"io"
	"os"
)

//Statistics struct for CSV
type Statistics struct {
	Filename string
	Splits   []Split
}

//Add Split to Statistics
func (s *Statistics) Add(split Split) []Split {
	s.Splits = append(s.Splits, split)
	return s.Splits
}

//Update Split in Statistics
func (s *Statistics) Update(location int, split Split) []Split {
	s.Splits[location] = split
	return s.Splits
}

//Split struct for splitted lines
type Split struct {
	Count int
	Lines []Line
}

//Line struct for each line
type Line struct {
	Data []string
}

//Add Line to Splits
func (s *Split) Add(data Line) []Line {
	s.Lines = append(s.Lines, data)
	return s.Lines
}

//Read function takes in fileLocation and delimiter and returns Statistics struct
func Read(fileLocation string, delimiter rune) (*Statistics, error) {
	var stats = Statistics{}

	file, err := os.Open(fileLocation)
	if err != nil {
		return &stats, err
	}

	fileInfo, _ := file.Stat()
	stats.Filename = fileInfo.Name()

	r := csv.NewReader(file)
	r.Comma = delimiter

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return &stats, err
		}

		recordSplitCount := len(record)

		// if no split exists, create one
		if len(stats.Splits) == 0 {
			currentSplit := Split{Count: recordSplitCount}
			currentSplit.Add(Line{Data: record})
			stats.Add(currentSplit)
		} else {
			//test if recordSplitCount exists as Split already
			for i, s := range stats.Splits {
				var currentSplit = Split{}
				if s.Count == recordSplitCount {
					currentSplit = s
					currentSplit.Add(Line{Data: record})
					stats.Update(i, currentSplit)
				} else {
					currentSplit = Split{Count: recordSplitCount}
					currentSplit.Add(Line{Data: record})
					stats.Add(currentSplit)
				}
			}
		}
	}
	/*
		records, err := r.ReadAll()
		if err != nil {
			return &stats, err
		}
	*/
	return &stats, nil
}
