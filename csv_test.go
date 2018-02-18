package csvtogs

import (
	"reflect"
	"testing"
)

func TestCsv(t *testing.T) {
	expected := reflect.TypeOf([][]interface{}{})
	actual := reflect.TypeOf(prepare("testdata/test.csv"))
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}

	if len(prepare("testdata/test.csv")) < 1 {
		t.Errorf("The resulting interface has a length of %d", len(prepare("testdata/test.csv")))
	}
}
