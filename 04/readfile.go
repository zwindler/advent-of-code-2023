package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

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
