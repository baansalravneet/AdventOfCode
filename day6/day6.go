package day6

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day6() {
	fmt.Println("--- Day 6: Wait For It ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart2Answer(lines []string) int {
	time := getPart2Value(lines[0])
	distance := getPart2Value(lines[1])
	return getResult(time, distance)
}

func getPart2Value(line string) int {
	line = strings.Split(line, ":")[1]
	line = strings.Trim(line, " ")
	re := regexp.MustCompile("\\s+")
	resultString := ""
	for _, w := range re.Split(line, -1) {
		resultString += w
	}
	result, _ := strconv.Atoi(resultString)
	return result
}

func getPart1Answer(lines []string) int {
	times := getValues(lines[0])
	distances := getValues(lines[1])
	result := 1
	for i, time := range times {
		result *= getResult(time, distances[i])
	}
	return result
}

func getResult(time, distance int) int {
	minTimeHeld := getMinTimeHeld(time, distance)
	maxTimeHeld := getMaxTimeHeld(time, distance)
	if minTimeHeld == -1 {
		return 0
	}
	return maxTimeHeld - minTimeHeld + 1
}

func getMaxTimeHeld(time, distance int) int {
	for i := time; i >= 0; i-- {
		if (time-i)*i > distance {
			return i
		}
	}
	return -1
}

func getMinTimeHeld(time, distance int) int {
	for i := 0; i <= time; i++ {
		if (time-i)*i > distance {
			return i
		}
	}
	return -1
}

func winnable(timeHeld, totalTime, distance int) bool {
	speed := timeHeld
	timeLeft := totalTime - timeHeld
	return distance <= speed*timeLeft
}

func getValues(line string) []int {
	values := []int{}
	line = strings.Split(line, ":")[1]
	line = strings.Trim(line, " ")
	re := regexp.MustCompile("\\s+")
	for _, w := range re.Split(line, -1) {
		value, _ := strconv.Atoi(w)
		values = append(values, value)
	}
	return values
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day6/input.txt")
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
