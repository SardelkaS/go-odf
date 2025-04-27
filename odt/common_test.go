package odt

import (
	"reflect"
	"strings"
	"testing"
)

func expectEqual(t *testing.T, expected, actual any) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nExpected: %v (%T)\nActual:   %v (%T)", expected, expected, actual, actual)
	}
}

func expectNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("\nError: %s", err.Error())
	}
}

func cleanString(s string) string {
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	s = strings.Replace(s, " ", "", -1)
	return s
}
