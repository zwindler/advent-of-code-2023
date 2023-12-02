package main

import (
	"testing"
)

func TestGetCalibrationValue(t *testing.T) {
	testData := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	expectedTestValue := 142

	testValue := getCalibrationValue(testData)

	if testValue != expectedTestValue {
		t.Errorf("Calibration value is %v, expected %d", testValue, expectedTestValue)
	}
}

func TestGetCalibrationValueV2(t *testing.T) {
	tcs := []struct {
		data          []string
		expectedValue int
	}{
		{
			data:          []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"},
			expectedValue: 281},
		{
			data:          []string{"twoneoneight"},
			expectedValue: 28},
	}

	for _, tc := range tcs {
		testValue := getCalibrationValueV2(tc.data)
		if testValue != tc.expectedValue {
			t.Errorf("Calibration value is %v, expected %d", testValue, tc.expectedValue)
		}
	}
}
