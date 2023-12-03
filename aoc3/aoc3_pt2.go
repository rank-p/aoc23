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


func getParts(parsedLines []ParsedLine, index int) []LocatedPart {
	var parts []LocatedPart
	parts = append(parts, parsedLines[index].parts...)
	if index > 0 {
		 parts = append(parts, parsedLines[index-1].parts...)
	}
	if index < len(parsedLines)-1 {
		parts = append(parts, parsedLines[index+1].parts...)
	}
	return parts 
}

func gearRatio(symbol LocatedSymbol, parts []LocatedPart) int {
	if symbol.symbol != "*" {
		return 0
	}
	gearRatio := 1
	num := 0
	for _, part := range parts {
		length := len(strconv.Itoa(part.partNumber))
		if symbol.index >= part.index-1 && symbol.index <= part.index+length {
			gearRatio *= part.partNumber
			num++
		}
	}

	if num == 2 {
		return gearRatio
	} else {
		return 0
	}
}

func evaluateLine(parsedLines []ParsedLine,  index int) int {
	currentLine := parsedLines[index]
	parts := getParts(parsedLines, index)
	lineValue := 0
	for _, symbol := range currentLine.symbols {
		lineValue += gearRatio(symbol, parts)	
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
