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
	h1_1 := strings.Replace(h1, "J", mostCommonChar(h1), -1)
	h2_1 := strings.Replace(h2, "J", mostCommonChar(h2), -1)
	h1_1 = strings.Replace(h1_1, "'", "", -1)
	h2_1 = strings.Replace(h2_1, "'", "", -1)
	if rankHand(h1_1) > rankHand(h2_1) {
		return true
	} else if rankHand(h2_1) > rankHand(h1_1) {
		return false
	} else {
		return compareHandsByRune(h1, h2)
	}

}

func hasCharWithCount(s string, count int) bool {
	charCount := make(map[rune]int)
	for _, char := range s {
		if char == 39 {continue}
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
		if char == 39 {continue}
		charCount[char]++
	}
	countPair := 0
	fmt.Println(s, charCount)
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

func mostCommonChar(s string) string {
	charCount := make(map[rune]int)
	for _, char := range s {
		charCount[char]++
	}
	var mostCommonChar rune
	maxCount := 0
	for char, count := range charCount {
		if char == 'J' {
			continue
		}
		if count > maxCount {
			mostCommonChar = char
			maxCount = count
		}
	}

	return strconv.QuoteRune(mostCommonChar)

}

func main() {
	file, _ := os.Open("aoc7test.txt")
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
		fmt.Println(arr[i].hand, rankHand(strings.Replace(arr[i].hand, "J", mostCommonChar(arr[i].hand), -1)))
		score += (i+1) * arr[i].bid
	}
	fmt.Println(arr)
	fmt.Println(score)
}
