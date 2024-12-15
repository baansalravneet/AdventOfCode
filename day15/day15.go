package day15

import (
	"bufio"
	"fmt"
	"os"
)

func Day15() {
	fmt.Println("--- Day 15: Warehouse Woes ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

var dir [][]int = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func getGrid(lines []string) [][]byte {
	grid := [][]byte{}
	for _, line := range lines {
		if line == "" {
			break
		}
		grid = append(grid, []byte(line))
	}
	return grid
}

func getGrid2(lines []string) [][]byte {
	grid := [][]byte{}
	for _, line := range lines {
		if line == "" {
			break
		}
		row := []byte{}
		for _, c := range line {
			if c == '#' {
				row = append(row, '#', '#')
			} else if c == 'O' {
				row = append(row, '[', ']')
			} else if c == '.' {
				row = append(row, '.', '.')
			} else if c == '@' {
				row = append(row, '@', '.')
			}
		}
		grid = append(grid, row)
	}
	return grid
}

func robotLocation(grid [][]byte) (int, int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '@' {
				return i, j
			}
		}
	}
	return -1, -1
}

func getCommands(lines []string, index int) []int {
	commands := []int{}
	for ; index < len(lines); index++ {
		for _, c := range lines[index] {
			switch c {
			case '^':
				commands = append(commands, 0)
			case '>':
				commands = append(commands, 1)
			case 'v':
				commands = append(commands, 2)
			case '<':
				commands = append(commands, 3)
			}
		}
	}
	return commands
}

func getScore(grid [][]byte, c byte) int {
	answer := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == c {
				answer += 100*i + j
			}
		}
	}
	return answer
}

func canMove(grid [][]byte, x, y, c int) bool {
	for {
		if grid[x][y] == '#' {
			return false
		}
		if grid[x][y] == '.' {
			return true
		}
		x, y = x+dir[c][0], y+dir[c][1]
	}
}

func canMove2(grid [][]byte, x, y, c int) bool {
	if c == 1 || c == 3 {
		return canMove(grid, x, y, c)
	}
	nx, ny := x+dir[c][0], y+dir[c][1]
	next := grid[nx][ny]
	switch next {
	case '#':
		return false
	case '[':
		return canMove2(grid, nx, ny, c) && canMove2(grid, nx, ny+1, c)
	case ']':
		return canMove2(grid, nx, ny, c) && canMove2(grid, nx, ny-1, c)
	}
	return true
}

func move(grid [][]byte, x, y, c int) {
	if grid[x][y] == '.' {
		return
	}
	nx, ny := x+dir[c][0], y+dir[c][1]
	move(grid, nx, ny, c)
	grid[nx][ny], grid[x][y] = grid[x][y], '.'
}

func move2(grid [][]byte, x, y, c int) {
	if c == 1 || c == 3 {
		move(grid, x, y, c)
		return
	}
	if grid[x][y] == '.' {
		return
	}
	nx, ny := x+dir[c][0], y+dir[c][1]
	if grid[x][y] == '[' {
		move2(grid, nx, ny, c)
		move2(grid, nx, ny+1, c)
		grid[nx][ny], grid[nx][ny+1] = grid[x][y], grid[x][y+1]
		grid[x][y], grid[x][y+1] = '.', '.'
	} else if grid[x][y] == ']' {
		move2(grid, nx, ny, c)
		move2(grid, nx, ny-1, c)
		grid[nx][ny], grid[nx][ny-1] = grid[x][y], grid[x][y-1]
		grid[x][y], grid[x][y-1] = '.', '.'
	} else if grid[x][y] == '@' {
		move2(grid, nx, ny, c)
		grid[nx][ny], grid[x][y] = grid[x][y], '.'
	}
}

func getPart1Answer(lines []string) int {
	grid := getGrid(lines)
	x, y := robotLocation(grid)
	commands := getCommands(lines, len(grid)+1)
	for _, c := range commands {
		if canMove(grid, x, y, c) {
			move(grid, x, y, c)
			x, y = x+dir[c][0], y+dir[c][1]
		}
	}
	return getScore(grid, 'O')
}

func getPart2Answer(lines []string) int {
	grid := getGrid2(lines)
	x, y := robotLocation(grid)
	commands := getCommands(lines, len(grid)+1)
	for _, c := range commands {
		if canMove2(grid, x, y, c) {
			move2(grid, x, y, c)
			x, y = x+dir[c][0], y+dir[c][1]
		}
	}
	return getScore(grid, '[')
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day15/input.txt")
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
