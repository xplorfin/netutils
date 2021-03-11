package testutils

import (
	"encoding/json"
	"testing"

	"github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
)

// NOTE: Does not currently work on arrays
func AssertJsonEquals(a, b []byte, t *testing.T) {
	d := gojsondiff.New()
	diff, err := d.Compare(a, b)
	if err != nil {
		t.Error(err)
	}
	if diff.Modified() {
		var aJson map[string]interface{}
		err := json.Unmarshal(a, &aJson)
		if err != nil {
			t.Error(err)
		}

		config := formatter.AsciiFormatterConfig{
			ShowArrayIndex: true,
			Coloring:       true,
		}

		format := formatter.NewAsciiFormatter(aJson, config)
		diffString, _ := format.Format(diff)
		t.Error(diffString)
	}
}
