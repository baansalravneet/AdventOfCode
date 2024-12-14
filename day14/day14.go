package day14

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Day14() {
	fmt.Println("--- Day 14: Restroom Redoubt ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(lines []string) int {
	matcher := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
	boundX, boundY := 101, 103
	seconds := 100
	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, line := range lines {
		match := matcher.FindAllStringSubmatch(line, 5)
		x, _ := strconv.Atoi(match[0][1])
		y, _ := strconv.Atoi(match[0][2])
		vx, _ := strconv.Atoi(match[0][3])
		vy, _ := strconv.Atoi(match[0][4])
		finalX := (((x + vx*seconds) % boundX) + boundX) % boundX
		finalY := (((y + vy*seconds) % boundY) + boundY) % boundY
		if finalX < boundX/2 && finalY < boundY/2 {
			q1++
		} else if finalX > boundX/2 && finalY < boundY/2 {
			q2++
		} else if finalX < boundX/2 && finalY > boundY/2 {
			q3++
		} else if finalX > boundX/2 && finalY > boundY/2 {
			q4++
		}
	}
	return q1 * q2 * q3 * q4
}

func getPart2Answer(lines []string) int {
	matcher := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
	boundX, boundY := 101, 103
	positions, velocities := make([][2]int, len(lines)*2), make([][2]int, len(lines)*2)
	for _, line := range lines {
		match := matcher.FindAllStringSubmatch(line, 5)
		x, _ := strconv.Atoi(match[0][1])
		y, _ := strconv.Atoi(match[0][2])
		vx, _ := strconv.Atoi(match[0][3])
		vy, _ := strconv.Atoi(match[0][4])
		positions = append(positions, [2]int{x, y})
		velocities = append(velocities, [2]int{vx, vy})
	}
	grid := make([][]byte, boundY)
	for i := range grid {
		grid[i] = make([]byte, boundX)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}
	for it := 0; it < 101*103; it++ {
		for i := 0; i < boundY; i++ {
			for j := 0; j < boundX; j++ {
				grid[i][j] = ' '
			}
		}
		for i := 0; i < len(positions); i++ {
			x := ((positions[i][0]+velocities[i][0]*it)%boundX + boundX) % boundX
			y := ((positions[i][1]+velocities[i][1]*it)%boundY + boundY) % boundY
			grid[y][x] = '#'
		}
		for i := 0; i < boundY; i++ {
			// fmt.Println(string(grid[i]))
		}
		// fmt.Println(it, "========================")
	}
	return 0
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day14/input.txt")
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
