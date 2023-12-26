package day19

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day19() {
	fmt.Println("--- Day 19: Aplenty ---")
	lines := getInput()
	workflows := []string{}
	parts := []string{}
	i := 0
	for lines[i] != "" {
		workflows = append(workflows, lines[i])
		i++
	}
	i++
	for i < len(lines) {
		parts = append(parts, lines[i])
		i++
	}
	part1Answer := getPart1Answer(workflows, parts)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(workflows)
	fmt.Println("Part2:", part2Answer)
}

type partsByRange struct {
	xvalues         []int
	mvalues         []int
	avalues         []int
	svalues         []int
	currentWorkflow string
	empty           bool
}

func (p partsByRange) getTotal() int {
	result := p.xvalues[1] - p.xvalues[0] + 1
	result *= p.mvalues[1] - p.mvalues[0] + 1
	result *= p.avalues[1] - p.avalues[0] + 1
	result *= p.svalues[1] - p.svalues[0] + 1
	return result
}

type part struct {
	values map[byte]int
}

func (p part) getValue(b byte) int {
	return p.values[b]
}

func (p partsByRange) getValueRange(b byte) []int {
	if b == 'x' {
		return p.xvalues
	}
	if b == 'm' {
		return p.mvalues
	}
	if b == 'a' {
		return p.avalues
	}
	return p.svalues
}

type rule struct {
	property    byte
	value       int
	comp        bool
	destination string
}

func (r rule) evaluateByRange(p partsByRange) (partsByRange, partsByRange) {
	if r.property == 0 {
		p.currentWorkflow = r.destination
		return p, partsByRange{empty: true}
	}
	partRange := p.getValueRange(r.property)
	if r.comp {
		if partRange[1] <= r.value { // everything got rejected
			return partsByRange{empty: true}, p
		}
		if partRange[0] > r.value { // everything got accepted
			return p, partsByRange{empty: true}
		}
		acceptedRange := []int{r.value + 1, partRange[1]}
		rejectedRange := []int{partRange[0], r.value}
		acceptedParts := partsByRange{
			xvalues:         p.xvalues,
			mvalues:         p.mvalues,
			avalues:         p.avalues,
			svalues:         p.svalues,
			currentWorkflow: r.destination,
		}
		rejectedParts := partsByRange{
			xvalues:         p.xvalues,
			mvalues:         p.mvalues,
			avalues:         p.avalues,
			svalues:         p.svalues,
			currentWorkflow: p.currentWorkflow,
		}
		acceptedParts.setRange(acceptedRange, r.property)
		rejectedParts.setRange(rejectedRange, r.property)
		return acceptedParts, rejectedParts
	} else {
		if partRange[0] >= r.value { // everything got rejected
			return partsByRange{empty: true}, p
		}
		if partRange[1] < r.value { // everything got accepted
			return p, partsByRange{empty: true}
		}
		acceptedRange := []int{partRange[0], r.value - 1}
		rejectedRange := []int{r.value, partRange[1]}
		acceptedParts := partsByRange{
			xvalues:         p.xvalues,
			mvalues:         p.mvalues,
			avalues:         p.avalues,
			svalues:         p.svalues,
			currentWorkflow: r.destination,
		}
		rejectedParts := partsByRange{
			xvalues:         p.xvalues,
			mvalues:         p.mvalues,
			avalues:         p.avalues,
			svalues:         p.svalues,
			currentWorkflow: p.currentWorkflow,
		}
		acceptedParts.setRange(acceptedRange, r.property)
		rejectedParts.setRange(rejectedRange, r.property)
		return acceptedParts, rejectedParts
	}
}

func (p *partsByRange) setRange(newRange []int, property byte) {
	switch property {
	case 'x':
		(*p).xvalues = newRange
	case 'm':
		(*p).mvalues = newRange
	case 'a':
		(*p).avalues = newRange
	case 's':
		(*p).svalues = newRange
	}
}

func (r rule) evaluate(p part) string {
	if r.property == 0 {
		return r.destination
	}
	partValue := p.getValue(r.property)
	if r.comp && partValue > r.value {
		return r.destination
	} else if !r.comp && partValue < r.value {
		return r.destination
	}
	return ""
}

type workflow struct {
	rules []rule
}

func parseRule(s string) rule {
	idx := strings.Index(s, ":")
	if idx == -1 {
		return rule{destination: s}
	}
	value, _ := strconv.Atoi(s[2:idx])
	result := rule{
		destination: s[idx+1:],
		property:    s[0],
		value:       value,
	}
	if s[1] == '<' {
		result.comp = false
	} else {
		result.comp = true
	}
	return result
}

func parseWorkflow(s string) (string, workflow) {
	bs := strings.Index(s, "{")
	be := strings.Index(s, "}")
	name := s[:bs]
	ruleStrings := strings.Split(s[bs+1:be], ",")
	rules := []rule{}
	for _, r := range ruleStrings {
		rules = append(rules, parseRule(r))
	}
	return name, workflow{rules: rules}
}

func (p part) getTotal() int {
	result := 0
	for _, v := range p.values {
		result += v
	}
	return result
}

func parsePart(s string) part {
	s = s[1 : len(s)-1]
	propertyMap := make(map[byte]int)
	for _, sp := range strings.Split(s, ",") {
		property := sp[0]
		value, _ := strconv.Atoi(sp[2:])
		propertyMap[property] = value
	}
	return part{values: propertyMap}
}

func getPart1Answer(workflows, partStrings []string) int {
	workflowMap := make(map[string]workflow)
	for _, w := range workflows {
		name, workflow := parseWorkflow(w)
		workflowMap[name] = workflow
	}
	parts := []part{}
	for _, partString := range partStrings {
		parts = append(parts, parsePart(partString))
	}
	result := 0
	for _, part := range parts {
		current := "in"
		for current != "R" && current != "A" {
			workflow := workflowMap[current]
			for _, r := range workflow.rules {
				ruleResult := r.evaluate(part)
				if ruleResult == "" {
					continue
				}
				current = ruleResult
				break
			}
		}
		if current == "A" {
			result += part.getTotal()
		}
	}
	return result
}

func getPart2Answer(workflows []string) int {
	workflowMap := make(map[string]workflow)
	for _, w := range workflows {
		name, workflow := parseWorkflow(w)
		workflowMap[name] = workflow
	}
	part := partsByRange{
		xvalues:         []int{1, 4000},
		mvalues:         []int{1, 4000},
		avalues:         []int{1, 4000},
		svalues:         []int{1, 4000},
		currentWorkflow: "in",
	}
	q := []partsByRange{}
	q = append(q, part)
	result := 0
	for len(q) > 0 {
		currentPart := q[0]
		q = q[1:]
		if currentPart.currentWorkflow == "R" {
			continue
		}
		if currentPart.currentWorkflow == "A" {
			result += currentPart.getTotal()
			continue
		}
		workflow := workflowMap[currentPart.currentWorkflow]
		for _, rule := range workflow.rules {
			if currentPart.empty {
				break
			}
			accepted, rejected := rule.evaluateByRange(currentPart)
			if !accepted.empty {
				q = append(q, accepted)
			}
			currentPart = rejected
		}
	}
	return result
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day19/input.txt")
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
