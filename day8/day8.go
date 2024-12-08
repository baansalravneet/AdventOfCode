package day8

import (
	"bufio"
	"fmt"
	"os"
)

func Day8() {
	fmt.Println("--- Day 8: Resonant Collinearity ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(lines []string) int {
	freq := getFreq(lines)
	locations := make(map[[2]int]bool)
	for _, v := range freq {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				x1, y1, x2, y2 := v[i][0], v[i][1], v[j][0], v[j][1]
				x3, y3 := x1+(x1-x2), y1+(y1-y2)
				x4, y4 := x2-(x1-x2), y2-(y1-y2)
				if x3 >= 0 && x3 < len(lines) && y3 >= 0 && y3 < len(lines[0]) {
					locations[[2]int{x3, y3}] = true
				}
				if x4 >= 0 && x4 < len(lines) && y4 >= 0 && y4 < len(lines[0]) {
					locations[[2]int{x4, y4}] = true
				}
			}
		}
	}
	return len(locations)
}

func getPart2Answer(lines []string) int {
	freq := getFreq(lines)
	locations := make(map[[2]int]bool)
	for _, v := range freq {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				x1, y1, x2, y2 := v[i][0], v[i][1], v[j][0], v[j][1]
				for mul := -200; mul <= 200; mul++ {
					x3, y3 := x1+mul*(x1-x2), y1+mul*(y1-y2)
					if x3 >= 0 && x3 < len(lines) && y3 >= 0 && y3 < len(lines[0]) {
						locations[[2]int{x3, y3}] = true
					}
				}
			}
		}
	}
	return len(locations)
}

func getFreq(lines []string) map[byte][][2]int {
	freq := make(map[byte][][2]int)
	for i := range lines {
		for j := range lines[i] {
			c := lines[i][j]
			if c == '.' {
				continue
			}
			if _, ok := freq[c]; !ok {
				freq[c] = [][2]int{}
			}
			freq[c] = append(freq[c], [2]int{i, j})
		}
	}
	return freq
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day8/input.txt")
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
