package testutils

import (
	"encoding/json"
	"testing"

	"github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
)

// AssertJSONEquals makes sure two json byte slices are equal
// Note: Does not currently work on arrays
func AssertJSONEquals(a, b []byte, t *testing.T) {
	d := gojsondiff.New()
	diff, err := d.Compare(a, b)
	if err != nil {
		t.Error(err)
	}
	if diff.Modified() {
		var aJSON map[string]interface{}
		err := json.Unmarshal(a, &aJSON)
		if err != nil {
			t.Error(err)
		}

		config := formatter.AsciiFormatterConfig{
			ShowArrayIndex: true,
			Coloring:       true,
		}

		format := formatter.NewAsciiFormatter(aJSON, config)
		diffString, _ := format.Format(diff)
		t.Error(diffString)
	}
}
