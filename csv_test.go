package main

import (
	"testing"
	"reflect"
)

func TestCsv(t *testing.T) {
	expected := reflect.TypeOf([][]interface{}{})
	actual := reflect.TypeOf(prepare("testdata/test.csv"))
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

