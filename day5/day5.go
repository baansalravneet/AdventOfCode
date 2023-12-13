package day5

import (
	"bufio"
	"fmt"
	"os"
)

func Day5() {
	fmt.Println("--- Day 5: If You Give A Seed A Fertilizer ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	// part2Answer := getPart2Answer(lines)
	// fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(lines []string) int {
	return 0
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day4/input.txt")
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
