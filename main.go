package main

import (
	day1 "adventofcode/day1"
	day2 "adventofcode/day2"
	day3 "adventofcode/day3"
	day4 "adventofcode/day4"
	day5 "adventofcode/day5"
	day6 "adventofcode/day6"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print()
		fmt.Println("Enter day number to see result")
		fmt.Print(">")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		fmt.Println()
		dayNumber, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid Input")
			fmt.Println()
			continue
		}
		switch dayNumber {
		case 1:
			day1.Day1()
		case 2:
			day2.Day2()
		case 3:
			day3.Day3()
		case 4:
			day4.Day4()
		case 5:
			day5.Day5()
		case 6:
			day6.Day6()
		default:
			fmt.Println("Invalid input")
		}
		fmt.Println()
	}
}
