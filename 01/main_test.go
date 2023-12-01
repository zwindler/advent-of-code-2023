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
	testData := []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}
	expectedTestValue := 281

	testValue := getCalibrationValueV2(testData)

	if testValue != expectedTestValue {
		t.Errorf("Calibration value is %v, expected %d", testValue, expectedTestValue)
	}
}
