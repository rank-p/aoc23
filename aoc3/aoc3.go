package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
	"strconv"
)

type LocatedPart struct {
	partNumber int
	index int
}

type LocatedSymbol struct {
	symbol string
	index int
}

type ParsedLine struct {
	lineNumber int
	symbols []LocatedSymbol
	parts []LocatedPart
}

func parseLine(line string, lineNumber int) ParsedLine {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllStringIndex(line, -1)
	var parts []LocatedPart
	for _, match := range matches {
		number, _ := strconv.Atoi(line[match[0]:match[1]])
		parts = append(parts, LocatedPart{partNumber: number, index: match[0]})
	}

	re = regexp.MustCompile(`[^0-9.]`)
	matches = re.FindAllStringIndex(line, -1)
	var symbols []LocatedSymbol
	for _, match := range matches {
		symbol := line[match[0]:match[1]]
		symbols = append(symbols, LocatedSymbol{symbol: symbol, index: match[0]})
	}
	
	return ParsedLine{lineNumber: lineNumber, symbols: symbols, parts: parts}
	


}


func getSymbols(parsedLines []ParsedLine, index int) []LocatedSymbol {
	var symbols []LocatedSymbol
	symbols = append(symbols, parsedLines[index].symbols...)
	if index > 0 {
		symbols = append(symbols, parsedLines[index-1].symbols...)
	}
	if index < len(parsedLines)-1 {
		symbols = append(symbols, parsedLines[index+1].symbols...)
	}
	return symbols
}

func hasAdjacentSymbol(part LocatedPart, symbols []LocatedSymbol) bool {
	length := len(strconv.Itoa(part.partNumber))
	for _, symbol := range symbols {
		if symbol.index >= part.index-1 && symbol.index <= part.index+length {
			return true
		}
	}
	return false

}

func evaluateLine(parsedLines []ParsedLine,  index int) int {
	currentLine := parsedLines[index]
	symbols := getSymbols(parsedLines, index)
	lineValue := 0
	for _, part := range currentLine.parts {
		if hasAdjacentSymbol(part, symbols) {
			lineValue += part.partNumber
		}
	}
	return lineValue
}

func main() {

	file, _ := os.Open("aoc3.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	var parsedLines []ParsedLine

	i := 0
	for scanner.Scan() {
		line:=scanner.Text()
		parsedLines = append(parsedLines, parseLine(line, i))
		i++
	}

	answer := 0

	for index, _ := range parsedLines {
		answer += evaluateLine(parsedLines, index)
	}

	fmt.Println(answer)

	

}
