package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day7() {
	fmt.Println("--- Day 7: Bridge Repair ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(lines []string) int {
	return getAnswer(lines, true)
}

func getPart2Answer(lines []string) int {
	return getAnswer(lines, false)
}

func possible(nums []int, target, current, index int) bool {
	if index >= len(nums) {
		return target == current
	}
	return possible(nums, target, current+nums[index], index+1) ||
		possible(nums, target, current*nums[index], index+1)
}

func possible2(nums []int, target, current, index int) bool {
	if index >= len(nums) {
		return target == current
	}
	return possible2(nums, target, current+nums[index], index+1) ||
		possible2(nums, target, current*nums[index], index+1) ||
		possible2(nums, target, concat(current, nums[index]), index+1)
}

func concat(a, b int) int {
	digitCount := 1
	tempB := b
	for tempB != 0 {
		digitCount *= 10
		tempB /= 10
	}
	return a*digitCount + b
}

func getAnswer(lines []string, part1 bool) int {
	answer := 0
	for _, line := range lines {
		result, _ := strconv.Atoi(line[0:strings.Index(line, ":")])
		numS := strings.Split(line[strings.Index(line, ":")+2:len(line)], " ")
		num := []int{}
		for _, n := range numS {
			val, _ := strconv.Atoi(n)
			num = append(num, val)
		}
		if part1 {
			if possible(num, result, num[0], 1) {
				answer += result
			}
		} else {
			if possible2(num, result, num[0], 1) {
				answer += result
			}
		}
	}
	return answer
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
