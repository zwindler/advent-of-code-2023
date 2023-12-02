package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type game struct {
	gameID   int
	drawList []draw
}

type draw struct {
	blue  int
	green int
	red   int
}

var (
	emptyGame = []game{{gameID: 1, drawList: []draw{{red: 0, green: 0, blue: 0}}}}
	maxRed    = 12
	maxGreen  = 13
	maxBlue   = 14
)

func main() {
	filename := "puzzleData.txt"
	puzzleData, _ := readPuzzleData(filename)

	fmt.Println(getPossibleGamesSum(puzzleData))
}

func getPossibleGamesSum(puzzleInput []game) int {
	var sum int
	for _, game := range puzzleInput {
		if isGamePossible(game) {
			sum += game.gameID
		}
	}
	return sum
}

func isGamePossible(currentGame game) bool {
	for _, draw := range currentGame.drawList {
		if draw.blue > maxBlue || draw.green > maxGreen || draw.red > maxRed {
			return false
		}
	}
	return true
}

func readPuzzleData(filename string) ([]game, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	data := string(content)
	lines := strings.Split(data, "\n")

	var puzzleData []game
	for _, line := range lines {
		if line != "" {
			currentGame := parseLine(line)
			puzzleData = append(puzzleData, currentGame)
		}
	}

	return puzzleData, nil
}

func parseLine(line string) game {
	var currentGame game
	re := regexp.MustCompile(`Game (\d+): (.+)`)
	matches := re.FindStringSubmatch(line)

	currentGame.gameID, _ = strconv.Atoi(matches[1])

	draws := strings.Split(matches[2], ";")
	for _, drawString := range draws {
		blue, green, red := parseDraw(drawString)
		currentDraw := draw{blue: blue, green: green, red: red}
		currentGame.drawList = append(currentGame.drawList, currentDraw)
	}

	fmt.Println(currentGame)

	return currentGame
}

func parseDraw(currentDraw string) (blue, green, red int) {
	reBlue := regexp.MustCompile(`(\d+) blue`)
	reGreen := regexp.MustCompile(`(\d+) green`)
	reRed := regexp.MustCompile(`(\d+) red`)

	matchesBlue := reBlue.FindStringSubmatch(currentDraw)
	if matchesBlue != nil {
		blue, _ = strconv.Atoi(matchesBlue[1])
	}
	matchesGreen := reGreen.FindStringSubmatch(currentDraw)
	if matchesGreen != nil {
		green, _ = strconv.Atoi(matchesGreen[1])
	}
	matchesRed := reRed.FindStringSubmatch(currentDraw)
	if matchesRed != nil {
		red, _ = strconv.Atoi(matchesRed[1])
	}
	return blue, green, red
}
