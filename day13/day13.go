package day13

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func Day13() {
	fmt.Println("--- Day 13: Claw Contraption ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

type Button struct {
	xStep int
	yStep int
}

func getCostLimited(a, b Button, xPrize, yPrize int) int {
	cost := math.MaxInt32
	for i := 0; i < 100; i++ {
		xRemain := xPrize - a.xStep*i
		yRemain := yPrize - a.yStep*i
		if xRemain < 0 || yRemain < 0 {
			continue
		}
		if xRemain%b.xStep != 0 || yRemain%b.yStep != 0 || xRemain/b.xStep != yRemain/b.yStep {
			continue
		}
		if xRemain/b.xStep > 100 {
			continue
		}
		cost = min(cost, 3*i+xRemain/b.xStep)
	}
	if cost == math.MaxInt32 {
		return -1
	}
	return cost
}

func getCost(a, b Button, xPrize, yPrize int) int {
	if a.xStep*b.yStep == a.yStep*b.xStep && a.xStep*yPrize == a.yStep*xPrize && b.xStep*yPrize == b.yStep*xPrize {
		if xPrize%a.xStep != 0 && xPrize%b.xStep != 0 {
			return -1
		}
		if yPrize%a.yStep != 0 && yPrize%b.yStep != 0 {
			return -1
		}
		if xPrize%a.xStep == 0 && xPrize%b.xStep != 0 {
			return 3 * (xPrize / a.xStep)
		}
		if xPrize%a.xStep != 0 && xPrize%b.xStep == 0 {
			return xPrize / b.xStep
		}
		return min(3*(xPrize/a.xStep), xPrize/b.xStep)
	}
	if a.xStep*b.yStep == a.yStep*b.xStep {
		return -1
	}
	xD := a.xStep*b.yStep - a.yStep*b.xStep
	xN := b.yStep*xPrize - b.xStep*yPrize
	if xN%xD != 0 {
		return -1
	}
	yD := -xD
	yN := xPrize*a.yStep - yPrize*a.xStep
	if yN%yD != 0 {
		return -1
	}
	return 3*(xN/xD) + (yN / yD)

}

func getButtonsAndPrize(lines []string, i int) (Button, Button, int, int) {
	buttonMatch := regexp.MustCompile(`Button .\: X\+(\d+), Y\+(\d+)`)
	prizeMatch := regexp.MustCompile(`Prize\: X=(\d+), Y=(\d+)`)
	bAmatches := buttonMatch.FindAllStringSubmatch(lines[i], 3)
	xStepA, _ := strconv.Atoi(bAmatches[0][1])
	yStepA, _ := strconv.Atoi(bAmatches[0][2])
	bBmatches := buttonMatch.FindAllStringSubmatch(lines[i+1], 3)
	xStepB, _ := strconv.Atoi(bBmatches[0][1])
	yStepB, _ := strconv.Atoi(bBmatches[0][2])
	prizeMatches := prizeMatch.FindAllStringSubmatch(lines[i+2], 3)
	xPrize, _ := strconv.Atoi(prizeMatches[0][1])
	yPrize, _ := strconv.Atoi(prizeMatches[0][2])
	return Button{xStepA, yStepA}, Button{xStepB, yStepB}, xPrize, yPrize
}

func getPart1Answer(lines []string) int {
	cost := 0
	for i := 0; i < len(lines); {
		buttonA, buttonB, xPrize, yPrize := getButtonsAndPrize(lines, i)
		if thisCost := getCostLimited(buttonA, buttonB, xPrize, yPrize); thisCost != -1 {
			cost += thisCost
		}
		i += 4
	}
	return cost
}

func getPart2Answer(lines []string) int {
	cost := 0
	for i := 0; i < len(lines); {
		buttonA, buttonB, xPrize, yPrize := getButtonsAndPrize(lines, i)
		if thisCost := getCost(buttonA, buttonB, xPrize+10000000000000, yPrize+10000000000000); thisCost != -1 {
			cost += thisCost
		}
		i += 4
	}
	return cost
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day13/input.txt")
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
