package day5

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func Day5() {
	fmt.Println("--- Day 5: If You Give A Seed A Fertilizer ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart2Answer(lines []string) int {
	seeds := getSeedValues(lines[0])
	maps := getMaps(lines)
	ranges := [][]int{}
	for i := 0; i < len(seeds); i += 2 {
		ranges = append(ranges, []int{seeds[i], seeds[i] + seeds[i+1] - 1})
	}
	sort.SliceStable(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	for i := 0; i < len(maps); i++ {
		transformation := maps[i]
		nextRanges := [][]int{}
		for _, r := range ranges {
			newRanges := [][]int{}
			splits := [][]int{}
			rangeStart := r[0]
			rangeEnd := r[1]
			for _, t := range transformation {
				tStart := t[0]
				tEnd := t[1]
				if overlap(rangeStart, rangeEnd, tStart, tEnd) {
					start := max(rangeStart, tStart)
					end := min(rangeEnd, tEnd)
					splits = append(splits, []int{start, end})
					dStart := getNext(start, transformation)
					dEnd := getNext(end, transformation)
					newRanges = append(newRanges, []int{dStart, dEnd})
				}
			}
			if len(splits) == 0 {
				newRanges = append(newRanges, []int{rangeStart, rangeEnd})
			} else {
				// add all the parts of the range which were not mapped
				if splits[0][0] > rangeStart {
					newRanges = append(newRanges, []int{rangeStart, splits[0][0] - 1})
				}
				for i := 1; i < len(splits); i++ {
					if splits[i][0] > splits[i-1][1]+1 {
						newRanges = append(newRanges, []int{splits[i-1][1] + 1, splits[i][0] - 1})
					}
				}
				if splits[len(splits)-1][1] < rangeEnd {
					newRanges = append(newRanges, []int{splits[len(splits)-1][1] + 1, rangeEnd})
				}
			}
			nextRanges = append(nextRanges, newRanges...)
		}
		ranges = nextRanges
		sort.SliceStable(ranges, func(i, j int) bool {
			return ranges[i][0] < ranges[j][0]
		})
	}
	return ranges[0][0]
}

func overlap(a, b, c, d int) bool {
	return max(b, d)-min(a, c)+1 < b-a+1+d-c+1
}

func getPart1Answer(lines []string) int {
	seeds := getSeedValues(lines[0])
	maps := getMaps(lines)
	result := math.MaxInt32
	for _, seed := range seeds {
		result = min(result, getLocation(seed, maps))
	}
	return result
}

func getMaps(lines []string) [][][]int {
	idx := 0
	maps := [][][]int{}
	mapIndex := 0
	for idx < len(lines) {
		if len(lines[idx]) == 0 || !unicode.IsDigit(rune(lines[idx][0])) {
			idx++
		} else {
			maps = append(maps, [][]int{})
			for idx < len(lines) && len(lines[idx]) > 0 {
				maps[mapIndex] = append(maps[mapIndex], getPath(lines[idx]))
				idx++
			}
			sort.SliceStable(maps[mapIndex], func(i, j int) bool {
				return maps[mapIndex][i][0] < maps[mapIndex][j][0]
			})
			mapIndex++
			idx++
		}
	}
	return maps
}

func getPath(line string) []int {
	result := []int{}
	numbers := strings.Split(line, " ")
	destinationStart, _ := strconv.Atoi(numbers[0])
	sourceStart, _ := strconv.Atoi(numbers[1])
	length, _ := strconv.Atoi(numbers[2])
	result = append(result, sourceStart)
	result = append(result, sourceStart+length-1)
	result = append(result, destinationStart)
	result = append(result, destinationStart+length-1)
	return result
}

func getLocation(seed int, maps [][][]int) int {
	for i := range maps {
		seed = getNext(seed, maps[i])
	}
	return seed
}

func getNext(seed int, thisMap [][]int) int {
	for _, _thisMap := range thisMap {
		if seed >= _thisMap[0] && seed <= _thisMap[1] {
			return seed - _thisMap[0] + _thisMap[2]
		}
	}
	return seed
}

func getSeedValues(line string) []int {
	values := strings.Split(line, ":")[1]
	values = strings.Trim(values, " ")
	result := []int{}
	for _, v := range strings.Split(values, " ") {
		i, _ := strconv.Atoi(v)
		result = append(result, i)
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
