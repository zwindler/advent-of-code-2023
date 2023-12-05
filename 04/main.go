package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
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

func readPuzzleData(filename string) ([]card, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	data := string(content)
	lines := strings.Split(data, "\n")

	var puzzleData []card
	for _, line := range lines {
		if line != "" {
			currentGame := parseLine(line)
			puzzleData = append(puzzleData, currentGame)
		}
	}

	return puzzleData, nil
}

func parseLine(line string) card {
	var currentCard card
	re := regexp.MustCompile(`Card\s+(\d+):(.+)\|(.+)`)
	matches := re.FindStringSubmatch(line)

	currentCard.cardID, _ = strconv.Atoi(matches[1])
	for _, stringNumber := range strings.Split(matches[2], " ") {
		number, _ := strconv.Atoi(stringNumber)
		if number != 0 {
			currentCard.winningNumbers = append(currentCard.winningNumbers, number)
		}
	}
	for _, stringNumber := range strings.Split(matches[3], " ") {
		number, _ := strconv.Atoi(stringNumber)
		if number != 0 {
			currentCard.playingNumbers = append(currentCard.playingNumbers, number)
		}
	}
	return currentCard
}
