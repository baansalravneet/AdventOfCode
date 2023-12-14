package day4

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "regexp"
    "strings"
)

func Day4() {
    fmt.Println("--- Day 4: Scratchcards ---")
    lines := getInput()
    part1Answer := getPart1Answer(lines)
    fmt.Println("Part1:", part1Answer)
    part2Answer := getPart2Answer(lines)
    fmt.Println("Part2:", part2Answer)
}

func getPart2Answer(lines []string) int {
    winnings := []int{}
    for _, line := range lines {
        cardInfo := strings.Split(line, ":")
        numbers := strings.Split(cardInfo[1], "|")
        re := regexp.MustCompile("\\s+")
        winningNumbers := re.Split(strings.Trim(numbers[0], " "), -1)
        cardNumbers := re.Split(strings.Trim(numbers[1], " "), -1)
        winnings = append(winnings, getMatchingCount(winningNumbers, cardNumbers))
    }
    cards := make([]int, len(winnings))
    for i := range cards {
        cards[i] = 1
    }
    result := 0
    for idx, count := range winnings {
        current := cards[idx]
        for i := idx + 1; i <= idx+count && i < len(cards); i++ {
            cards[i] += current
        }
        result += current
    }
    return result
}

func getPart1Answer(lines []string) int {
    answer := 0
    for _, line := range lines {
        cardInfo := strings.Split(line, ":")
        numbers := strings.Split(cardInfo[1], "|")
        re := regexp.MustCompile("\\s+")
        winningNumbers := re.Split(strings.Trim(numbers[0], " "), -1)
        cardNumbers := re.Split(strings.Trim(numbers[1], " "), -1)
        answer += getMatchValue(winningNumbers, cardNumbers)
    }
    return answer
}

func getMatchValue(winningNumbers, cardNumbers []string) int {
    matching := getMatchingCount(winningNumbers, cardNumbers)
    if matching == 0 {
        return 0
    }
    return int(math.Pow(2, float64(matching-1)))
}

func getMatchingCount(winningNumbers, cardNumbers []string) int {
    count := 0
    for _, number := range cardNumbers {
        if contains(winningNumbers, number) {
            count++
        }
    }
    return count
}

func contains(arr []string, target string) bool {
    for _, s := range arr {
        if s == target {
            return true
        }
    }
    return false
}

func getInput() []string {
    lines := []string{}
    file, err := os.Open("day4/input.txt")
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
