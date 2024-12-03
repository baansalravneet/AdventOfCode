package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day3() {
	fmt.Println("--- Day 3: Mull It Over ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(lines []string) int {
	result := 0
	r := regexp.MustCompile("mul\\([0-9]{1,3},[0-9]{1,3}\\)")
	for _, line := range lines {
		matches := r.FindAllString(line, -1)
		for _, match := range matches {
			index := strings.Index(match, ",")
			firstNumber, _ := strconv.Atoi(match[4:index])
			secondNumber, _ := strconv.Atoi(match[index+1 : len(match)-1])
			result += firstNumber * secondNumber
		}
	}
	return result
}

func getPart2Answer(lines []string) int {
	result := 0
	rMul := regexp.MustCompile("mul\\([0-9]{1,3},[0-9]{1,3}\\)")
	rDont := regexp.MustCompile("don't\\(\\)")
	rDo := regexp.MustCompile("do\\(\\)")
	do := true
	for _, line := range lines {
		index, n := 0, len(line)
		for index <= len(line) {
			mulIndex := rMul.FindStringIndex(line[index:n])
			doIndex := rDo.FindStringIndex(line[index:n])
			dontIndex := rDont.FindStringIndex(line[index:n])
			if mulIndex != nil && ((doIndex == nil || mulIndex[0] < doIndex[0]) && (dontIndex == nil || mulIndex[0] < dontIndex[0])) {
				if do {
					match := rMul.FindString(line[index:n])
					i := strings.Index(match, ",")
					firstNumber, _ := strconv.Atoi(match[4:i])
					secondNumber, _ := strconv.Atoi(match[i+1 : len(match)-1])
					result += firstNumber * secondNumber
				}
				index += mulIndex[1]
			} else if doIndex != nil && ((mulIndex == nil || doIndex[0] < mulIndex[0]) && (dontIndex == nil || doIndex[0] < dontIndex[0])) {
				do = true
				index += doIndex[1]
			} else if dontIndex != nil && ((mulIndex == nil || dontIndex[0] < mulIndex[0]) && (doIndex == nil || dontIndex[0] < doIndex[0])) {
				do = false
				index += dontIndex[1]
			} else {
				break
			}
		}
	}
	return result
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
