package main

import "testing"

func TestSumCardValue(t *testing.T) {
	filename := "testData.txt"
	puzzleData, _ := readPuzzleData(filename)

	tcs := []struct {
		data          []card
		expectedValue int
	}{
		{data: puzzleData,
			expectedValue: 13},
	}
	for _, tc := range tcs {
		testValue := sumCardValue(tc.data)
		if testValue != tc.expectedValue {
			t.Errorf("Sum value is %v, expected %d", testValue, tc.expectedValue)
		}
	}
}


func TestTotalScratchcards(t *testing.T) {
	filename := "testData.txt"
	puzzleData, _ := readPuzzleData(filename)

	tcs := []struct {
		data          []card
		expectedValue int
	}{
		{data: puzzleData,
			expectedValue: 30},
	}
	for _, tc := range tcs {
		testValue := totalScratchcards(tc.data)
		if testValue != tc.expectedValue {
			t.Errorf("Sum value is %v, expected %d", testValue, tc.expectedValue)
		}
	}
}
