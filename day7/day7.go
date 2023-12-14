package day7

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day7() {
	fmt.Println("--- Day 7: Camel Cards ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(lines []string) int {
	sort.SliceStable(lines, func(i, j int) bool {
		return compareHands(lines[i], lines[j])
	})
	result := 0
	for i, line := range lines {
		value, _ := strconv.Atoi(strings.Split(line, " ")[1])
		result += value * (i + 1)
	}
	return result
}

func compareHands(a, b string) bool {
	handsMap := map[string]int{
		"FiveOfAKind":  7,
		"FourOfAKind":  6,
		"FullHouse":    5,
		"ThreeOfAKind": 4,
		"TwoPair":      3,
		"OnePair":      2,
		"HighCard":     1,
	}
	handA := getHandType(strings.Split(a, " ")[0])
	handB := getHandType(strings.Split(b, " ")[0])
	if handsMap[handA] == handsMap[handB] {
		for i := range a {
			if a[i] != b[i] {
				return compareCards(a[i], b[i])
			}
		}
		return false
	}
	return handsMap[handA] < handsMap[handB]
}

func compareHandsPart2(a, b string) bool {
	handsMap := map[string]int{
		"FiveOfAKind":  7,
		"FourOfAKind":  6,
		"FullHouse":    5,
		"ThreeOfAKind": 4,
		"TwoPair":      3,
		"OnePair":      2,
		"HighCard":     1,
	}
	handA := getHandTypePart2(strings.Split(a, " ")[0])
	handB := getHandTypePart2(strings.Split(b, " ")[0])
	if handsMap[handA] == handsMap[handB] {
		for i := range a {
			if a[i] != b[i] {
				return compareCardsPart2(a[i], b[i])
			}
		}
		return true
	}
	return handsMap[handA] < handsMap[handB]
}

func getHandTypePart2(hand string) string {
	freq := make(map[rune]int)
	numJokers := 0
	for _, char := range hand {
		if char == 'J' {
			numJokers++
			continue
		}
		freq[char]++
	}
	switch len(freq) {
	case 0, 1:
		return "FiveOfAKind"
	case 2:
		for _, v := range freq {
			if v == 1 {
				return "FourOfAKind"
			}
		}
		return "FullHouse"
	case 3:
		if numJokers > 0 {
			return "ThreeOfAKind"
		}
		for _, v := range freq {
			if v == 3 {
				return "ThreeOfAKind"
			}
		}
		return "TwoPair"
	case 4:
		return "OnePair"
	default:
		return "HighCard"
	}
}

func getHandType(hand string) string {
	freq := make(map[rune]int)
	for _, char := range hand {
		freq[char]++
	}
	switch len(freq) {
	case 1:
		return "FiveOfAKind"
	case 2:
		for _, v := range freq {
			if v == 1 {
				return "FourOfAKind"
			}
		}
		return "FullHouse"
	case 3:
		for _, v := range freq {
			if v == 3 {
				return "ThreeOfAKind"
			}
		}
		return "TwoPair"
	case 4:
		return "OnePair"
	default:
		return "HighCard"
	}
}

func compareCardsPart2(a, b byte) bool {
	cardsMap := map[byte]int{
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'J': 0,
		'Q': 12,
		'K': 13,
		'A': 14,
	}
	return cardsMap[a] < cardsMap[b]
}

func compareCards(a, b byte) bool {
	cardsMap := map[byte]int{
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'J': 11,
		'Q': 12,
		'K': 13,
		'A': 14,
	}
	return cardsMap[a] < cardsMap[b]
}

func getPart2Answer(lines []string) int {
	sort.SliceStable(lines, func(i, j int) bool {
		return compareHandsPart2(lines[i], lines[j])
	})
	result := 0
	for i, line := range lines {
		value, _ := strconv.Atoi(strings.Split(line, " ")[1])
		result += value * (i + 1)
	}
	return result
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day7/input.txt")
	if err != nil {
		fmt.Println("error reading input")
		return lines
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}
