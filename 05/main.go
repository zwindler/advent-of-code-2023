package main

import "fmt"

type xToY struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeSize             int
}

type almanac struct {
	seedsList []int
	s2S       []xToY
	S2F       []xToY
	f2W       []xToY
	w2L       []xToY
	l2T       []xToY
	t2H       []xToY
	h2L       []xToY
}

func main() {
	puzzleData := almanac{}
	fmt.Println(findLowestLocation(puzzleData))
}

func findLowestLocation(almanac) (lowest int) {
	return
}
