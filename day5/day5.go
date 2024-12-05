package day5

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day5() {
	fmt.Println("--- Day 5: Print Queue ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(lines []string) int {
	before := make(map[int]map[int]bool)
	after := make(map[int]map[int]bool)
	i := 0
	for ; i < len(lines); i++ {
		if lines[i] == "" {
			i++
			break
		}
		a, _ := strconv.Atoi(strings.Split(lines[i], "|")[0])
		b, _ := strconv.Atoi(strings.Split(lines[i], "|")[1])
		if _, ok := before[a]; !ok {
			before[a] = make(map[int]bool)
		}
		before[a][b] = true
		if _, ok := after[b]; !ok {
			after[b] = make(map[int]bool)
		}
		after[b][a] = true
	}
	result := 0
	for ; i < len(lines); i++ {
		nums := strings.Split(lines[i], ",")
		ints := make([]int, len(nums))
		for j := 0; j < len(nums); j++ {
			ints[j], _ = strconv.Atoi(nums[j])
		}
		correct := true
		for j := 0; j < len(ints); j++ {
			for k := j - 1; k >= 0; k-- {
				if _, ok := after[ints[k]]; ok {
					if after[ints[k]][ints[j]] {
						correct = false
					}
				}
				if _, ok := before[ints[j]]; ok {
					if before[ints[j]][ints[k]] {
						correct = false
					}
				}
			}
			for k := j + 1; k < len(ints); k++ {
				if _, ok := before[ints[k]]; ok {
					if before[ints[k]][ints[j]] {
						correct = false
					}
				}
				if _, ok := after[ints[j]]; ok {
					if after[ints[j]][ints[k]] {
						correct = false
					}
				}
			}
		}
		if correct {
			result += ints[len(ints)/2]
		}
	}
	return result
}

func getPart2Answer(lines []string) int {
	before := make(map[int]map[int]bool)
	after := make(map[int]map[int]bool)
	i := 0
	orderUpdate := func(ints []int) int {
		sort.Slice(ints, func(i, j int) bool {
			a, b := ints[i], ints[j]
			if _, ok := before[a]; ok {
				if before[a][b] {
					return false
				}
			}
			if _, ok := after[a]; ok {
				if after[a][b] {
					return true
				}
			}
			if _, ok := before[b]; ok {
				if before[b][a] {
					return true
				}
			}
			if _, ok := after[b]; ok {
				if after[b][a] {
					return false
				}
			}
			return a < b
		})
		return ints[len(ints)/2]
	}
	for ; i < len(lines); i++ {
		if lines[i] == "" {
			i++
			break
		}
		a, _ := strconv.Atoi(strings.Split(lines[i], "|")[0])
		b, _ := strconv.Atoi(strings.Split(lines[i], "|")[1])
		if _, ok := before[a]; !ok {
			before[a] = make(map[int]bool)
		}
		before[a][b] = true
		if _, ok := after[b]; !ok {
			after[b] = make(map[int]bool)
		}
		after[b][a] = true
	}
	result := 0
	for ; i < len(lines); i++ {
		nums := strings.Split(lines[i], ",")
		ints := make([]int, len(nums))
		for j := 0; j < len(nums); j++ {
			ints[j], _ = strconv.Atoi(nums[j])
		}
		correct := true
		for j := 0; j < len(ints); j++ {
			for k := j - 1; k >= 0; k-- {
				if _, ok := after[ints[k]]; ok {
					if after[ints[k]][ints[j]] {
						correct = false
					}
				}
				if _, ok := before[ints[j]]; ok {
					if before[ints[j]][ints[k]] {
						correct = false
					}
				}
			}
			for k := j + 1; k < len(ints); k++ {
				if _, ok := before[ints[k]]; ok {
					if before[ints[k]][ints[j]] {
						correct = false
					}
				}
				if _, ok := after[ints[j]]; ok {
					if after[ints[j]][ints[k]] {
						correct = false
					}
				}
			}
		}
		if !correct {
			result += orderUpdate(ints)
		}
	}
	return result
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day5/input.txt")
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
