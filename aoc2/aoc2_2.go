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

func max(a int, b int) int {
	if a>=b {
		return a
	} else {
		return b
	}


}

func parseLine(line string, gameConfig map[string]int) int {
	splitLine := strings.Split(line,":")		
	games := strings.Split(splitLine[1], ";")
	maximalRed := 0
	maximalGreen := 0
	maximalBlue := 0
	for i:=0; i < len(games); i++ {
		parsedGame := parseGame(strings.TrimSpace(games[i])) 
		maximalRed = max(maximalRed, parsedGame["red"])
		maximalGreen = max(maximalGreen, parsedGame["green"])
		maximalBlue = max(maximalBlue, parsedGame["blue"])
		
	}	
	return maximalRed * maximalBlue * maximalGreen 
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
		answer += parseLine(line, gameConfig)
	}

	if err:= scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

	fmt.Println(answer)

}
