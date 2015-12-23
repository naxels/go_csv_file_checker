# CSV file checker written in Golang

Written to help troubleshoot faulty CSV files quickly

Works with custom delimiters

Tells you statistics about the CSV file(s) you input:
* Split count(s) in the file
* Amount of records containing that split count
* the actual records splitted within that split (valuable for using them or troubleshooting)
  * ability to return limited actual records

You can use a go routine to input a volume of files and get statistics back

## Installation
```
go get github.com/naxels/go_csv_file_checker
```

## Basic Usage
```
import "github.com/naxels/go_csv_file_checker"
// be sure to import fmt package for this example

//single file:
fileLocation := "filelocation/filename.extension"
delimiter := ',' //any rune character
recordLimit := 0 //no recordLimit

statistics, err := csvfilechecker.Read(fileLocation, delimiter, recordLimit)
if err != nil {
  fmt.Println(err)
}

fmt.Println(statistics.Filename)

for _, s := range statistics.Splits {
  fmt.Printf("Split count: %v, containing %v records\n", s.Count, s.RecordCount)

  for i, l := range s.Records {
    fmt.Printf("%v, %v\n", i, l.Data)
    //add the index number after l.Data[] to get specific column
  }
}
```
