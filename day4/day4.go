package day4

import (
	"bufio"
	"fmt"
	"os"
)

func Day4() {
	fmt.Println("--- Day 4: Ceres Search ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(lines []string) int {
	grid := [][]rune{}
	for _, line := range lines {
		row := []rune{}
		for _, c := range line {
			row = append(row, c)
		}
		grid = append(grid, row)
	}
	answer := 0
	directions := [][]int{
		{1, 0}, {-1, 0}, {1, 1}, {-1, 1},
		{0, 1}, {1, -1}, {-1, -1}, {0, -1},
	}
	word := []rune{'M', 'A', 'S'}
	getCount := func(i, j int) int {
		count := 0
		for _, d := range directions {
			w, ni, nj := 0, i, j
			for ; w < 3; w++ {
				ni, nj = ni+d[0], nj+d[1]
				if ni < 0 || ni >= len(grid) {
					break
				}
				if nj < 0 || nj >= len(grid[0]) {
					break
				}
				if grid[ni][nj] != word[w] {
					break
				}
			}
			if w == 3 {
				count++
			}
		}
		return count
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'X' {
				answer += getCount(i, j)
			}
		}
	}
	return answer
}

func getPart2Answer(lines []string) int {
	grid := [][]rune{}
	for _, line := range lines {
		row := []rune{}
		for _, c := range line {
			row = append(row, c)
		}
		grid = append(grid, row)
	}
	answer := 0
	directions := [][]int{{-1, -1}, {1, 1}, {1, -1}, {-1, 1}}
	check := func(i, j int) bool {
		mCount, sCount := 0, 0
		for _, d := range directions {
			ni, nj := i+d[0], j+d[1]
			if grid[ni][nj] == 'M' {
				mCount++
			} else if grid[ni][nj] == 'S' {
				sCount++
			}
		}
		if mCount != 2 || sCount != 2 {
			return false
		}
		return grid[i+1][j+1] != grid[i-1][j-1]
	}
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[1])-1; j++ {
			if grid[i][j] == 'A' && check(i, j) {
				answer++
			}
		}
	}
	return answer
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day4/input.txt")
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
