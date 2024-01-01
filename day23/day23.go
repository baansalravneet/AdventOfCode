package day23

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day23() {
	fmt.Println("--- Day 23: A Long Walk ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(lines []string) int {
	start := strings.Index(lines[0], ".")
	result := 0
	var dfs func(int, int, int, int, int)
	dfs = func(x, y, px, py, length int) {
		if x == len(lines)-1 {
			result = max(result, length)
			return
		}
		for i, next := range [][]int{{x + 1, y}, {x - 1, y}, {x, y + 1}, {x, y - 1}} {
			nx, ny := next[0], next[1]
			if nx < 0 || ny < 0 || nx >= len(lines) || ny >= len(lines[0]) {
				continue
			}
			if lines[nx][ny] == '#' {
				continue
			}
			if nx == px && ny == py {
				continue
			}
			if i == 0 && lines[nx][ny] == '^' {
				continue
			}
			if i == 1 && lines[nx][ny] == 'v' {
				continue
			}
			if i == 2 && lines[nx][ny] == '<' {
				continue
			}
			if i == 3 && lines[nx][ny] == '>' {
				continue
			}
			dfs(nx, ny, x, y, length+1)
		}
	}
	dfs(0, start, -1, -1, 0)
	return result
}

func getPart2Answer(lines []string) int {
	edges := getEdges(lines)
	start := strings.Index(lines[0], ".")
	visited := make(map[[2]int]bool)
	result := 0
	var dfs func(int, int, int)
	dfs = func(x, y, distance int) {
		key := [2]int{x, y}
		if x == len(lines)-1 {
			result = max(result, distance)
		}
		visited[key] = true
		next := edges[key]
		for _, edge := range next {
			nx, ny, d := edge[0], edge[1], edge[2]
			if _, ok := visited[[2]int{nx, ny}]; ok {
				continue
			}
			dfs(nx, ny, distance+d)
		}
		delete(visited, key)
	}
	dfs(0, start, 0)
	return result
}

func getEdges(lines []string) map[[2]int][][3]int {
	result := make(map[[2]int][][3]int)
	poi := getPointsOfInterest(lines)
	for k := range poi {
		result[k] = append(result[k], getPaths(k[0], k[1], poi, lines)...)
	}
	return result
}

func getPaths(x, y int, poi map[[2]int]bool, lines []string) [][3]int {
	visited := make(map[[2]int]bool)
	result := [][3]int{}
	var dfs func(int, int, int)
	dfs = func(x, y, distance int) {
		if _, ok := poi[[2]int{x, y}]; ok && distance != 0 {
			result = append(result, [3]int{x, y, distance})
			return
		}
		visited[[2]int{x, y}] = true
		for _, d := range [][]int{{x + 1, y}, {x - 1, y}, {x, y + 1}, {x, y - 1}} {
			nx, ny := d[0], d[1]
			if nx < 0 || nx >= len(lines) || ny < 0 || ny >= len(lines[0]) {
				continue
			}
			if lines[nx][ny] == '#' {
				continue
			}
			if _, ok := visited[[2]int{nx, ny}]; ok {
				continue
			}
			dfs(nx, ny, distance+1)
		}
	}
	dfs(x, y, 0)
	return result
}

func getPointsOfInterest(lines []string) map[[2]int]bool {
	result := make(map[[2]int]bool)
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == '#' {
				continue
			}
			if i == 0 || i == len(lines)-1 || getNeighbours(lines, i, j) >= 3 {
				result[[2]int{i, j}] = true
			}
		}
	}
	return result
}

func getNeighbours(lines []string, x, y int) int {
	count := 0
	for _, d := range [][]int{{x + 1, y}, {x - 1, y}, {x, y + 1}, {x, y - 1}} {
		nx, ny := d[0], d[1]
		if nx < 0 || ny < 0 || nx >= len(lines) || ny >= len(lines[0]) {
			continue
		}
		if lines[nx][ny] == '#' {
			continue
		}
		count++
	}
	return count
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day23/input.txt")
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
