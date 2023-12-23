package day12

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day12() {
	fmt.Println("--- Day 12: Hot Springs ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(lines []string) int {
	result := 0
	for _, line := range lines {
		pattern := strings.Split(line, " ")[0]
		groups := []int{}
		for _, n := range strings.Split(strings.Split(line, " ")[1], ",") {
			num, _ := strconv.Atoi(n)
			groups = append(groups, num)
		}
		result += findCount(pattern, groups)
	}
	return result
}

func findCount(pattern string, groups []int) int {
	// we had groups to match but no characters left
	if len(pattern) == 0 && len(groups) != 0 {
		return 0
	}
	if len(groups) == 0 {
		// we don't have groups to match but there are still # left in the pattern
		if strings.Contains(pattern, "#") {
			return 0
		}
		// we don't have groups to match and there are no # left
		return 1
	}

	for i := 0; i < len(pattern); i++ {
		if pattern[i] == '.' {
			// we do nothing since it does not add to the group
			continue
		}
		answer := 0
		if pattern[i] == '?' {
			answer += findCount("."+pattern[i+1:], groups)
		}
		if pattern[i] == '#' || pattern[i] == '?' {
			count := 0
			toMake := groups[0]
			// we will keep counting until we either
			// 1. reach the end of the pattern with # or ?
			// 2. reach the count we want to make
			idx := i
			for idx < len(pattern) && (pattern[idx] == '?' || pattern[idx] == '#') && count < toMake {
				idx++
				count++
			}
			if count != toMake {
				return answer
			}
			// if we have reached the end of the pattern
			if idx == len(pattern) {
				if len(groups) > 1 {
					return answer
				}
				return answer + 1
			}
			if pattern[idx] == '#' {
				return answer
			}
			// if the next is a '?', make it a dot and find the count for the rest
			if pattern[idx] == '?' {
				return answer + findCount(pattern[idx+1:], groups[1:])
			}
			// if the next is a '.', find the coutn for the rest
			if pattern[idx] == '.' {
				return answer + findCount(pattern[idx:], groups[1:])
			}
		}
		return answer
	}
	return 0
}

func findCountMemoised(pattern string, groups []int, cache map[string]int) int {
	// we had groups to match but no characters left
	if len(pattern) == 0 && len(groups) != 0 {
		return 0
	}
	if len(groups) == 0 {
		// we don't have groups to match but there are still # left in the pattern
		if strings.Contains(pattern, "#") {
			return 0
		}
		// we don't have groups to match and there are no # left
		return 1
	}

	key := fmt.Sprintf("%s:%d", pattern, len(groups))
	if v, ok := cache[key]; ok {
		return v
	}

	for i := 0; i < len(pattern); i++ {
		if pattern[i] == '.' {
			// we do nothing since it does not add to the group
			continue
		}
		answer := 0
		if pattern[i] == '?' {
			answer += findCountMemoised("."+pattern[i+1:], groups, cache)
		}
		if pattern[i] == '#' || pattern[i] == '?' {
			count := 0
			toMake := groups[0]
			// we will keep counting until we either
			// 1. reach the end of the pattern with # or ?
			// 2. reach the count we want to make
			idx := i
			for idx < len(pattern) && (pattern[idx] == '?' || pattern[idx] == '#') && count < toMake {
				idx++
				count++
			}
			if count != toMake {
				cache[key] = answer
				return cache[key]
			}
			// if we have reached the end of the pattern
			if idx == len(pattern) {
				if len(groups) > 1 {
					cache[key] = answer
					return cache[key]
				}
				cache[key] = answer + 1
				return cache[key]
			}
			if pattern[idx] == '#' {
				cache[key] = answer
				return cache[key]
			}
			// if the next is a '?', make it a dot and find the count for the rest
			if pattern[idx] == '?' {
				cache[key] = answer + findCountMemoised(pattern[idx+1:], groups[1:], cache)
				return cache[key]
			}
			// if the next is a '.', find the coutn for the rest
			if pattern[idx] == '.' {
				cache[key] = answer + findCountMemoised(pattern[idx:], groups[1:], cache)
				return cache[key]
			}
		}
		cache[key] = answer
		return cache[key]
	}
	cache[key] = 0
	return 0
}

func getPart2Answer(lines []string) int {
	result := 0
	for _, line := range lines {
		pattern := strings.Split(line, " ")[0]
		pattern = strings.Join([]string{pattern, pattern, pattern, pattern, pattern}, "?")
		a := strings.Split(line, " ")[1]
		a = strings.Join([]string{a, a, a, a, a}, ",")
		groups := []int{}
		for _, n := range strings.Split(a, ",") {
			num, _ := strconv.Atoi(n)
			groups = append(groups, num)
		}
		cache := make(map[string]int)
		result += findCountMemoised(pattern, groups, cache)
	}
	return result
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day12/input.txt")
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
