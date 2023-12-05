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
	fmt.Println(totalScratchcards(puzzleData))
}

func sumCardValue(puzzleData []card) (sum int) {
	for _, card := range puzzleData {
		sum += getCardValue(card)
	}
	return
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
	return
}

func totalScratchcards(puzzleData []card) (sum int) {
	cardCopiesMap := make(map[int]int)

	// initialize map
	for _, card := range puzzleData {
		cardCopiesMap[card.cardID] = 1
	}

	// find copies
	for _, card := range puzzleData {
		result := getCardCopies(card)
		if result > 0 {
			for i := 1; i <= getCardCopies(card); i++ {
				currentIndex := card.cardID + i
				// we add as much copies as there are cards of the currentCard number
				cardCopiesMap[currentIndex] += cardCopiesMap[card.cardID]
			}
		}
	}

	// sum all cards
	for _, value := range cardCopiesMap {
		sum += value
	}

	return
}

func getCardCopies(currentCard card) (result int) {
	for _, playing := range currentCard.playingNumbers {
		if slices.Contains(currentCard.winningNumbers, playing) {
			result += 1
		}
	}
	return
}
