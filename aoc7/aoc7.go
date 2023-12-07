package main

import (
	"fmt"
	 "bufio"
	 "os"
	"strings"
	"strconv"
	"sort"
)

type HandAndBid struct {
	hand string
	bid int
}

func compareHandsByRune(h1 string, h2 string) bool {
	var strengths = map[byte]int{
    		'A': 14,
    		'K': 13,
    		'Q': 12,
    		'J': 11,
    		'T': 10,
    		'9': 9,
    		'8': 8,
    		'7': 7,
    		'6': 6,
    		'5': 5,
    		'4': 4,
    		'3': 3,
   		'2': 2,
	}

	if strengths[h1[0]] > strengths[h2[0]] {
		return true
	} else if strengths[h2[0]] > strengths[h1[0]] {
		return false
	} else {
		return compareHandsByRune(h1[1:], h2[1:])
	} 
		
}

func compareHands(h1 string, h2 string) bool {
	if rankHand(h1) > rankHand(h2) {
		return true
	} else if rankHand(h2) > rankHand(h1) {
		return false
	} else {
		return compareHandsByRune(h1, h2)
	}

}

func hasCharWithCount(s string, count int) bool {
	charCount := make(map[rune]int)
	for _, char := range s {
		charCount[char]++
	}
	
	for _, c := range charCount {
		if c == count {
			return true
		}
	}
	return false
}

func hasDoublePair(s string) bool {
	charCount := make(map[rune]int)
	for _, char := range s {
		charCount[char]++
	}
	countPair := 0
	for _, c := range charCount {
		if c == 2 {
			countPair++
		}
	}
	return countPair == 2
}

func rankHand(hand string) int {
	if hasCharWithCount(hand, 5) {
		return 7
	} else if hasCharWithCount(hand, 4) {
		return 6
	} else if hasCharWithCount(hand, 3) && hasCharWithCount(hand, 2) {
		return 5
	} else if hasCharWithCount(hand, 3) {
		return 4
	} else if hasDoublePair(hand) {
		return 3
	} else if hasCharWithCount(hand, 2) {
		return 2
	} else {
		return 1
	}
}

func main() {
	file, _ := os.Open("aoc7.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	var arr []HandAndBid
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), " ")
		bid, _ := strconv.Atoi(tmp[1])
		arr = append(arr, HandAndBid{hand:tmp[0], bid:bid})
	}
	
	sort.Slice(arr, func(i, j int) bool {
		return compareHands(arr[j].hand, arr[i].hand)
	})	
	score := 0
	for i := range arr {
		score += (i+1) * arr[i].bid
	}
	fmt.Println(arr)
	fmt.Println(score)
}
