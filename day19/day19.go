package day19

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day19() {
	fmt.Println("--- Day 19: Linen Layout ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func isPossible(line string, options map[string]bool, index int, cache map[int]bool) bool {
	if index >= len(line) {
		return true
	}
	if v, ok := cache[index]; ok {
		return v
	}
	for i := index + 1; i <= len(line); i++ {
		if options[line[index:i]] {
			if isPossible(line, options, i, cache) {
				cache[index] = true
				return true
			}
		}
	}
	cache[index] = false
	return false
}

func getPart1Answer(lines []string) int {
	options := getOptions(lines[0])
	count := 0
	for i := 2; i < len(lines); i++ {
		cache := make(map[int]bool)
		if isPossible(lines[i], options, 0, cache) {
			count++
		}
	}
	return count
}

func getCount(line string, options map[string]bool, index int, cache map[int]int) int {
	if index >= len(line) {
		return 1
	}
	if v, ok := cache[index]; ok {
		return v
	}
	answer := 0
	for i := index + 1; i <= len(line); i++ {
		if options[line[index:i]] {
			answer += getCount(line, options, i, cache)
		}
	}
	cache[index] = answer
	return answer
}

func getPart2Answer(lines []string) int {
	options := getOptions(lines[0])
	count := 0
	for i := 2; i < len(lines); i++ {
		cache := make(map[int]int)
		count += getCount(lines[i], options, 0, cache)
	}
	return count
}

func getOptions(line string) map[string]bool {
	options := make(map[string]bool)
	for _, s := range strings.Split(line, ", ") {
		options[s] = true
	}
	return options
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day19/input.txt")
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
