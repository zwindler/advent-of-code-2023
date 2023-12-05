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

func findLowestLocation(alm almanac) (lowest int) {
	s2SMap := buildXToYMap(alm.s2S)
	fmt.Println(s2SMap)
	return
}

func buildXToYMap(CurrentXToY []xToY) (realMap map[int]int) {
	realMap = make(map[int]int)
	for _, xtoy := range CurrentXToY {
		fmt.Println(xtoy)
		for i := 0; i < xtoy.rangeSize; i++ {
			source := xtoy.sourceRangeStart + i
			destination := xtoy.destinationRangeStart + i
			realMap[source] = destination
		}
	}

	return
}
