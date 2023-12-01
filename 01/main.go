package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	filename := "puzzleData.txt"
	puzzleData, _ := readPuzzleData(filename)

	fmt.Println(getCalibrationValue(puzzleData))
	fmt.Println(getCalibrationValueV2(puzzleData))
}

func getCalibrationValue(puzzleInput []string) int {
	var sum = 0
	for _, line := range puzzleInput {
		re := regexp.MustCompile("[0-9]")
		allDigits := re.FindAllString(line, -1)
		result, _ := strconv.Atoi(allDigits[0] + allDigits[len(allDigits)-1])
		sum += result
	}
	return sum
}

func getCalibrationValueV2(puzzleInput []string) int {
	var fixedPuzzleInput []string
	// this is a bit tricky because there might be some overlapping words like oneight
	// for the first digit you need to take 1ight, for the last on8
	r := strings.NewReplacer("one", "o1e",
		"two", "t2o",
		"three", "t3e",
		"four", "f4r",
		"five", "f5e",
		"six", "s6x",
		"seven", "s7n",
		"eight", "e8t",
		"nine", "n9e")

	for _, line := range puzzleInput {
		// run replacer twice to deal with overlapping words
		fixedPuzzleInput = append(fixedPuzzleInput, r.Replace(r.Replace(line)))
	}
	return getCalibrationValue(fixedPuzzleInput)
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
