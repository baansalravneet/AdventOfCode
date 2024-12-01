package day1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day1() {
	fmt.Println("--- Day 1: Historian Hysteria ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(lines []string) int {
	leftList, rightList := []int{}, []int{}
	for _, l := range lines {
		numbers := strings.Split(l, "   ")
		left, _ := strconv.Atoi(numbers[0])
		right, _ := strconv.Atoi(numbers[1])
		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}
	sort.Ints(leftList)
	sort.Ints(rightList)
	var result int
	for i := range leftList {
		result += int(math.Abs(float64(leftList[i] - rightList[i])))
	}
	return result
}

func getPart2Answer(lines []string) int {
	leftList, rightMap := []int{}, make(map[int]int)
	for _, l := range lines {
		numbers := strings.Split(l, "   ")
		left, _ := strconv.Atoi(numbers[0])
		right, _ := strconv.Atoi(numbers[1])
		leftList = append(leftList, left)
		rightMap[right]++
	}
	var result int
	for _, v := range leftList {
        result += v * rightMap[v]
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
