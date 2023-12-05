package main

import "testing"

var (
	seeds      = []int{79, 14, 55, 13}
	seedToSoil = []xToY{
		{sourceRangeStart: 50, destinationRangeStart: 98, rangeSize: 2},
		{sourceRangeStart: 52, destinationRangeStart: 50, rangeSize: 48},
	}
	soilToFertilizer = []xToY{
		{sourceRangeStart: 0, destinationRangeStart: 15, rangeSize: 37},
		{sourceRangeStart: 37, destinationRangeStart: 52, rangeSize: 2},
		{sourceRangeStart: 39, destinationRangeStart: 0, rangeSize: 15},
	}
	fertilizerToWater = []xToY{
		{sourceRangeStart: 49, destinationRangeStart: 53, rangeSize: 8},
		{sourceRangeStart: 0, destinationRangeStart: 11, rangeSize: 42},
		{sourceRangeStart: 42, destinationRangeStart: 0, rangeSize: 7},
		{sourceRangeStart: 57, destinationRangeStart: 7, rangeSize: 4},
	}
	waterToLight = []xToY{
		{sourceRangeStart: 88, destinationRangeStart: 18, rangeSize: 7},
		{sourceRangeStart: 18, destinationRangeStart: 25, rangeSize: 70},
	}
	lightToTemperature = []xToY{
		{sourceRangeStart: 45, destinationRangeStart: 77, rangeSize: 23},
		{sourceRangeStart: 81, destinationRangeStart: 45, rangeSize: 19},
		{sourceRangeStart: 68, destinationRangeStart: 64, rangeSize: 13},
	}
	temperatureToHumidity = []xToY{
		{sourceRangeStart: 0, destinationRangeStart: 69, rangeSize: 1},
		{sourceRangeStart: 1, destinationRangeStart: 0, rangeSize: 69},
	}
	humidityToLocation = []xToY{
		{sourceRangeStart: 60, destinationRangeStart: 56, rangeSize: 37},
		{sourceRangeStart: 56, destinationRangeStart: 93, rangeSize: 4},
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
