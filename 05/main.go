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
	// S2FMap := buildXToYMap(alm.s2F)
	// f2WMap := buildXToYMap(alm.f2W)
	// w2LMap := buildXToYMap(alm.w2L)
	// l2TMap := buildXToYMap(alm.l2T)
	// t2HMap := buildXToYMap(alm.t2H)
	// h2LMap := buildXToYMap(alm.h2L)

	for _, seed := range alm.seedsList {
		fmt.Println(s2SMap[seed])
	}
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
