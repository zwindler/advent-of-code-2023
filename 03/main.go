package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type mapNumber struct {
	line         int
	begin        int
	end          int
	value        int
	isPartNumber bool
}

var numbersFound = []mapNumber{}

func main() {
	filename := "puzzleData.txt"
	puzzleData, _ := readPuzzleData(filename)

	fmt.Println(sumPartNumbers(puzzleData))
}

func sumPartNumbers(partsMap []string) (sum int) {
	findNumbers(partsMap)

	// find all symbols and update numbers that are parts
	lineNumber := 0
	for _, line := range partsMap {
		lineNumber += 1
		for _, symbol := range findSymbols(line) {
			checkIfSymbolConnectsWithANumber(lineNumber, symbol[0])
		}
	}

	for _, number := range numbersFound {
		if number.isPartNumber {
			sum += number.value
		}
	}
	return sum
}

func findNumbers(partsMap []string) {
	lineNumber := 0
	for _, line := range partsMap {
		lineNumber += 1
		numbersFound = append(numbersFound, findContiguousDigitsIndex(line, lineNumber)...)
	}
}

func findContiguousDigitsIndex(line string, lineNumber int) (numbersInLine []mapNumber) {
	re := regexp.MustCompile(`[0-9]+`)
	for _, numberMatch := range re.FindAllStringIndex(line, -1) {
		value, _ := strconv.Atoi(line[numberMatch[0]:numberMatch[1]])
		currentNumber := mapNumber{
			line:         lineNumber,
			begin:        numberMatch[0],
			end:          numberMatch[1],
			value:        value,
			isPartNumber: false,
		}
		numbersInLine = append(numbersInLine, currentNumber)
	}
	return numbersInLine
}

func findSymbols(currentLine string) [][]int {
	re := regexp.MustCompile(`[\/\*\&\+\-\#\@\$\=\%]`)

	return re.FindAllStringIndex(currentLine, -1)
}

func findStars(currentLine string) [][]int {
	re := regexp.MustCompile(`[\*]`)

	return re.FindAllStringIndex(currentLine, -1)
}

func checkIfSymbolConnectsWithANumber(row, column int) {
	for idx, number := range numbersFound {
		if (column <= number.end && column >= number.begin-1) &&
			(row >= number.line-1 && row <= number.line+1) {
			// confirmed part number
			numbersFound[idx].isPartNumber = true
		}
	}
}

func sumGearRatio(partsMap []string) (sum int) {
	findNumbers(partsMap)

	// find all "*" and check if they have 2 numbers attached
	lineNumber := 0
	for _, line := range partsMap {
		lineNumber += 1
		for _, star := range findStars(line) {
			sum += checkIfStarConnectsWith2Numbers(lineNumber, star[0])
		}
	}
	return sum
}

func checkIfStarConnectsWith2Numbers(row, column int) int {
	connected := []int{}
	for _, number := range numbersFound {
		if (column <= number.end && column >= number.begin-1) &&
			(row >= number.line-1 && row <= number.line+1) {
			// confirmed part connected to star
			connected = append(connected, number.value)
		}
	}
	if len(connected) == 2 {
		return connected[0] * connected[1]
	}
	return 0
}

func readPuzzleData(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	data := string(content)
	lines := strings.Split(data, "\n")

	var puzzleData []string
	for _, line := range lines {
		if line != "" {
			puzzleData = append(puzzleData, line)
		}
	}

	return puzzleData, nil
}
