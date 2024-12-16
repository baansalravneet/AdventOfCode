package day12

import (
	"bufio"
	"fmt"
	"os"
)

func Day12() {
	fmt.Println("--- Day 12: Garden Groups ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getDimensions(grid [][]byte, i, j int, target byte) (int, int) {
	if grid[i][j] != target {
		return 0, 0
	}
	grid[i][j] = '_'
	area := 1
	perimeter := 0
	for _, d := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		ni, nj := i+d[0], j+d[1]
		if ni < 0 || nj < 0 || ni >= len(grid) || nj >= len(grid[0]) {
			perimeter++
			continue
		}
		if grid[ni][nj] != '_' && grid[ni][nj] != target {
			perimeter++
			continue
		}
		na, np := getDimensions(grid, ni, nj, target)
		area += na
		perimeter += np
	}
	return area, perimeter
}

func floodFill(grid [][]byte, i, j int) {
	if grid[i][j] != '_' {
		return
	}
	grid[i][j] = '0'
	for _, d := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		ni, nj := i+d[0], j+d[1]
		if ni < 0 || nj < 0 || ni >= len(grid) || nj >= len(grid[0]) {
			continue
		}
		floodFill(grid, ni, nj)
	}
}

func getGrid(lines []string) [][]byte {
	grid := [][]byte{}
	for _, line := range lines {
		grid = append(grid, []byte(line))
	}
	return grid
}

func checkShare(boundarySet map[[3]int]bool, key [3]int) bool {
	if key[2] == 0 || key[2] == 1 {
		nKey1 := [3]int{key[0] - 1, key[1], key[2]}
		nKey2 := [3]int{key[0] + 1, key[1], key[2]}
		if boundarySet[nKey1] || boundarySet[nKey2] {
			return true
		}
	} else {
		nKey1 := [3]int{key[0], key[1] - 1, key[2]}
		nKey2 := [3]int{key[0], key[1] + 1, key[2]}
		if boundarySet[nKey1] || boundarySet[nKey2] {
			return true
		}
	}
	return false
}

func getSides(grid [][]byte, i, j int) int {
	count := 0
	boundarySet := make(map[[3]int]bool)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '_' {
				for idx, d := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
					ni, nj := i+d[0], j+d[1]
					if ni >= 0 && nj >= 0 && ni < len(grid) && nj < len(grid[0]) {
						if grid[ni][nj] == '_' {
							continue
						}
					}
					key := [3]int{i, j, idx}
					if !checkShare(boundarySet, key) {
						count++
					}
					boundarySet[key] = true
				}
			}
		}
	}
	return count
}

func getPart1Answer(lines []string) int {
	grid := getGrid(lines)
	cost := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				continue
			}
			area, perimeter := getDimensions(grid, i, j, grid[i][j])
			floodFill(grid, i, j)
			cost += perimeter * area
		}
	}
	return cost
}

func getPart2Answer(lines []string) int {
	grid := getGrid(lines)
	cost := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				continue
			}
			area, _ := getDimensions(grid, i, j, grid[i][j])
			sides := getSides(grid, i, j)
			floodFill(grid, i, j)
			cost += sides * area
		}
	}
	return cost
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day12/input.txt")
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
