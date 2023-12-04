package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var testPartsMap = []string{
	"467..114..",
	"...*......",
	"..35..633.",
	"......#...",
	"617*......",
	".....+.58.",
	"..592.....",
	"......755.",
	"...$.*....",
	".664.598..",
}

type mapNumber struct {
	line         int
	begin        int
	end          int
	value        int
	isPartNumber bool
}

var numbersFound = []mapNumber{}

func main() {
	fmt.Println(sumPartNumbers(testPartsMap))

}

func sumPartNumbers(partsMap []string) (sum int) {
	// find all numbers
	lineNumber := 0
	for _, line := range partsMap {
		lineNumber += 1
		numbersFound = append(numbersFound, findContiguousDigitsIndex(line, lineNumber)...)
	}

	// find all symbols and update numbers that are parts
	lineNumber = 0
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
	re := regexp.MustCompile(`[\*\#\+\$\/\%\-\@\=]`)
	return re.FindAllStringIndex(currentLine, -1)
}

func checkIfSymbolConnectsWithANumber(row, column int) {
	for idx, number := range numbersFound {
		if (column <= number.end+1 && column >= number.begin-1) &&
			(row >= number.line-1 && row <= number.line+1) {
			// confirmed part number
			numbersFound[idx].isPartNumber = true
		}
	}
}
