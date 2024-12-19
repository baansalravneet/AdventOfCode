package day18

import (
	"bufio"
	"fmt"
	"os"
)

func Day18() {
	fmt.Println("--- Day 18: RAM Run ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func getPart1(lines []string, byteLimit int) int {
	xBound, yBound := 71, 71
	corrupted := [][]bool{}
	for i := 0; i < xBound; i++ {
		corrupted = append(corrupted, make([]bool, yBound))
	}
	for i := 0; i < byteLimit; i++ {
		var x, y int
		fmt.Sscanf(lines[i], "%d,%d", &x, &y)
		corrupted[x][y] = true
	}
	steps := 0
	bfsQ := [][2]int{}
	bfsQ = append(bfsQ, [2]int{0, 0})
	for len(bfsQ) > 0 {
		count := len(bfsQ)
		for count > 0 {
			count--
			x, y := bfsQ[0][0], bfsQ[0][1]
			bfsQ = bfsQ[1:]
			if x == xBound-1 && y == yBound-1 {
				return steps
			}
			if corrupted[x][y] {
				continue
			}
			corrupted[x][y] = true
			for _, d := range directions {
				nx, ny := x+d[0], y+d[1]
				if nx < 0 || ny < 0 || nx >= xBound || ny >= yBound {
					continue
				}
				if corrupted[nx][ny] {
					continue
				}
				bfsQ = append(bfsQ, [2]int{nx, ny})
			}
		}
		steps++
	}
	return -1
}

func getPart1Answer(lines []string) int {
	return getPart1(lines, 1024)
}

func getPart2Answer(lines []string) string {
	left, right := 0, len(lines)-1
	answer := -1
	for left <= right {
		mid := (left + right) / 2
		if getPart1(lines, mid) == -1 {
			answer = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return lines[answer-1]
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day18/input.txt")
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
