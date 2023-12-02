package main

import (
	"testing"
)

var exampleGame = []game{{gameID: 1,
	drawList: []draw{{blue: 3, red: 4},
		{red: 1, green: 2, blue: 6},
		{green: 2}}},
	{gameID: 2,
		drawList: []draw{{blue: 1, green: 2},
			{green: 3, blue: 4, red: 1},
			{green: 1, blue: 1}}},
	{gameID: 3,
		drawList: []draw{{green: 8, blue: 6, red: 20},
			{blue: 5, red: 4, green: 13},
			{green: 5, red: 1}}},
	{gameID: 4,
		drawList: []draw{{green: 1, red: 3, blue: 6},
			{green: 3, red: 6},
			{green: 3, blue: 15, red: 14}}},
	{gameID: 5,
		drawList: []draw{{red: 6, blue: 1, green: 3},
			{blue: 2, red: 1, green: 2}}}}

func TestGetPossibleGamesSum(t *testing.T) {
	tcs := []struct {
		data          []game
		expectedValue int
	}{
		{data: emptyGame,
			expectedValue: 1},
		{data: exampleGame,
			expectedValue: 8},
	}
	for _, tc := range tcs {
		testValue := getPossibleGamesSum(tc.data)
		if testValue != tc.expectedValue {
			t.Errorf("Games Sum value is %v, expected %d", testValue, tc.expectedValue)
		}
	}
}

func TestTetMinimumGamePowerSum(t *testing.T) {
	tcs := []struct {
		data          []game
		expectedValue int
	}{
		{data: exampleGame,
			expectedValue: 2286},
	}

	for _, tc := range tcs {
		testValue := getMinimumGamePowerSum(tc.data)
		if testValue != tc.expectedValue {
			t.Errorf("Games Power Sum value is %v, expected %d", testValue, tc.expectedValue)
		}
	}
}
