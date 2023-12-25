package day15

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Day15() {
	fmt.Println("--- Day 15: Lens Library ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(lines []string) int {
	result := 0
	for _, l := range lines {
		steps := strings.Split(l, ",")
		for _, step := range steps {
			result += getBox(step)
		}
	}
	return result
}

func getLabel(step string) (string, int) {
	if step[len(step)-1] == '-' {
		return step[:len(step)-1], -1
	}
	value := step[len(step)-1] - '0'
	return step[:len(step)-2], int(value)
}

func getBox(label string) int {
	result := 0
	for _, c := range label {
		result += int(c)
		result *= 17
		result %= 256
	}
	return result
}

func getPart2Answer(lines []string) int {
	hashMap := make(map[string]int)
	labelOrder := make(map[string]int)
	stepCounter := 0
	for _, l := range lines {
		steps := strings.Split(l, ",")
		for _, step := range steps {
			label, v := getLabel(step)
			if v == -1 {
				delete(hashMap, label)
				delete(labelOrder, label)
			} else {
				hashMap[label] = v
				if _, ok := labelOrder[label]; !ok {
					labelOrder[label] = stepCounter
					stepCounter++
				}
			}
		}
	}
	boxes := make([][]string, 256)
	for k := range hashMap {
		box := getBox(k)
		boxes[box] = append(boxes[box], k)
	}
	for _, box := range boxes {
		sort.Slice(box, func(i, j int) bool {
			return labelOrder[box[i]] < labelOrder[box[j]]
		})
	}
	focusingPower := 0
	for boxNumber, box := range boxes {
		if len(box) == 0 {
			continue
		}
		for slotNumber, label := range box {
			focusingPower += (boxNumber + 1) * (slotNumber + 1) * (hashMap[label])
		}
	}
	return focusingPower
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
