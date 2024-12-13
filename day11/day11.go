package day11

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day11() {
	fmt.Println("--- Day 11: Plutonian Pebbles ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func splitDigits(num, count int) (int, int) {
	intString := strconv.Itoa(num)
	a, _ := strconv.Atoi(intString[:count/2])
	b, _ := strconv.Atoi(intString[count/2:])
	return a, b
}

func digitCount(num int) int {
	count := 0
	for num != 0 {
		num /= 10
		count++
	}
	return count
}

func recurseCached(num, blinks int, cache map[[2]int]int) int {
	if blinks == 0 {
		return 1
	}
	if val, ok := cache[[2]int{num, blinks}]; ok {
		return val
	}
	answer := 0
	if num == 0 {
		answer = recurseCached(1, blinks-1, cache)
	} else if count := digitCount(num); count%2 == 0 {
		a, b := splitDigits(num, count)
		answer = recurseCached(a, blinks-1, cache) + recurseCached(b, blinks-1, cache)
	} else {
		answer = recurseCached(num*2024, blinks-1, cache)
	}
	cache[[2]int{num, blinks}] = answer
	return answer
}

func getAnswer(line string, blinks int) int {
    nums := strings.Split(line, " ")
    cache := make(map[[2]int]int)
    answer := 0
    for _, num := range nums {
        numInt, _ := strconv.Atoi(num)
        answer += recurseCached(numInt, blinks, cache)
    }
    return answer
}

func getPart1Answer(lines []string) int {
    return getAnswer(lines[0], 25)
}

func getPart2Answer(lines []string) int {
    return getAnswer(lines[0], 75)
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day11/input.txt")
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
