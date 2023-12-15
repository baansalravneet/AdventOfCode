package day8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Day8() {
	fmt.Println("--- Day 8: Haunted Wasteland ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart2Answer(lines []string) int {
	sequence := lines[0]
	adjList := getAdjList(lines[2:])
	startingNodes := []string{}
	for k := range adjList {
		if k[2] == 'A' {
			startingNodes = append(startingNodes, k)
		}
	}
	nodeAnswers := []int{}
	for _, node := range startingNodes {
		index := 0
		current := node
		count := 0
		visited := make(map[string]bool)
		for {
			visited[current+strconv.Itoa(index)] = true
			code := sequence[index]
			count++
			if code == 'L' {
				current = adjList[current][0]
			} else {
				current = adjList[current][1]
			}
			index++
			if index == len(sequence) {
				index = 0
			}
			if _, ok := visited[current+strconv.Itoa(index)]; ok { // we found a cycle
				break
			}
			if current[2] == 'Z' {
				nodeAnswers = append(nodeAnswers, count)
			}
		}
	}
	return getLCM(nodeAnswers)
}

func getLCM(nums []int) int {
	lcm := 1
	for _, i := range nums {
		lcm = (lcm * i) / getGCD(lcm, i)
	}
	return lcm
}

func getGCD(a, b int) int {
	if a > b {
		return getGCD(b, a)
	}
	if a == 0 {
		return b
	}
	return getGCD(b%a, a)
}

func getPart1Answer(lines []string) int {
	sequence := lines[0]
	adjList := getAdjList(lines[2:])
	answer := 0
	index := 0
	current := "AAA"
	for {
		if current == "ZZZ" {
			return answer
		}
		answer++
		code := sequence[index]
		index += 1
		index %= len(sequence)
		if code == 'L' {
			current = adjList[current][0]
		} else {
			current = adjList[current][1]
		}
	}
}

func getAdjList(lines []string) map[string][]string {
	result := make(map[string][]string)
	for _, line := range lines {
		key := line[:3]
		result[key] = []string{line[7:10], line[12:15]}
	}
	return result
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day8/input.txt")
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
