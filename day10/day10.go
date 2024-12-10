package day10

import (
	"bufio"
	"fmt"
	"os"
)

func Day10() {
	fmt.Println("--- Day 10: Hoof It ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(lines []string) int {
	grid := getGrid(lines)
	answer := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				visited := make([][]bool, len(grid))
				for i := range visited {
					visited[i] = make([]bool, len(grid[i]))
				}
				answer += dfs(grid, i, j, visited)
			}
		}
	}
	return answer
}

func getPart2Answer(lines []string) int {
	grid := getGrid(lines)
	answer := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				answer += dfsNoVisit(grid, i, j)
			}
		}
	}
	return answer
}

func dfs(grid [][]int, i, j int, visited [][]bool) int {
	if grid[i][j] == 9 {
		visited[i][j] = true
		return 1
	}
	if visited[i][j] {
		return 0
	}
	visited[i][j] = true
	answer := 0
	for _, d := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		ni, nj := i+d[0], j+d[1]
		if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[i]) {
			continue
		}
		if visited[ni][nj] {
			continue
		}
		if grid[ni][nj] != grid[i][j]+1 {
			continue
		}
		answer += dfs(grid, ni, nj, visited)
	}
	return answer
}

func getGrid(lines []string) [][]int {
	grid := [][]int{}
	for _, l := range lines {
		row := []int{}
		for _, c := range l {
			row = append(row, int(c-'0'))
		}
		grid = append(grid, row)
	}
	return grid
}

func dfsNoVisit(grid [][]int, i, j int) int {
	if grid[i][j] == 9 {
		return 1
	}
	answer := 0
	for _, d := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		ni, nj := i+d[0], j+d[1]
		if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[i]) {
			continue
		}
		if grid[ni][nj] != grid[i][j]+1 {
			continue
		}
		answer += dfsNoVisit(grid, ni, nj)
	}
	return answer
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
