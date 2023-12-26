package day16

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day16() {
	fmt.Println("--- Day 16: The Floor Will Be Lava ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(lines []string) int {
	return runBFS(lines, "0,0,>")
}

func runBFS(lines []string, start string) int {
	visited := make(map[string]bool)
	visitedCoords := make(map[string]bool)
	bfsQueue := []string{}
	bfsQueue = append(bfsQueue, start)
	for len(bfsQueue) > 0 {
		current := bfsQueue[0]
		bfsQueue = bfsQueue[1:]
		visitedCoords[getCoords(current)] = true
		visited[current] = true
		for _, next := range getNext(current, lines) {
			if _, ok := visited[next]; !ok && isValid(next, len(lines)) {
				bfsQueue = append(bfsQueue, next)
			}
		}
	}
	return len(visitedCoords)
}

func getCoords(s string) string {
	return strings.Split(s, ",")[0] + " " + strings.Split(s, ",")[1]
}

func isValid(s string, n int) bool {
	x, _ := strconv.Atoi(strings.Split(s, ",")[0])
	y, _ := strconv.Atoi(strings.Split(s, ",")[1])
	return x >= 0 && x < n && y >= 0 && y < n
}

func getNext(current string, lines []string) []string {
	result := []string{}
	x, _ := strconv.Atoi(strings.Split(current, ",")[0])
	y, _ := strconv.Atoi(strings.Split(current, ",")[1])
	d := strings.Split(current, ",")[2]
	switch lines[x][y] {
	case '.':
		switch d {
		case ">":
			return append(result, fmt.Sprintf("%d,%d,%s", x, y+1, d))
		case "<":
			return append(result, fmt.Sprintf("%d,%d,%s", x, y-1, d))
		case "v":
			return append(result, fmt.Sprintf("%d,%d,%s", x+1, y, d))
		case "^":
			return append(result, fmt.Sprintf("%d,%d,%s", x-1, y, d))
		}
	case '\\':
		switch d {
		case ">":
			return append(result, fmt.Sprintf("%d,%d,%s", x+1, y, "v"))
		case "<":
			return append(result, fmt.Sprintf("%d,%d,%s", x-1, y, "^"))
		case "v":
			return append(result, fmt.Sprintf("%d,%d,%s", x, y+1, ">"))
		case "^":
			return append(result, fmt.Sprintf("%d,%d,%s", x, y-1, "<"))
		}
	case '/':
		switch d {
		case ">":
			return append(result, fmt.Sprintf("%d,%d,%s", x-1, y, "^"))
		case "<":
			return append(result, fmt.Sprintf("%d,%d,%s", x+1, y, "v"))
		case "v":
			return append(result, fmt.Sprintf("%d,%d,%s", x, y-1, "<"))
		case "^":
			return append(result, fmt.Sprintf("%d,%d,%s", x, y+1, ">"))
		}
	case '-':
		switch d {
		case ">":
			return append(result, fmt.Sprintf("%d,%d,%s", x, y+1, d))
		case "<":
			return append(result, fmt.Sprintf("%d,%d,%s", x, y-1, d))
		case "v", "^":
			result = append(result, fmt.Sprintf("%d,%d,%s", x, y-1, "<"))
			return append(result, fmt.Sprintf("%d,%d,%s", x, y+1, ">"))
		}
	case '|':
		switch d {
		case ">", "<":
			result = append(result, fmt.Sprintf("%d,%d,%s", x+1, y, "v"))
			return append(result, fmt.Sprintf("%d,%d,%s", x-1, y, "^"))
		case "v":
			return append(result, fmt.Sprintf("%d,%d,%s", x+1, y, d))
		case "^":
			return append(result, fmt.Sprintf("%d,%d,%s", x-1, y, d))
		}
	}
	return result
}

func getPart2Answer(lines []string) int {
	result := 0
	n := len(lines)
	for i := 0; i < n; i++ {
		result = max(result, runBFS(lines, fmt.Sprintf("0,%d,v", i)))
		result = max(result, runBFS(lines, fmt.Sprintf("%d,%d,^", len(lines)-1, i)))
		result = max(result, runBFS(lines, fmt.Sprintf("%d,0,>", i)))
		result = max(result, runBFS(lines, fmt.Sprintf("%d,%d,<", i, len(lines)-1)))
	}
	return result
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day16/input.txt")
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
