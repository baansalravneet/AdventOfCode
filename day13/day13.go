package day13

import (
	"bufio"
	"fmt"
	"os"
)

func Day13() {
	fmt.Println("--- Day 13: Point of Incidence ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(patterns [][]string) int {
	result := 0
	for _, pattern := range patterns {
		horizontal := getHorizontalMirror(pattern, 0)
		if horizontal != 0 {
			result += 100 * horizontal
			continue
		}
		vertical := getVerticalMirror(pattern, 0)
		result += vertical
	}
	return result
}

func getTranspose(mat []string) []string {
	transpose := []string{}
	for i := 0; i < len(mat[0]); i++ {
		current := []byte{}
		for j := 0; j < len(mat); j++ {
			current = append(current, mat[j][i])
		}
		transpose = append(transpose, string(current))
	}
	return transpose
}

func getVerticalMirror(pattern []string, allowance int) int {
	transpose := getTranspose(pattern)
	return getHorizontalMirror(transpose, allowance)
}

func getPatterns(lines []string) [][]string {
	result := [][]string{}
	current := []string{}
	for _, l := range lines {
		if len(l) == 0 {
			result = append(result, current)
			current = []string{}
		} else {
			current = append(current, l)
		}
	}
	return append(result, current)
}

func checkMirror(pattern []string, position int, allowance int) bool {
	for i, j := position, position-1; i < len(pattern) && j >= 0; i, j = i+1, j-1 {
		for x := range pattern[i] {
			if pattern[i][x] != pattern[j][x] {
				if allowance == 0 {
					return false
				} else {
					allowance--
				}
			}
		}
	}
	return allowance == 0

}

func getHorizontalMirror(pattern []string, allowance int) int {
	for i := 1; i < len(pattern); i++ {
		if checkMirror(pattern, i, allowance) {
			return i
		}
	}
	return 0
}

func getPart2Answer(patterns [][]string) int {
	result := 0
	for _, pattern := range patterns {
		horizontal := getHorizontalMirror(pattern, 1)
		if horizontal != 0 {
			result += 100 * horizontal
			continue
		}
		vertical := getVerticalMirror(pattern, 1)
		result += vertical
	}
	return result
}

func getInput() [][]string {
	lines := []string{}
	file, err := os.Open("day13/input.txt")
	if err != nil {
		fmt.Println("error reading input")
		return [][]string{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return getPatterns(lines)
}
