package day6

import (
	"bufio"
	"fmt"
	"os"
)

func Day6() {
	fmt.Println("--- Day 6: Guard Gallivant ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(lines []string) int {
	grid := [][]rune{}
	seen := [][]bool{}
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	i, j, dir, count := -1, -1, 0, 0
	for x, line := range lines {
		row := []rune{}
		for y, c := range line {
			row = append(row, c)
			if c == '^' {
				i, j = x, y
				row[y] = 'X'
			}
		}
		grid = append(grid, row)
		seen = append(seen, make([]bool, len(row)))
	}
	for {
		if !seen[i][j] {
			count++
		}
		seen[i][j] = true
		ni, nj := i+directions[dir][0], j+directions[dir][1]
		if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[0]) {
			return count
		}
		if grid[ni][nj] == '#' {
			dir = (dir + 1) % 4
		} else {
			i, j = ni, nj
		}
	}
}

func getPart2Answer(lines []string) int {
	grid := [][]rune{}
	seen := [][]bool{}
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	si, sj, dir, count := -1, -1, 0, 0
	for x, line := range lines {
		row := []rune{}
		for y, c := range line {
			row = append(row, c)
			if c == '^' {
				si, sj = x, y
				row[y] = 'X'
			}
		}
		grid = append(grid, row)
		seen = append(seen, make([]bool, len(row)))
	}
	populateSeen := func(i, j, dir int) {
		for {
			seen[i][j] = true
			ni, nj := i+directions[dir][0], j+directions[dir][1]
			if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[0]) {
				return
			}
			if grid[ni][nj] == '#' {
				dir = (dir + 1) % 4
			} else {
				i, j = ni, nj
			}
		}
	}
	checkLoop := func(i, j, dir int) bool {
		newSeen := make([][][]int, len(grid))
		for i := range seen {
			newSeen[i] = make([][]int, len(grid[0]))
		}
		for {
			for _, s := range newSeen[i][j] {
				if s == dir {
					return true
				}
			}
			newSeen[i][j] = append(newSeen[i][j], dir)
			ni, nj := i+directions[dir][0], j+directions[dir][1]
			if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[0]) {
				return false
			}
			if grid[ni][nj] == '#' {
				dir = (dir + 1) % 4
			} else {
				i, j = ni, nj
			}
		}
	}
	populateSeen(si, sj, dir)
	for x := range seen {
		for y := range seen[x] {
			if !seen[x][y] {
				continue
			}
			if x == si && y == sj {
				continue
			}
			grid[x][y] = '#'
			if checkLoop(si, sj, 0) {
				count++
			}
			grid[x][y] = '.'
		}
	}
	return count
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day6/input.txt")
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
