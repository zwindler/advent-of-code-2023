package main

import "testing"

func TestSumPartNumbers(t *testing.T) {
	tcs := []struct {
		data          []string
		expectedValue int
	}{
		{data: testPartsMap,
			expectedValue: 4361},
	}
	for _, tc := range tcs {
		testValue := sumPartNumbers(tc.data)
		if testValue != tc.expectedValue {
			t.Errorf("Sum value is %v, expected %d", testValue, tc.expectedValue)
		}
	}
}
