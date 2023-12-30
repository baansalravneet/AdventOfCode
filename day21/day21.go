package day21

import (
	"bufio"
	"fmt"
	"os"
)

func Day21() {
	fmt.Println("--- Day 21: Step Counter ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(lines []string) int {
	x, y := findS(lines)
	return findCover(lines, x, y, 64)
}

/*
assumptions:
 1. the grid is square
 2. the starting point is in the middle of the grid
 3. the rocks(#) are sparse
 4. grid width is odd
 5. there are no rocks between the start and the end of the grid
 6. steps % size == size/2
*/
func getPart2Answer(lines []string) int {
	answer := 0
	x, y := findS(lines)
	steps := 26501365
	size := len(lines)
	reachWidth := (steps / size) - 1

	gridsWithOddStartingSteps := ((reachWidth/2)*2 + 1)
	gridsWithOddStartingSteps = gridsWithOddStartingSteps * gridsWithOddStartingSteps
	gridsWithEvenStartingSteps := (((reachWidth + 1) / 2) * 2)
	gridsWithEvenStartingSteps = gridsWithEvenStartingSteps * gridsWithEvenStartingSteps

	oddPoints := findCover(lines, x, y, size*2+1)
	evenPoints := findCover(lines, x, y, size*2)

	answer += oddPoints * gridsWithOddStartingSteps
	answer += evenPoints * gridsWithEvenStartingSteps

	cornerSteps := size - 1
	cornerTop := findCover(lines, size-1, y, cornerSteps)
	cornerRight := findCover(lines, x, 0, cornerSteps)
	cornerBottom := findCover(lines, 0, y, cornerSteps)
	cornerLeft := findCover(lines, x, size-1, cornerSteps)
	answer += cornerTop + cornerRight + cornerBottom + cornerLeft

	smallGridSteps := (size / 2) - 1
	smallTopRight := findCover(lines, size-1, 0, smallGridSteps)
	smallTopLeft := findCover(lines, size-1, size-1, smallGridSteps)
	smallBottomRight := findCover(lines, 0, 0, smallGridSteps)
	smallBottomLeft := findCover(lines, 0, size-1, smallGridSteps)
	answer += (reachWidth + 1) * (smallTopRight + smallTopLeft + smallBottomRight + smallBottomLeft)

	largeGridSteps := (size / 2) - 1 + size
	largeTopRight := findCover(lines, size-1, 0, largeGridSteps)
	largeTopLeft := findCover(lines, size-1, size-1, largeGridSteps)
	largeBottomRight := findCover(lines, 0, 0, largeGridSteps)
	largeBottomLeft := findCover(lines, 0, size-1, largeGridSteps)
	answer += reachWidth * (largeTopRight + largeTopLeft + largeBottomRight + largeBottomLeft)

	return answer
}

func findS(lines []string) (int, int) {
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == 'S' {
				return i, j
			}
		}
	}
	return -1, -1
}

func findCover(lines []string, x, y, steps int) int {
	n, m := len(lines), len(lines[0])
	keyFormat := "%d,%d"
	visited := map[string]bool{
		fmt.Sprintf(keyFormat, x, y): true,
	}
	answer := make(map[string]bool)
	bfsQueue := [][]int{{x, y, steps}}
	for len(bfsQueue) > 0 {
		current := bfsQueue[0]
		bfsQueue = bfsQueue[1:]
		key := fmt.Sprintf(keyFormat, current[0], current[1])
		if current[2]%2 == 0 {
			answer[key] = true
		}
		if current[2] == 0 {
			continue
		}
		visited[key] = true
		for _, coords := range [][]int{
			{current[0] + 1, current[1]},
			{current[0] - 1, current[1]},
			{current[0], current[1] + 1},
			{current[0], current[1] - 1},
		} {
			nx, ny := coords[0], coords[1]
			if nx < 0 || nx >= n || ny < 0 || ny >= m {
				continue
			}
			if lines[nx][ny] == '#' {
				continue
			}
			key := fmt.Sprintf(keyFormat, nx, ny)
			if _, ok := visited[key]; ok {
				continue
			}
			visited[key] = true
			bfsQueue = append(bfsQueue, []int{nx, ny, current[2] - 1})
		}
	}
	return len(answer)
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day21/input.txt")
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
