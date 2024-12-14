package day14

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day14() {
	fmt.Println("--- Day 14: Restroom Redoubt ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

const boundX, boundY = 101, 103

func getDangerScore(xPos, yPos, xVel, yVel []int, seconds int) int {
	q1, q2, q3, q4 := 0, 0, 0, 0
	for i := 0; i < len(xPos); i++ {
		x, y, vx, vy := xPos[i], yPos[i], xVel[i], yVel[i]
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

func getPositionsAndVelocities(lines []string) ([]int, []int, []int, []int) {
	matcher := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
	xPos, yPos, xVel, yVel := []int{}, []int{}, []int{}, []int{}
	for _, line := range lines {
		match := matcher.FindAllStringSubmatch(line, 5)
		x, _ := strconv.Atoi(match[0][1])
		y, _ := strconv.Atoi(match[0][2])
		vx, _ := strconv.Atoi(match[0][3])
		vy, _ := strconv.Atoi(match[0][4])
		xPos, yPos = append(xPos, x), append(yPos, y)
		xVel, yVel = append(xVel, vx), append(yVel, vy)
	}
	return xPos, yPos, xVel, yVel
}

func getPart1Answer(lines []string) int {
	xPos, yPos, xVel, yVel := getPositionsAndVelocities(lines)
	return getDangerScore(xPos, yPos, xVel, yVel, 100)
}

func getPart2Answer(lines []string) int {
	xPos, yPos, xVel, yVel := getPositionsAndVelocities(lines)
	grid := make([][]byte, boundX)
	for i := 0; i < boundX; i++ {
		grid[i] = make([]byte, boundY)
	}
	for seconds := 0; seconds < boundX*boundY; seconds++ {
		for i := 0; i < len(xPos); i++ {
			x := (((xPos[i] + xVel[i]*seconds) % boundX) + boundX) % boundX
			y := (((yPos[i] + yVel[i]*seconds) % boundY) + boundY) % boundY
			grid[x][y] = '#'
		}
		for i := range grid {
			if strings.Contains(string(grid[i]), "###########") {
				return seconds
			}
			for j := range grid[i] {
				grid[i][j] = ' '
			}
		}
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
