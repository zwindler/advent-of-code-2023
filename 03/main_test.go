package main

import "testing"

var testPartsMap = []string{
	"467..114..",
	"...*......",
	"..35..633.",
	"......#...",
	"617*......",
	".....+.58.",
	"..592.....",
	"......755.",
	"...$.*....",
	".664.598..",
}

func TestSumPartNumbers(t *testing.T) {
	tcs := []struct {
		data          []string
		expectedValue int
	}{
		{data: testPartsMap,
			expectedValue: 4361},
	}
	for _, tc := range tcs {
		findNumbers(tc.data)
		testValue := sumPartNumbers(tc.data)
		if testValue != tc.expectedValue {
			t.Errorf("Sum value is %v, expected %d", testValue, tc.expectedValue)
		}
	}
}

func TestSumGearRatio(t *testing.T) {
	tcs := []struct {
		data          []string
		expectedValue int
	}{
		{data: testPartsMap,
			expectedValue: 467835},
	}
	for _, tc := range tcs {
		findNumbers(tc.data)
		testValue := sumGearRatio(tc.data)
		if testValue != tc.expectedValue {
			t.Errorf("Sum value is %v, expected %d", testValue, tc.expectedValue)
		}
	}
}
