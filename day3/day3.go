package day3

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func Day3() {
	fmt.Println("--- Day 3: Gear Ratios ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart2Answer(lines []string) int {
	answer := 0
	for i, line := range lines {
		for j, char := range line {
			if char == '*' {
				numbers := findNumbersAround(lines, i, j)
				if len(numbers) == 2 {
					answer += numbers[0] * numbers[1]
				}
			}
		}
	}
	return answer
}

func findNumbersAround(lines []string, i, j int) []int {
	numbers := []int{}
	directions := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for idx, dir := range directions {
		_i := i + dir[0]
		_j := j + dir[1]
		if _i >= 0 && _j >= 0 && _i < len(lines) && _j < len(lines[_i]) {
			if unicode.IsDigit(rune(lines[_i][_j])) {
                if idx == 1 || idx == 2 || idx == 6 || idx == 7 {
                    if unicode.IsDigit(rune(lines[i + directions[idx-1][0]][j + directions[idx-1][1]])) {
                        continue
                    }
                }
                numbers = append(numbers, collectNumber(lines, _i, _j))
			}
		}
	}
	return numbers
}

func collectNumber(lines []string, i, j int) int {
	left := j
	right := j
	for left >= 0 && unicode.IsDigit(rune(lines[i][left])) {
		left--
	}
	for right < len(lines[i]) && unicode.IsDigit(rune(lines[i][right])) {
		right++
	}
	value := 0
	for idx := left + 1; idx < right; idx++ {
		value = value*10 + int(lines[i][idx]-'0')
	}
	return value
}

func getPart1Answer(lines []string) int {
	answer := 0
	for i := 0; i < len(lines); i++ {
		j := 0
		for j < len(lines[i]) {
			if unicode.IsDigit(rune(lines[i][j])) {
				consider := false
				value := 0
				for j < len(lines[i]) && unicode.IsDigit(rune(lines[i][j])) {
					value = value*10 + int(lines[i][j]-'0')
					if nextToSymbol(lines, i, j) {
						consider = true
					}
					j++
				}
				if consider {
					answer += value
				}
			} else {
				j++
			}
		}
	}
	return answer
}

func nextToSymbol(lines []string, i, j int) bool {
	directions := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for _, dir := range directions {
		_i := i + dir[0]
		_j := j + dir[1]
		if _i >= 0 && _j >= 0 && _i < len(lines) && _j < len(lines[_i]) {
			if !unicode.IsDigit(rune(lines[_i][_j])) && lines[_i][_j] != '.' {
				return true
			}
		}
	}
	return false
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day3/input.txt")
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
