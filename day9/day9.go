package day9

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day9() {
	fmt.Println("--- Day 9: Mirage Maintenance ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(lines []string) int {
	result := 0
	for _, line := range lines {
		result += getResult(strings.Split(line, " "))
	}
	return result
}

func getResult(input []string) int {
	values := []int{}
	for _, i := range input {
		value, _ := strconv.Atoi(i)
		values = append(values, value)
	}
	return getNext(values)
}

func getResultPart2(input []string) int {
	values := []int{}
	for _, i := range input {
		value, _ := strconv.Atoi(i)
		values = append(values, value)
	}
	return getPrevious(values)
}

func getPrevious(values []int) int {
	if allZeros(values) {
		return 0
	}
	diff := []int{}
	for i := 1; i < len(values); i++ {
		diff = append(diff, values[i]-values[i-1])
	}
	previous := getPrevious(diff)
	return values[0] - previous
}

func getNext(values []int) int {
	if allZeros(values) {
		return 0
	}
	diff := []int{}
	for i := 1; i < len(values); i++ {
		diff = append(diff, values[i]-values[i-1])
	}
	next := getNext(diff)
	return next + values[len(values)-1]
}

func allZeros(values []int) bool {
	for _, i := range values {
		if i != 0 {
			return false
		}
	}
	return true
}

func getPart2Answer(lines []string) int {
	result := 0
	for _, line := range lines {
		result += getResultPart2(strings.Split(line, " "))
	}
	return result
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day9/input.txt")
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
