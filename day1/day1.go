package day1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

func Day1() {
	fmt.Println("--- Day 1: Trebuchet?! ---")
	part1Answer := getPart1Answer()
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer()
	fmt.Println("Part2:", part2Answer)
}

func getPart2Answer() int {
	lines := getInput()
	result := 0
	numbersMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"zero":  0,
	}
	reverseMap := map[string]int{
		"eno":   1,
		"owt":   2,
		"eerht": 3,
		"ruof":  4,
		"evif":  5,
		"xis":   6,
		"neves": 7,
		"thgie": 8,
		"enin":  9,
		"orez":  0,
	}
	for _, line := range lines {
		left := -1
		right := -1
		indexLeft := len(line)
		indexRight := -1
		for i, char := range line {
			if unicode.IsDigit(char) {
				right = int(char - '0')
				indexRight = i
				if left == -1 {
					left = int(char - '0')
					indexLeft = i
				}
			}
		}
		indexWordLeft, wordLeft := findFirstSpelledNumber(line, numbersMap)
		indexWordRight, wordRight := findLastSpelledNumber(line, reverseMap)
		if indexWordLeft < indexLeft && wordLeft != -1 {
			left = wordLeft
		}
		if indexWordRight > indexRight && wordRight != -1 {
			right = wordRight
		}
		result += 10*left + right
	}
	return result
}

func findFirstSpelledNumber(line string, numbersMap map[string]int) (int, int) {
	leftMost := math.MaxInt32
	number := -1
	for k, v := range numbersMap {
		index := strings.Index(line, k)
		if index == -1 {
			continue
		}
		if index < leftMost {
			leftMost = index
			number = v
		}
	}
	return leftMost, number
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func findLastSpelledNumber(line string, numbersMap map[string]int) (int, int) {
	line = reverse(line)
	leftMost := math.MaxInt32
	number := -1
	for k, v := range numbersMap {
		index := strings.Index(line, k)
		if index == -1 {
			continue
		}
		if index < leftMost {
			leftMost = index
			number = v
		}
	}
	return len(line) - 1 - leftMost, number
}

func getPart1Answer() int {
	lines := getInput()
	result := 0
	for _, line := range lines {
		left := -1
		right := -1
		for _, char := range line {
			if unicode.IsDigit(char) {
				right = int(char - '0')
				if left == -1 {
					left = int(char - '0')
				}
			}
		}
		result += 10*left + right
	}
	return result
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day1/input.txt")
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
