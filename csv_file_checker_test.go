package csvfilechecker

import (
	"os"
	"testing"
)

const (
	testDataDir    = "testdata"
	testFileSuffix = ".csv"
)

/* Old trial for all files
func TestAllCSVParse(t *testing.T) {
	fileInfos, err := ioutil.ReadDir(testDataDir)
	if err != nil {
		t.Fatalf("ioutil.ReadDir(%q) err = %v, expected nil", testDataDir, err)
	}
	for _, fileInfo := range fileInfos {
		fileName := fileInfo.Name()
		if !strings.HasSuffix(fileName, testFileSuffix) {
			continue
		}
		if _, err := Read(testDataDir+string(os.PathSeparator)+fileName, ','); err != nil {
			t.Fatalf("Read(%q) err = %v, expected nil", fileName, err)
		}
	}
}
*/

func TestCommaCSVParse(t *testing.T) {
	fileName := "csv_comma.csv"
	delimiter := ','
	expectedResult := "splits\n3 -> 5 lines"

	result, err := Read(testDataDir+string(os.PathSeparator)+fileName, delimiter)
	if err != nil {
		t.Fatalf("Read(%q) err = %v, expected nil", fileName, err)
	}
	if result != expectedResult {
		t.Fatalf("Read(%q) statistics = %v, expected %v", fileName, result, expectedResult)
	}
}

func TestCommaEnclosedCSVParse(t *testing.T) {
	fileName := "csv_comma_enclosed.csv"
	delimiter := ','
	expectedResult := "splits\n3 -> 5 lines"

	result, err := Read(testDataDir+string(os.PathSeparator)+fileName, delimiter)
	if err != nil {
		t.Fatalf("Read(%q) err = %v, expected nil", fileName, err)
	}
	if result != expectedResult {
		t.Fatalf("Read(%q) statistics = %v, expected %v", fileName, result, expectedResult)
	}
}

func TestCommaEnclosedHeaderCSVParse(t *testing.T) {
	fileName := "csv_comma_enclosed_header.csv"
	delimiter := ','
	expectedResult := "splits\n3 -> 6 lines"

	result, err := Read(testDataDir+string(os.PathSeparator)+fileName, delimiter)
	if err != nil {
		t.Fatalf("Read(%q) err = %v, expected nil", fileName, err)
	}
	if result != expectedResult {
		t.Fatalf("Read(%q) statistics = %v, expected %v", fileName, result, expectedResult)
	}
}

/*TODO
func TestCommaFewCSVParse(t *testing.T) {
	fileName := "csv_comma_few.csv"
	delimiter := ','
	expectedResult := "splits\n3 -> 4 lines\n2 -> 1 lines"

	result, err := Read(testDataDir+string(os.PathSeparator)+fileName, delimiter)
	if err != nil {
		t.Fatalf("Read(%q) err = %v, expected nil", fileName, err)
	}
	if result != expectedResult {
		t.Fatalf("Read(%q) statistics = %v, expected %v", fileName, result, expectedResult)
	}
}
*/

/*TODO
func TestCommaMoreCSVParse(t *testing.T) {
	fileName := "csv_comma_more.csv"
	delimiter := ','
	expectedResult := "splits\n3 -> 4 lines\n4 -> 1 lines"

	result, err := Read(testDataDir+string(os.PathSeparator)+fileName, delimiter)
	if err != nil {
		t.Fatalf("Read(%q) err = %v, expected nil", fileName, err)
	}
	if result != expectedResult {
		t.Fatalf("Read(%q) statistics = %v, expected %v", fileName, result, expectedResult)
	}
}
*/

func TestPipeCSVParse(t *testing.T) {
	fileName := "csv_pipe.csv"
	delimiter := '|'
	expectedResult := "splits\n3 -> 5 lines"

	result, err := Read(testDataDir+string(os.PathSeparator)+fileName, delimiter)
	if err != nil {
		t.Fatalf("Read(%q) err = %v, expected nil", fileName, err)
	}
	if result != expectedResult {
		t.Fatalf("Read(%q) statistics = %v, expected %v", fileName, result, expectedResult)
	}
}
