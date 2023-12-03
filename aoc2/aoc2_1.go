package main

import ( 
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func parseGame(game string) map[string]int {
	cubeSets := strings.Split(game, ", ")
	parsedGame := make(map[string]int)
	for i := 0; i < len(cubeSets); i++ {
		set := strings.Split(cubeSets[i], " ")
		parsedGame[set[1]], _ = strconv.Atoi(set[0])
	}
	return parsedGame

}

func isGameValid(game map[string]int, gameConfig map[string]int) bool {
	for k, v:= range game {
		if v > gameConfig[k] {
			return false
		}
	}
	return true
}

func parseLine(line string, gameConfig map[string]int) int {
	splitLine := strings.Split(line,":")		
	gameID, _ := strconv.Atoi(strings.Split(splitLine[0], " ")[1])
	games := strings.Split(splitLine[1], ";")
	for i:=0; i < len(games); i++ {
		if !isGameValid(parseGame(strings.TrimSpace(games[i])), gameConfig) {
			return 0
		}
	}	
	return gameID
}

func main() {
	gameConfig := make(map[string]int)
	gameConfig["green"] = 13
	gameConfig["red"] = 12
	gameConfig["blue"] = 14

	answer := 0

	file, err := os.Open("aoc2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line:= scanner.Text()
		fmt.Println(parseLine(line, gameConfig))
		answer += parseLine(line, gameConfig)
	}

	if err:= scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

	fmt.Println(answer)

}
