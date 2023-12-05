package main

import "testing"

var (
	seeds      = []int{79, 14, 55, 13}
	seedToSoil = []xToY{
		{destinationRangeStart: 50, sourceRangeStart: 98, rangeSize: 2},
		{destinationRangeStart: 52, sourceRangeStart: 50, rangeSize: 48},
	}
	soilToFertilizer = []xToY{
		{destinationRangeStart: 0, sourceRangeStart: 15, rangeSize: 37},
		{destinationRangeStart: 37, sourceRangeStart: 52, rangeSize: 2},
		{destinationRangeStart: 39, sourceRangeStart: 0, rangeSize: 15},
	}
	fertilizerToWater = []xToY{
		{destinationRangeStart: 49, sourceRangeStart: 53, rangeSize: 8},
		{destinationRangeStart: 0, sourceRangeStart: 11, rangeSize: 42},
		{destinationRangeStart: 42, sourceRangeStart: 0, rangeSize: 7},
		{destinationRangeStart: 57, sourceRangeStart: 7, rangeSize: 4},
	}
	waterToLight = []xToY{
		{destinationRangeStart: 88, sourceRangeStart: 18, rangeSize: 7},
		{destinationRangeStart: 18, sourceRangeStart: 25, rangeSize: 70},
	}
	lightToTemperature = []xToY{
		{destinationRangeStart: 45, sourceRangeStart: 77, rangeSize: 23},
		{destinationRangeStart: 81, sourceRangeStart: 45, rangeSize: 19},
		{destinationRangeStart: 68, sourceRangeStart: 64, rangeSize: 13},
	}
	temperatureToHumidity = []xToY{
		{destinationRangeStart: 0, sourceRangeStart: 69, rangeSize: 1},
		{destinationRangeStart: 1, sourceRangeStart: 0, rangeSize: 69},
	}
	humidityToLocation = []xToY{
		{destinationRangeStart: 60, sourceRangeStart: 56, rangeSize: 37},
		{destinationRangeStart: 56, sourceRangeStart: 93, rangeSize: 4},
	}
	testData = almanac{
		seedsList: seeds,
		s2S:       seedToSoil,
		S2F:       soilToFertilizer,
		f2W:       fertilizerToWater,
		w2L:       waterToLight,
		l2T:       lightToTemperature,
		t2H:       temperatureToHumidity,
		h2L:       humidityToLocation,
	}
)

func TestFindLowestLocation(t *testing.T) {
	tcs := []struct {
		data          almanac
		expectedValue int
	}{
		{data: testData,
			expectedValue: 35},
	}
	for _, tc := range tcs {
		testValue := findLowestLocation(tc.data)
		if testValue != tc.expectedValue {
			t.Errorf("Lowest value is %v, expected %d", testValue, tc.expectedValue)
		}
	}
}
