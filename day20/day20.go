package day20

import (
	"bufio"
	"fmt"
	"os"
)

func Day20() {
	fmt.Println("--- Day 20: Race Condition ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(lines []string) int {
	grid := getGrid(lines)
	x, y := getStart(grid)
	distance := getDistances(grid, x, y)
	return getCount(distance, 2, 100)
}

func getPart2Answer(lines []string) int {
	grid := getGrid(lines)
	x, y := getStart(grid)
	distance := getDistances(grid, x, y)
	return getCount(distance, 20, 100)
}

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func getGrid(lines []string) [][]byte {
	grid := [][]byte{}
	for _, line := range lines {
		grid = append(grid, []byte(line))
	}
	return grid
}

func getStart(grid [][]byte) (int, int) {
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 'S' {
				return x, y
			}
		}
	}
	return -1, -1
}

func getCount(distance [][]int, maxSteps, saved int) int {
	count := 0
	for x := range distance {
		for y := range distance[x] {
			if distance[x][y] == -1 {
				continue
			}
			count += getSubCount(distance, x, y, maxSteps, saved)
		}
	}
	return count
}

func getSubCount(distance [][]int, x, y, maxSteps, saved int) int {
	count := 0
	for i := range distance {
		for j := range distance[i] {
			if distance[i][j] == -1 {
				continue
			}
			if steps := getSteps(x, y, i, j); steps <= maxSteps {
				if distance[i][j]-distance[x][y]-steps >= saved {
					count++
				}
			}
		}
	}
	return count
}

func getSteps(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func getDistances(grid [][]byte, x, y int) [][]int {
	distance := make([][]int, len(grid))
	for i := range distance {
		distance[i] = make([]int, len(grid[i]))
		for j := range distance[i] {
			distance[i][j] = -1
		}
	}
	current := 0
	prevX, prevY := -1, -1
	for grid[x][y] != 'E' {
		distance[x][y] = current
		for _, d := range directions {
			nx, ny := x+d[0], y+d[1]
			if nx == prevX && ny == prevY {
				continue
			}
			if grid[nx][ny] == '#' {
				continue
			}
			prevX, prevY = x, y
			x, y = nx, ny
			break
		}
		current++
	}
	distance[x][y] = current
	return distance
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day20/input.txt")
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
