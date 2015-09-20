CSV file checker written in Golang

can work with custom delimiters and enclosures

tells you statistics about the file(s) you input:
* amount of splits
* lines containing that split
* TODO: the lines within that split

you can use a go routine to input a volume of files

## Installation
```
go get github.com/naxels/go_csv_file_checker
```

## Basic Usage
```
import "github.com/naxels/go_csv_file_checker"

//single file:
file := "filelocation/filename.extension"
delimiter := ',' //any rune character

statistics, err := csvfilechecker.Read(fileLocation, delimiter)
if err != nil {
  fmt.Println(err)
}

fmt.Println(statistics.Filename)

for _, s := range statistics.Splits {
  fmt.Printf("Split count: %v\n", s.Count)

  for i, l := range s.Lines {
    fmt.Printf("%v, %v\n", i, l.Data)
  }
}
```

#### Roadmap
* make the few and more CSV files work
