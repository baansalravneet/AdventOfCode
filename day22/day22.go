package day22

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day22() {
	fmt.Println("--- Day 22: Sand Slabs ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(lines []string) int {
	bricks := getBricks(lines)
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i][0][2] < bricks[j][0][2]
	})
	groundBricks := drop(bricks)
	brickDependencies := getDependencies(groundBricks)
	inDegrees := getInDegrees(brickDependencies)
	result := 0
loop:
	for _, edge := range brickDependencies {
		if len(edge) == 0 {
			result++
			continue
		}
		for k := range edge {
			if inDegrees[k] <= 1 {
				continue loop
			}
		}
		result++
	}
	return result
}

func getPart2Answer(lines []string) int {
	bricks := getBricks(lines)
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i][0][2] < bricks[j][0][2]
	})
	groundBricks := drop(bricks)
	brickDependencies := getDependencies(groundBricks)
	inDegrees := getInDegrees(brickDependencies)
	result := 0
	for i := range inDegrees {
		copied := make([]int, len(inDegrees))
		copy(copied, inDegrees)
		result += removeThis(i, brickDependencies, copied)
	}
	return result
}

func removeThis(i int, edges map[int]map[int]bool, inDegrees []int) int {
	result := 0
	bfsQ := []int{}
	bfsQ = append(bfsQ, i)
	for len(bfsQ) > 0 {
		current := bfsQ[0]
		bfsQ = bfsQ[1:]
		for k := range edges[current] {
			inDegrees[k]--
			if inDegrees[k] == 0 {
				result++
				bfsQ = append(bfsQ, k)
			}
		}
	}
	return result
}

func getInDegrees(edges map[int]map[int]bool) []int {
	result := make([]int, len(edges))
	for _, v := range edges {
		for k := range v {
			result[k]++
		}
	}
	for i, v := range result {
		if v == 0 {
			result[i] = -1
		}
	}
	return result
}

func getDependencies(bricks [][][3]int) map[int]map[int]bool {
	result := make(map[int]map[int]bool)
	for i := range bricks {
		result[i] = make(map[int]bool)
	}
	for i, brick := range bricks {
		for _, coord := range brick {
			if j, ok := under(coord, bricks); ok && i != j {
				result[j][i] = true
			}
		}
	}
	return result
}

func under(coord [3]int, bricks [][][3]int) (int, bool) {
	check := [3]int{coord[0], coord[1], coord[2] - 1}
	for i, brick := range bricks {
		for _, c := range brick {
			if check == c {
				return i, true
			}
		}
	}
	return 0, false
}

func drop(bricks [][][3]int) [][][3]int {
	droppedBricks := [][][3]int{}
	occupied := make(map[[3]int]bool)
	for _, brick := range bricks {
		height := brick[0][2]
		newHeight := height
	loop:
		for ; newHeight > 0; newHeight-- {
			for _, coords := range brick {
				newCoords := [3]int{coords[0], coords[1], coords[2] + newHeight - height}
				spaceOccupied := false
				if _, ok := occupied[newCoords]; ok {
					spaceOccupied = true
				}
				if spaceOccupied {
					break loop
				}
			}
		}
		newHeight = newHeight + 1
		newBrick := [][3]int{}
		for _, coords := range brick {
			newCoords := [3]int{coords[0], coords[1], coords[2] + newHeight - height}
			occupied[newCoords] = true
			newBrick = append(newBrick, newCoords)
		}
		droppedBricks = append(droppedBricks, newBrick)
	}
	return droppedBricks
}

func getBricks(lines []string) [][][3]int {
	result := [][][3]int{}
	for _, l := range lines {
		startCoord := strings.Split(l, "~")[0]
		endCoord := strings.Split(l, "~")[1]
		xa, _ := strconv.Atoi(strings.Split(startCoord, ",")[0])
		ya, _ := strconv.Atoi(strings.Split(startCoord, ",")[1])
		za, _ := strconv.Atoi(strings.Split(startCoord, ",")[2])
		xb, _ := strconv.Atoi(strings.Split(endCoord, ",")[0])
		yb, _ := strconv.Atoi(strings.Split(endCoord, ",")[1])
		zb, _ := strconv.Atoi(strings.Split(endCoord, ",")[2])
		brick := [][3]int{}
		for x := min(xa, xb); x <= max(xa, xb); x++ {
			for y := min(ya, yb); y <= max(ya, yb); y++ {
				for z := min(za, zb); z <= max(za, zb); z++ {
					brick = append(brick, [3]int{x, y, z})
				}
			}
		}
		result = append(result, brick)
	}
	return result
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day22/input.txt")
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
