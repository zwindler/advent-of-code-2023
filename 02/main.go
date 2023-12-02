package main

type game struct {
	gameID   int
	drawList []draw
}

type draw struct {
	red   int
	blue  int
	green int
}

var (
	emptyGame = []game{{gameID: 1, drawList: []draw{{red: 0, green: 0, blue: 0}}}}
	maxRed    = 12
	maxGreen  = 13
	maxBlue   = 14
)

func main() {
	getPossibleGamesSum(emptyGame)
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
