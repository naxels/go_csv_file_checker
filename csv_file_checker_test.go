package csvfilechecker

import (
	"os"
	"reflect"
	"testing"
)

const (
	testDataDir    = "testdata"
	testFileSuffix = ".csv"
)

type testpair struct {
	fileName           string
	delimiter          rune
	expectedSplitCount []int
	// expectedResult string
}

var tests = []testpair{
	{"csv_comma" + testFileSuffix, ',', []int{3}},
	{"csv_comma_enclosed" + testFileSuffix, ',', []int{3}},
	{"csv_comma_enclosed_header" + testFileSuffix, ',', []int{3}},
	{"csv_pipe" + testFileSuffix, '|', []int{3}},
	{"csv_comma_few" + testFileSuffix, ',', []int{3, 2}},
	{"csv_comma_more" + testFileSuffix, ',', []int{3, 4}},
	{"csv_comma_more_few" + testFileSuffix, ';', []int{3, 2, 4}},
	{"csv_comma_quotes" + testFileSuffix, ',', []int{3}},
}

func TestRead(t *testing.T) {
	for _, pair := range tests {
		result, err := Read(testDataDir+string(os.PathSeparator)+pair.fileName, pair.delimiter)
		if err != nil {
			t.Fatalf("Read(%q) err = %v, expected nil", pair.fileName, err)
		}

		t.Logf("result: %+v", result)

		//test fileName match
		if result.Filename != pair.fileName {
			t.Fatalf("Read(%q) fileName = %v, expected %v", pair.fileName, result.Filename, pair.fileName)
		}

		//test no splits
		if len(result.Splits) == 0 {
			t.Fatalf("This file(%q) has no splits", pair.fileName)
		}

		//test expectedSplitCount match
		//collect result.Splits Counts
		var resultSplitCount []int
		for _, s := range result.Splits {
			resultSplitCount = append(resultSplitCount, s.Count)
		}

		//check if they are equal to what is expected
		if reflect.DeepEqual(resultSplitCount, pair.expectedSplitCount) == false {
			t.Fatalf("Read(%q) SplitCount = %v, expected %v", pair.fileName, resultSplitCount, pair.expectedSplitCount)
		}
	}
}
