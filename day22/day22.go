package day22

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Day22() {
	fmt.Println("--- Day 22: Monkey Market ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getNext(num int) int {
	num = (num ^ (num << 6)) & ((1 << 24) - 1)
	num = (num ^ (num >> 5)) & ((1 << 24) - 1)
	return (num ^ (num << 11)) & ((1 << 24) - 1)
}

func getAllSequences() map[[4]int]bool {
	result := make(map[[4]int]bool)
	for i := 0; i <= 99999; i++ {
		result[getSequence(i)] = true
	}
	return result
}

func getSequence(i int) [4]int {
	return [4]int{
		((i / 1000) % 10) - ((i / 10000) % 10),
		((i / 100) % 10) - ((i / 1000) % 10),
		((i / 10) % 10) - ((i / 100) % 10),
		(i % 10) - ((i / 10) % 10),
	}
}

func getPart1Answer(lines []string) int {
	answer := 0
	for _, l := range lines {
		input, _ := strconv.Atoi(l)
		for i := 0; i < 2000; i++ {
			input = getNext(input)
		}
		answer += input
	}
	return answer
}

func getPart2Answer(lines []string) int {
	answer := -1
	lookup := []map[[4]int]int{}
	for _, l := range lines {
		input, _ := strconv.Atoi(l)
		sequence := []int{input % 10}
		for i := 0; i < 1999; i++ {
			next := getNext(input)
			sequence = append(sequence, next%10)
			input = next
		}
		valueMap := make(map[[4]int]int)
		for i := 0; i+4 < len(sequence); i++ {
			seq := [4]int{
				sequence[i+1] - sequence[i],
				sequence[i+2] - sequence[i+1],
				sequence[i+3] - sequence[i+2],
				sequence[i+4] - sequence[i+3],
			}
			if _, ok := valueMap[seq]; ok {
				continue
			}
			valueMap[seq] = sequence[i+4]
		}
		lookup = append(lookup, valueMap)
	}
	sequences := getAllSequences()
	for seq, _ := range sequences {
		totalPrice := 0
		for _, valueMap := range lookup {
			if v, ok := valueMap[seq]; ok {
				totalPrice += v
			}
		}
		answer = max(answer, totalPrice)
	}
	return answer
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day22/input.txt")
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
