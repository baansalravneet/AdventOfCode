package day18

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day18() {
	fmt.Println("--- Day 18: Lavaduct Lagoon ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(lines []string) int {
	endPoints, length := getEndPointsPart1(lines)
	area := getArea(endPoints)
	return actualArea(area, length)
}

func getPart2Answer(lines []string) int {
	endPoints, length := getEndPointsPart2(lines)
	area := getArea(endPoints)
	return -area - (length / 2) + 1 + length
}

func getEndPointsPart1(lines []string) ([][]int, int) {
	length := 0
	x, y := 0, 0
	endPoints := [][]int{}
	for _, l := range lines {
		direction := strings.Split(l, " ")[0]
		distance, _ := strconv.Atoi(strings.Split(l, " ")[1])
		endPoints = append(endPoints, []int{x, y})
		length += distance
		switch direction {
		case "R":
			y += distance
		case "L":
			y -= distance
		case "U":
			x -= distance
		case "D":
			x += distance
		}
	}
	return endPoints, length
}

func getEndPointsPart2(lines []string) ([][]int, int) {
	endPoints := [][]int{}
	length := 0
	x, y := 0, 0
	for _, l := range lines {
		move := strings.Split(l, " ")[2]
		d, _ := strconv.ParseInt(move[2:7], 16, 32)
		distance := int(d)
		direction := move[7:8]
		endPoints = append(endPoints, []int{x, y})
		length += distance
		switch direction {
		case "0":
			y += distance
		case "1":
			x += distance
		case "2":
			y -= distance
		case "3":
			x -= distance
		}
	}
	return endPoints, length
}

/*
lets say the boundary is like this

###
# #
###

the area that we get from the shoelace algo is for the boundary

F-7
| |
L-J

so we remove half the length of the boundary+1 to get the internal area
and then add boundary to get the total area of the dig
*/
func actualArea(area, length int) int {
	if area < 0 {
		return actualArea(-area, length)
	}
	return area - (length / 2) + 1 + length
}

func getArea(endPoints [][]int) int { // shoelace algo
	endPoints = append(endPoints, endPoints[0])
	area := 0
	for i := 0; i < len(endPoints)-1; i++ {
		area += endPoints[i][0] * endPoints[i+1][1]
		area -= endPoints[i][1] * endPoints[i+1][0]
	}
	return area / 2
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day18/input.txt")
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
