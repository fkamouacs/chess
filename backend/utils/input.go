package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func GetMovesFromPGN(gameFile string) ([]Move, string) {
	movesString := ""

	index := strings.Index(gameFile, "\n1.")

	if index != -1 {

		movesString = gameFile[index+1:]

	} else {
		panic("no moves found")
	}

	movesArr, matchResult := SeperateTextMoves(movesString)

	moves := ConvertToMovesArr(movesArr)

	return moves, matchResult
}

func SeperateTextMoves(text string) ([]string, string) {
	pattern := `\d+\.`
	re := regexp.MustCompile(pattern)
	// Find all matches for "number."
	matches := re.FindAllStringIndex(text, -1)

	var results []string

	for i, match := range matches {
		start := match[0]
		var end int
		if i+1 < len(matches) {
			end = matches[i+1][0] // Start of the next match
		} else {
			end = len(text) // End of the string
		}

		// Extract the substring for this section
		section := strings.TrimSpace(text[start:end])
		results = append(results, section)
	}

	lastMove := results[len(results)-1]
	matchResult := FilterResultOfMatch(lastMove)
	
	results[len(results)-1] = strings.Replace(results[len(results)-1], " " + matchResult, "", -1)
	return results, matchResult
}


func ConvertToMovesArr(movesArr []string) []Move{
	var result []Move

	for i := 0; i < len(movesArr); i++ {

		result = append(result, ConvertMove(movesArr[i]))		
	}
	return result
}


func FilterResultOfMatch(move string) string{
	pattern := `\d+-\d+`
	// Compile regex
	re := regexp.MustCompile(pattern)

	// Find all matches
	matches := re.FindAllString(move, -1)

	if len(matches) > 0 {
		// Get the last match
		lastMatch := matches[len(matches)-1]
		return lastMatch
		} else {
		return ""
	}
}

func ConvertMove(move string) Move { 
	var result Move
	
	moveNumber, whiteMove, blackMove := SplitTurn(move)

	result.MoveNumber = moveNumber
	result.WhiteMove = whiteMove
	result.BlackMove = blackMove

	return result
}

func SplitTurn(move string) (int, string, string) {

	pattern := `(\d+)\.\s+(\S+)\s+(\S+)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(move, -1)


	moveNumber, err := strconv.Atoi(matches[0][1])
	whiteMove := matches[0][2]
	blackMove := matches[0][3]
	if err != nil {
		// Handle error
		fmt.Println("Error converting string to int:", err)
	} 
	return moveNumber, whiteMove, blackMove 
}

func StringToIntArray(str string) []int {
	var result []int
	words := strings.Fields(str)
	for _, word := range words {
		num, err := strconv.Atoi(word)

		if err != nil {
			panic(err)
		}
		result = append(result, num)
	}
	
	return result
}