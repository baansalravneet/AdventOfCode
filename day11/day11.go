package day11

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func Day11() {
	fmt.Println("--- Day 11: Cosmic Expansion ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(lines []string) int {
	return findDistances(lines, 2)
}

func findDistances(lines []string, scale int) int {
	emptyCols := getEmptyCols(lines)
	emptyRows := getEmptyRows(lines)
	points := getPoints(lines, emptyRows, emptyCols, scale)
	result := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			result += int(math.Abs(float64(points[i][0] - points[j][0])))
			result += int(math.Abs(float64(points[i][1] - points[j][1])))
		}
	}
	return result
}

func getPoints(lines []string, emptyRows, emptyCols []int, scale int) [][]int {
	result := [][]int{}
	rowIndex := 0
	for row := 0; row < len(lines); row++ {
		if rowIndex < len(emptyRows) && emptyRows[rowIndex] == row {
			rowIndex++
		}
		colIndex := 0
		for col := 0; col < len(lines[row]); col++ {
			if colIndex < len(emptyCols) && emptyCols[colIndex] == col {
				colIndex++
			}
			if lines[row][col] != '.' {
				result = append(result, []int{row + rowIndex*(scale-1), col + colIndex*(scale-1)})
			}
		}
	}
	return result
}

func getEmptyCols(lines []string) []int {
	result := []int{}
loop:
	for col := 0; col < len(lines[0]); col++ {
		for row := 0; row < len(lines); row++ {
			if lines[row][col] != '.' {
				continue loop
			}
		}
		result = append(result, col)
	}
	return result
}

func getEmptyRows(lines []string) []int {
	result := []int{}
loop:
	for row := 0; row < len(lines); row++ {
		for col := 0; col < len(lines[row]); col++ {
			if lines[row][col] != '.' {
				continue loop
			}
		}
		result = append(result, row)
	}
	return result
}

func getPart2Answer(lines []string) int {
	return findDistances(lines, 1_000_000)
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day11/input.txt")
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
