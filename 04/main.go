package main

import (
	"fmt"
	"slices"
)

type card struct {
	cardID         int
	winningNumbers []int
	playingNumbers []int
}

func main() {
	filename := "puzzleData.txt"
	puzzleData, _ := readPuzzleData(filename)
	fmt.Println(sumCardValue(puzzleData))
	// fmt.Println(puzzleData)
}

func sumCardValue(puzzleData []card) (sum int) {
	for _, card := range puzzleData {
		sum += getCardValue(card)
	}
	return sum
}

func getCardValue(currentCard card) (value int) {
	for _, playing := range currentCard.playingNumbers {
		if slices.Contains(currentCard.winningNumbers, playing) {
			if value == 0 {
				value = 1
			} else {
				value *= 2
			}
		}
	}
	return value
}
