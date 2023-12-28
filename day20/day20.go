package day20

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day20() {
	fmt.Println("--- Day 20: Pulse Propagation ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

type pulse struct {
	receiver string
	pt       pulseType
	sender   string
}

type pulseType int

const (
	low  pulseType = 0
	high pulseType = 1
)

type moduleType byte

const (
	ff  moduleType = '%'
	con moduleType = '&'
)

type state int

const (
	off state = 0
	on  state = 1
)

type module struct {
	kind        moduleType
	inputStates map[string]state
	receivers   []string
	name        string
}

func getModuleMap(lines []string) (map[string]module, string) {
	moduleMap := make(map[string]module)
	broadcaster := ""
	for _, s := range lines {
		moduleString := s[0:strings.Index(s, " ")]
		moduleName := moduleString[1:]
		if moduleString[0] == '%' {
			moduleMap[moduleName] = module{
				kind:        ff,
				inputStates: map[string]state{"self": off},
				name:        moduleName,
			}
		} else if moduleString[0] == '&' {
			moduleMap[moduleName] = module{
				kind:        con,
				inputStates: make(map[string]state),
				name:        moduleName,
			}
		} else {
			broadcaster = s
		}
	}
	return moduleMap, broadcaster
}

func getReceivers(l string) []string {
	result := []string{}
	for _, r := range strings.Split(strings.Trim(l[strings.Index(l, ">")+1:], " "), ",") {
		result = append(result, strings.Trim(r, " "))
	}
	return result
}

func addReceivers(mm *map[string]module, lines []string) {
	for _, l := range lines {
		if l[0] == '%' || l[0] == '&' {
			sender := l[1:strings.Index(l, " ")]
			for _, r := range getReceivers(l) {
				m := (*mm)[sender]
				m.receivers = append(m.receivers, r)
				(*mm)[sender] = m
				receiver, ok := (*mm)[r]
				if ok && receiver.kind == con {
					receiver.inputStates[sender] = off
				}
				if ok {
					(*mm)[r] = receiver
				}
			}
		}
	}
}

func (r *module) receive(p pulse) []pulse {
	result := []pulse{}
	if r.kind == ff {
		if p.pt == high {
			return result
		} else {
			if r.inputStates["self"] == on {
				for _, next := range r.receivers {
					result = append(result, pulse{receiver: next, pt: low, sender: r.name})
				}
				r.inputStates["self"] = off
			} else {
				for _, next := range r.receivers {
					result = append(result, pulse{receiver: next, pt: high, sender: r.name})
				}
				r.inputStates["self"] = on
			}
			return result
		}
	} else {
		if p.pt == high {
			r.inputStates[p.sender] = on
			allOn := true
			for _, v := range r.inputStates {
				if v == off {
					allOn = false
					break
				}
			}
			if allOn {
				for _, next := range r.receivers {
					result = append(result, pulse{receiver: next, pt: low, sender: r.name})
				}
			} else {
				for _, next := range r.receivers {
					result = append(result, pulse{receiver: next, pt: high, sender: r.name})
				}
			}
			return result
		} else {
			r.inputStates[p.sender] = off
			for _, next := range r.receivers {
				result = append(result, pulse{receiver: next, pt: high, sender: r.name})
			}
		}
	}
	return result
}

func getPart1Answer(lines []string) int {
	moduleMap, broadcaster := getModuleMap(lines)
	addReceivers(&moduleMap, lines)
	q := []pulse{} // modules receiving the first pulse
	countHigh := 0
	countLow := 0
	for i := 1; i <= 1000; i++ {
		for _, r := range getReceivers(broadcaster) {
			q = append(q, pulse{receiver: r, pt: low, sender: "broadcaster"})
		}
		countLow++
		for len(q) > 0 {
			size := len(q)
			for size > 0 {
				p := q[0]
				if p.pt == high {
					countHigh++
				} else {
					countLow++
				}
				q = q[1:]
				receiver, ok := moduleMap[p.receiver]
				if !ok {
					size--
					continue
				}
				outputs := receiver.receive(p)
				if len(outputs) != 0 {
					q = append(q, outputs...)
				}
				size--
			}
		}
	}
	return countHigh * countLow
}

// we can reverse engineer and see that rx is fed by only kz
// kz is a conjunction module which is fed by bg, sj, qq, and ls
// we can record the cycle lengths for each of these for
// when they emit a high pulse and take LCM
func getPart2Answer(lines []string) int {
	trackModules := []string{"bg", "qq", "sj", "ls"}
	lcm := []int{}
	for _, mod := range trackModules {
		moduleMap, broadcaster := getModuleMap(lines)
		addReceivers(&moduleMap, lines)
		lcm = append(lcm, getCycleLength(moduleMap, broadcaster, mod))
	}
	result := 1
	for _, length := range lcm {
		result = (result * length) / gcd(result, length)
	}
	return result
}

func gcd(a, b int) int {
	if a > b {
		return gcd(b, a)
	}
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func getCycleLength(moduleMap map[string]module, broadcaster, mod string) int {
	q := []pulse{} // modules receiving the first pulse
	for i := 1; ; i++ {
		for _, r := range getReceivers(broadcaster) {
			q = append(q, pulse{receiver: r, pt: low, sender: "broadcaster"})
		}
		for len(q) > 0 {
			size := len(q)
			for size > 0 {
				p := q[0]
				if p.receiver == "kz" && p.pt == high && p.sender == mod {
					return i
				}
				q = q[1:]
				receiver, ok := moduleMap[p.receiver]
				if !ok {
					size--
					continue
				}
				outputs := receiver.receive(p)
				if len(outputs) != 0 {
					q = append(q, outputs...)
				}
				size--
			}
		}
	}
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day20/input.txt")
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
