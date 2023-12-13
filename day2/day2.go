package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2() {
	fmt.Println("--- Day 2: Cube Conundrum ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart2Answer(lines []string) int {
	result := 0
	for _, line := range lines {
		result += getPower(line)
	}
	return result
}

func getPower(line string) int {
	gameInfo := strings.Split(line, ":")
	hands := strings.Split(gameInfo[1], ";")
	maxBlue := 0
	maxGreen := 0
	maxRed := 0
	for _, hand := range hands {
		numbers := getDieNumber(hand)
		maxBlue = max(maxBlue, numbers[0])
		maxGreen = max(maxGreen, numbers[1])
		maxRed = max(maxRed, numbers[2])
	}
	return maxBlue * maxGreen * maxRed
}

func getPart1Answer(lines []string) int {
	result := 0
	for _, line := range lines {
		if game_id, ok := isPossible(line); ok {
			result += game_id
		}
	}
	return result
}

func isPossible(line string) (int, bool) {
	gameInfo := strings.Split(line, ":")
	gameId, _ := strconv.Atoi(strings.Split(gameInfo[0], " ")[1])
	hands := strings.Split(gameInfo[1], ";")
	for _, hand := range hands {
		numbers := getDieNumber(hand)
		if numbers[0] > 14 || numbers[1] > 13 || numbers[2] > 12 {
			return gameId, false
		}
	}
	return gameId, true
}

func getDieNumber(hand string) [3]int {
	answer := [3]int{0, 0, 0}
	dieSplit := strings.Split(hand, ",")
	for _, split := range dieSplit {
		parts := strings.Split(split, " ")
		num := parts[1]
		colour := parts[2]
		if colour == "blue" {
			number, _ := strconv.Atoi(num)
			answer[0] += number
		}
		if colour == "green" {
			number, _ := strconv.Atoi(num)
			answer[1] += number
		}
		if colour == "red" {
			number, _ := strconv.Atoi(num)
			answer[2] += number
		}
	}
	return answer
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day2/input.txt")
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
