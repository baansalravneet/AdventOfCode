package day10

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func Day10() {
	fmt.Println("--- Day 10: Pipe Maze ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getVisited(lines []string) [][]bool {
	visited := make([][]bool, len(lines))
	for i := range lines {
		visited[i] = make([]bool, len(lines[i]))
	}
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == 'S' {
				visited[i][j] = true
				for _, neighbour := range [][]int{{i + 1, j}, {i - 1, j}, {i, j + 1}, {i, j - 1}} {
					n := getNext(neighbour[0], neighbour[1], &lines)
					if contains(n, []int{i, j}) {
						findLoop(neighbour[0], neighbour[1], &visited, &lines)
						break
					}
				}
			}
		}
	}
	return visited
}

func getPart1Answer(lines []string) int {
	visited := getVisited(lines)
	length := 0
	for _, v := range visited {
		for _, b := range v {
			if b {
				length++
			}
		}
	}
	return length / 2
}

func contains(arr [][]int, target []int) bool {
	for _, i := range arr {
		if target[0] == i[0] && target[1] == i[1] {
			return true
		}
	}
	return false
}

func getNeighbours(c byte) [][]int {
	switch c {
	case 'F':
		return [][]int{{0, 1}, {1, 0}}
	case 'J':
		return [][]int{{0, -1}, {-1, 0}}
	case 'L':
		return [][]int{{0, 1}, {-1, 0}}
	case '7':
		return [][]int{{1, 0}, {0, -1}}
	case '|':
		return [][]int{{-1, 0}, {1, 0}}
	case '-':
		return [][]int{{0, 1}, {0, -1}}
	}
	return [][]int{}
}

func getNext(x, y int, lines *[]string) [][]int {
	result := [][]int{}
	if x < 0 || y < 0 {
		return result
	}
	if x >= len(*lines) || y >= len((*lines)[0]) {
		return result
	}
	for _, next := range getNeighbours((*lines)[x][y]) {
		if next[0]+x < 0 || next[1]+y < 0 {
			continue
		}
		if next[0]+x >= len(*lines) || next[1]+y >= len((*lines)[x]) {
			continue
		}
		result = append(result, []int{next[0] + x, next[1] + y})
	}
	return result
}

func findLoop(x, y int, visited *[][]bool, lines *[]string) {
	if (*visited)[x][y] || (*lines)[x][y] == 'S' {
		return
	}
	(*visited)[x][y] = true
	for _, next := range getNext(x, y, lines) {
		findLoop(next[0], next[1], visited, lines)
	}
}

func countInversions(x, y int, lines *[]string, visited *[][]bool) int {
	count := 0
	for j := 0; j < y; j++ {
		if !(*visited)[x][j] {
			continue
		}
		if slices.Contains([]byte{'|', 'F', '7'}, (*lines)[x][j]) {
			count++
		}
	}
	return count
}

func getPart2Answer(lines []string) int {
	visited := getVisited(lines)

	result := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if visited[i][j] {
				continue
			}
			count := countInversions(i, j, &lines, &visited)
			if count%2 != 0 {
				result++
			}
		}
	}
	return result
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day10/input.txt")
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
