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

func main() {
	sumPartNumbers(testPartsMap)
}

func sumPartNumbers(partsMap []string) int {
	numbersFound := []mapNumber{}
	lineNumber := 0
	for _, line := range partsMap {
		lineNumber += 1
		numbersFound = append(numbersFound, findContiguousDigitsIndex(line, lineNumber)...)
	}

	fmt.Println(numbersFound)

	lineNumber = 0
	for _, line := range partsMap {
		lineNumber += 1
		fmt.Println(findSymbols(line))
	}
	return 0
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

func checkIfSymbolConnectsWithANumber(row, columns int){

}
