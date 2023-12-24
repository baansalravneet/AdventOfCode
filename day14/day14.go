package day14

import (
	"bufio"
	"fmt"
	"os"
)

func Day14() {
	fmt.Println("--- Day 14: Parabolic Reflector Dish ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(lines []string) int {
	rockPositions := getPositions(lines, 'O')
	cubePositions := getPositions(lines, '#')
	rockPositions, cubePositions = rotateAnticlock(rockPositions), rotateAnticlock(cubePositions)
	rockPositions = moveLeft(rockPositions, cubePositions)
	rockPositions, cubePositions = rotateClock(rockPositions), rotateClock(cubePositions)
	return calculateLoad(rockPositions)
}

func getState(positions [][]int) string {
	return fmt.Sprint(positions)
}

func getPart2Answer(lines []string) int {
	rockPositions := getPositions(lines, 'O') // position of a rock in each row
	cubePositions := getPositions(lines, '#') // position of a cube in each row
	visited := make(map[string]int)           // state and the index
	visited[getState(rockPositions)] = 0
	idx := 1
	cycleStart := 0
	for idx < 1000 {
		rockPositions = applyCycle(rockPositions, cubePositions)
		state := getState(rockPositions)
		if v, ok := visited[state]; ok {
			cycleStart = v
			break
		}
		visited[state] = idx
		idx++
	}
	cycleSize := idx - cycleStart
	for i := 0; i < (1000000000-cycleStart)%cycleSize; i++ {
		rockPositions = applyCycle(rockPositions, cubePositions)
	}
	return calculateLoad(rockPositions)
}

func applyCycle(rockPositions, cubePositions [][]int) [][]int {
	rockPositions, cubePositions = rotateAnticlock(rockPositions), rotateAnticlock(cubePositions)
	for i := 0; i < 4; i++ {
		rockPositions = moveLeft(rockPositions, cubePositions)
		rockPositions, cubePositions = rotateClock(rockPositions), rotateClock(cubePositions)
	}
	rockPositions, cubePositions = rotateClock(rockPositions), rotateClock(cubePositions)
	return rockPositions
}

func moveLeft(rockPositions, cubePositions [][]int) [][]int {
	result := [][]int{}
	for i := 0; i < len(rockPositions); i++ {
		current := []int{}
		rockIndex := 0
		cubeIndex := 0
		newPosition := 0
		for rockIndex < len(rockPositions[i]) && cubeIndex < len(cubePositions[i]) {
			if rockPositions[i][rockIndex] < cubePositions[i][cubeIndex] {
				current = append(current, newPosition)
				newPosition++
				rockIndex++
			} else {
				newPosition = cubePositions[i][cubeIndex] + 1
				cubeIndex++
			}
			for cubeIndex < len(cubePositions[i]) && newPosition == cubePositions[i][cubeIndex] {
				newPosition++
				cubeIndex++
			}
		}
		for rockIndex < len(rockPositions[i]) {
			current = append(current, newPosition)
			newPosition++
			rockIndex++
		}
		result = append(result, current)
	}
	return result
}

func transpose(mat [][]int) [][]int {
	n := len(mat)
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < len(mat[i]); j++ {
			result[mat[i][j]] = append(result[mat[i][j]], i)
		}
	}
	return result
}

func mirror(mat [][]int) [][]int {
	result := [][]int{}
	n := len(mat)
	for i := 0; i < n; i++ {
		current := []int{}
		for j := 0; j < len(mat[i]); j++ {
			current = append(current, n-1-mat[i][j])
		}
		result = append(result, current)
	}
	return result
}

func rotateClock(positions [][]int) [][]int {
	transposed := transpose(positions)
	mirrored := mirror(transposed)
	for _, v := range mirrored {
		for i := 0; i < len(v)/2; i++ {
			v[i], v[len(v)-1-i] = v[len(v)-1-i], v[i]
		}
	}
	return mirrored
}

func rotateAnticlock(positions [][]int) [][]int {
	mirrored := mirror(positions)
	return transpose(mirrored)
}

func calculateLoad(rockPositions [][]int) int {
	answer := 0
	for i := range rockPositions {
		answer += (len(rockPositions) - i) * len(rockPositions[i])
	}
	return answer
}

func getPositions(lines []string, char byte) [][]int {
	result := [][]int{}
	for i := 0; i < len(lines); i++ {
		current := []int{}
		for j := 0; j < len(lines[0]); j++ {
			if lines[i][j] == char {
				current = append(current, j)
			}
		}
		result = append(result, current)
	}
	return result
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
