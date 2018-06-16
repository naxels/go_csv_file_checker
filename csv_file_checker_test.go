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
	fileName                string
	delimiter               string
	recordLimit             int
	expectedSplitCount      []int
	expectedTotalCount      int
	expectedRecordCount     []int
	expectedRecordDataCount []int
}

var tests = []testpair{
	{"csv_comma_enclosed" + testFileSuffix, ",", 1, []int{3}, 1, []int{5}, []int{1}}, // recordlimit per Split test
	{"csv_comma" + testFileSuffix, ",", 0, []int{3}, 1, []int{5}, []int{5}},
	{"csv_comma_enclosed" + testFileSuffix, ",", 0, []int{3}, 1, []int{5}, []int{5}},
	{"csv_comma_enclosed_header" + testFileSuffix, ",", 0, []int{3}, 1, []int{6}, []int{6}},
	{"csv_pipe" + testFileSuffix, "|", 0, []int{3}, 1, []int{5}, []int{5}},
	{"csv_comma_few" + testFileSuffix, ",", 0, []int{3, 2}, 2, []int{4, 1}, []int{4, 1}},
	{"csv_comma_more" + testFileSuffix, ",", 0, []int{3, 4}, 2, []int{4, 1}, []int{4, 1}},
	{"csv_comma_more_few" + testFileSuffix, ";", 0, []int{3, 2, 4}, 3, []int{3, 1, 1}, []int{3, 1, 1}},
	{"csv_comma_quotes" + testFileSuffix, ",", 0, []int{3}, 1, []int{5}, []int{5}},
}

func TestRead(t *testing.T) {
	for _, pair := range tests {
		result, err := Read(testDataDir+string(os.PathSeparator)+pair.fileName, pair.delimiter, pair.recordLimit)
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

		//test expectedSplitCount, expectedRecordCount match
		//collect result.Splits Counts and result.Splits.Records Counts
		var resultSplitCount []int
		var recordSplitCount []int
		var recordSplitRecordDataCount []int
		for _, s := range result.Splits {
			t.Logf("stats: SplitCount: %v, RecordCount: %v, RecordDataCount: %v", s.Count, s.RecordCount, len(s.Records))
			resultSplitCount = append(resultSplitCount, s.Count)
			recordSplitCount = append(recordSplitCount, s.RecordCount)
			recordSplitRecordDataCount = append(recordSplitRecordDataCount, len(s.Records))
		}

		//check if resultSplitCount equal to what is expected
		if reflect.DeepEqual(resultSplitCount, pair.expectedSplitCount) == false {
			t.Fatalf("Read(%q) SplitCount = %v, expected %v", pair.fileName, resultSplitCount, pair.expectedSplitCount)
		}

		//check if recordSplitCount equal to what is expected
		if reflect.DeepEqual(recordSplitCount, pair.expectedRecordCount) == false {
			t.Fatalf("Read(%q) RecordCount = %v, expected %v", pair.fileName, recordSplitCount, pair.expectedRecordCount)
		}

		//test expectedTotalCount match
		if result.Count != pair.expectedTotalCount {
			t.Fatalf("Read(%q) Count = %v, expected %v", pair.fileName, result.Count, pair.expectedTotalCount)
		}

		//test expectedRecordDataCount match
		if reflect.DeepEqual(recordSplitRecordDataCount, pair.expectedRecordDataCount) == false {
			t.Fatalf("Read(%q) RecordDataCount = %v, expected %v", pair.fileName, recordSplitRecordDataCount, pair.expectedRecordDataCount)
		}
	}
}
