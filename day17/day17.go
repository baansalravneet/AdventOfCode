package day17

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Day17() {
	fmt.Println("--- Day 17: Chronospatial Computer ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func computer(program []string, regA, regB, regC int) []string {
	ip := 0
	result := []string{}
	for ip < len(program) {
		opCode, _ := strconv.Atoi(program[ip])
		operand, _ := strconv.Atoi(program[ip+1])
		comboOperand := operand
		if operand == 4 {
			comboOperand = regA
		} else if operand == 5 {
			comboOperand = regB
		} else if operand == 6 {
			comboOperand = regC
		}
		switch opCode {
		case 0:
			regA = regA >> comboOperand
		case 1:
			regB ^= operand
		case 2:
			regB = comboOperand & 7
		case 3:
			if regA != 0 {
				ip = operand
				ip -= 2
			}
		case 4:
			regB ^= regC
		case 5:
			result = append(result, strconv.Itoa(comboOperand&7))
		case 6:
			regB = regA >> comboOperand
		case 7:
			regC = regA >> comboOperand
		}
		ip += 2
	}
	return result
}

func getPart1Answer(lines []string) string {
	matcher := regexp.MustCompile(`(\d+)`)
	programMatcher := regexp.MustCompile(`Program\: (.*)`)
	regA, _ := strconv.Atoi(matcher.FindString(lines[0]))
	regB, _ := strconv.Atoi(matcher.FindString(lines[1]))
	regC, _ := strconv.Atoi(matcher.FindString(lines[2]))
	program := strings.Split(programMatcher.FindAllStringSubmatch(lines[4], 2)[0][1], ",")
	result := computer(program, regA, regB, regC)
	return strings.Join(result, ",")
}

func getPart2Answer(lines []string) int {
	programMatcher := regexp.MustCompile(`Program\: (.*)`)
	program := strings.Split(programMatcher.FindAllStringSubmatch(lines[4], 2)[0][1], ",")
	candidates := []int{0}
	for instr := range program {
		nextCandidates := []int{}
		for _, c := range candidates {
			for i := 0; i < 8; i++ {
				next := (c << 3) + i
				if slices.Equal(computer(program, next, 0, 0), program[len(program)-1-instr:]) {
					nextCandidates = append(nextCandidates, next)
				}
			}
		}
		candidates = nextCandidates
	}
	result := 1 << (3 * (len(program) + 1))
	for _, c := range candidates {
		result = min(result, c)
	}
	return result
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day17/input.txt")
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
