package csvfilechecker

import (
	"encoding/csv"
	"io"
	"os"
)

//Statistics struct for CSV
type Statistics struct {
	Filename string
	Count    int
	Splits   []Split
}

//Add Split to Statistics
func (s *Statistics) Add(split Split) {
	s.Splits = append(s.Splits, split)
	s.Count++
}

//Update Split in Statistics
func (s *Statistics) Update(location int, split Split) {
	s.Splits[location] = split
}

//ProcessRecord into Split
func (s *Statistics) ProcessRecord(recordRaw []string) {
	recordSplitCount := len(recordRaw)

	//initialize found and set to -1 to avoid first index (0) problem
	found := -1

	for i, s := range s.Splits {
		if s.Count == recordSplitCount {
			found = i
			break
		}
	}

	var currentSplit = Split{Count: recordSplitCount}

	// if no Splits or not found Add, else Update
	if len(s.Splits) == 0 || found == -1 {
		currentSplit.Add(Record{Data: recordRaw})
		s.Add(currentSplit)
	} else {
		currentSplit = s.Splits[found]
		currentSplit.Add(Record{Data: recordRaw})
		s.Update(found, currentSplit)
	}
}

//Split struct for splitted records
type Split struct {
	Count       int
	RecordCount int
	Records     []Record
}

//Add Record to Splits
func (s *Split) Add(data Record) {
	s.Records = append(s.Records, data)
	s.RecordCount++
}

//Record struct for each record
type Record struct {
	Data []string
}

//Read returns a new Statistics after opening file and processing it
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
	r.FieldsPerRecord = -1 // one of the key statistics is discovering broken CSV files
	r.LazyQuotes = true    // the unpredicability of incoming data should allow "'s to appear

	for {
		recordRaw, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return &stats, err
		}

		stats.ProcessRecord(recordRaw)
	}

	file.Close()

	return &stats, nil
}
