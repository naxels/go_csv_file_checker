# CSV file checker written in Golang

Works with custom delimiters

Tells you statistics about the CSV file(s) you input:
* Split count(s) in the file
* Amount of records containing that split count (by doing a simple `len()`)
* the actual records splitted within that split (valuable for using them or troubleshooting)

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

statistics, err := csvfilechecker.Read(fileLocation, delimiter)
if err != nil {
  fmt.Println(err)
}

fmt.Println(statistics.Filename)

for _, s := range statistics.Splits {
  fmt.Printf("Split count: %v, containing %v records\n", s.Count, len(s.Records))

  for i, l := range s.Records {
    fmt.Printf("%v, %v\n", i, l.Data)
    //add the index number after l.Data[] to get specific column
  }
}
```
