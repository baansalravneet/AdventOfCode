package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2() {
	fmt.Println("--- Day 2: Red-Nosed Reports ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(lines []string) int {
	count := 0
	for _, line := range lines {
		report := getReport(line)
		if checkNorm(report) {
			count++
		}
	}
	return count
}

func checkNorm(report []int) bool {
	inc, dec := true, true
	n := len(report)
	for i := 1; i < n; i++ {
		if report[i]-report[i-1] > 3 || report[i]-report[i-1] < 1 {
			inc = false
		}
		if report[n-1-i]-report[n-1-i+1] > 3 || report[n-1-i]-report[n-1-i+1] < 1 {
			dec = false
		}
	}
	return inc || dec
}

func checkInc(report []int) bool {
	check := func(ignore int) bool {
		previous, i := report[0], 1
		if ignore == 0 {
			previous = report[1]
			i = 2
		}
		for ; i < len(report); i++ {
			if i == ignore {
				continue
			}
			if report[i]-previous < 1 || report[i]-previous > 3 {
				return false
			}
			previous = report[i]
		}
		return true
	}
	for i := 0; i < len(report); i++ {
		if check(i) {
			return true
		}
	}
	return false
}

func checkDec(report []int) bool {
	n := len(report)
	check := func(ignore int) bool {
		previous, i := report[n-1], n-2
		if ignore == n-1 {
			previous = report[n-2]
			i = n - 3
		}
		for ; i >= 0; i-- {
			if i == ignore {
				continue
			}
			if report[i]-previous < 1 || report[i]-previous > 3 {
				return false
			}
			previous = report[i]
		}
		return true
	}
	for i := 0; i < len(report); i++ {
		if check(i) {
			return true
		}
	}
	return false
}

func getPart2Answer(lines []string) int {
	count := 0
	for _, line := range lines {
		report := getReport(line)
		if checkNorm(report) || checkInc(report) || checkDec(report) {
			count++
		}
	}
	return count
}

func getReport(line string) []int {
	split := strings.Split(line, " ")
	report := []int{}
	for _, s := range split {
		v, _ := strconv.Atoi(s)
		report = append(report, v)
	}
	return report
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day2/input.txt")
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
