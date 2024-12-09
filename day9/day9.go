package day9

import (
	"bufio"
	"fmt"
	"os"
)

func Day9() {
	fmt.Println("--- Day 9: Disk Fragmenter ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getDisk(lines []string) ([]int, int) {
	disk := []int{}
	line := lines[0]
	file := true
	id := 0
	for _, c := range line {
		val := int(c - '0')
		if file {
			for val > 0 {
				disk = append(disk, id)
				val--
			}
			id++
		} else {
			for val > 0 {
				disk = append(disk, -1)
				val--
			}
		}
		file = !file
	}
	return disk, id
}

func getAnswer(disk []int) int {
	answer := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			continue
		}
		answer += disk[i] * i
	}
	return answer
}

func getPart1Answer(lines []string) int {
	disk, _ := getDisk(lines)
	left, right := 0, len(disk)-1
	for left < right {
		if disk[right] == -1 {
			right--
			continue
		}
		if disk[left] != -1 {
			left++
			continue
		}
		disk[left], disk[right] = disk[right], disk[left]
		left++
		right--
	}
	return getAnswer(disk)
}

func getPart2Answer(lines []string) int {
	disk, id := getDisk(lines)
	id--
	right := len(disk) - 1
	for right >= 0 && id >= 0 {
		if disk[right] != id {
			right--
			continue
		}
		for left := 0; left < right; left++ {
			if disk[left] != -1 {
				continue
			}
			if canFit(disk, left, right) {
				for left < right && disk[right] == id {
					disk[left], disk[right] = disk[right], disk[left]
					left++
					right--
				}
				break
			}
		}
		id--
	}
	return getAnswer(disk)
}

func canFit(disk []int, left, right int) bool {
	emptyCount := 0
	for i := left; i < right && disk[i] == -1; i++ {
		emptyCount++
	}
	fileCount := 0
	for i := right; i > left && disk[i] == disk[right]; i-- {
		fileCount++
	}
	return emptyCount >= fileCount
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day9/input.txt")
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
