package tools

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	testCases := []struct {
		given string
		want  map[string]int
	}{
		{"aa", map[string]int{"a": 2}},
		{"abc bca", map[string]int{" ": 1, "a": 2, "b": 2, "c": 2}},
	}

	for _, testCase := range testCases {
		got := Count(testCase.given)
		isEqual := reflect.DeepEqual(testCase.want, got)
		if !isEqual {
			t.Errorf("not equel %v - %v", got, testCase.want)
		}
		//assert.Equalf(t, testCase.want, got, "not equel %v - %v", got, testCase.want)
	}

}
