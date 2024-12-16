package day16

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

func Day16() {
	fmt.Println("--- Day 16: Reindeer Maze ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

type State struct {
	x, y, dir int
}

type Node struct {
	current, previous *State
	score             int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].score < pq[j].score
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

var directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func getGrid(lines []string) [][]byte {
	grid := [][]byte{}
	for _, l := range lines {
		grid = append(grid, []byte(l))
	}
	return grid
}

func shortestPath(grid [][]byte) (map[State]map[State]bool, int, []State) {
	score := make(map[State]int)
	parents := make(map[State]map[State]bool)
	bestScore := math.MaxInt32
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	pq.Push(&Node{current: &State{len(grid) - 2, 1, 0}, previous: &State{-1, -1, -1}, score: 0})
	bfsQ := []State{}
	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)
		current := node.current
		// if we have already seen this state with a better score
		if v, ok := score[*current]; ok && v < node.score {
			continue
		}
		score[*current] = node.score
		if _, ok := parents[*current]; !ok {
			parents[*current] = map[State]bool{}
		}
		// add the parent of the current state
		parents[*current][*node.previous] = true
		// if we have reached the end
		if current.x == 1 && current.y == len(grid[0])-2 {
			// we have exhausted all paths to reach this state, end the loop
			if node.score > bestScore {
				break
			}
			// save bestScore
			bestScore = node.score
			bfsQ = append(bfsQ, *current)
		}
		nextState := &State{current.x, current.y, (current.dir + 1) % 4}
		if v, ok := score[*nextState]; !ok || v >= node.score+1000 {
			heap.Push(&pq, &Node{current: nextState, previous: current, score: node.score + 1000})
		}
		nextState = &State{current.x, current.y, (current.dir + 3) % 4}
		if v, ok := score[*nextState]; !ok || v >= node.score+1000 {
			heap.Push(&pq, &Node{current: nextState, previous: current, score: node.score + 1000})
		}
		nx, ny := current.x+directions[current.dir][0], current.y+directions[current.dir][1]
		if grid[nx][ny] != '#' {
			nextState = &State{nx, ny, current.dir}
			if v, ok := score[*nextState]; !ok || v >= node.score+1 {
				heap.Push(&pq, &Node{current: nextState, previous: current, score: node.score + 1})
			}
		}
	}
	return parents, bestScore, bfsQ
}

func getPart1Answer(lines []string) int {
	grid := getGrid(lines)
	_, bestScore, _ := shortestPath(grid)
	return bestScore
}

func getPart2Answer(lines []string) int {
	grid := getGrid(lines)
	parents, _, bfsQ := shortestPath(grid)
	// bfs to find the number of paths and counts
	count := 0
	visited := make(map[State]bool)
	finalVisited := make([][]bool, len(grid))
	for i := range finalVisited {
		finalVisited[i] = make([]bool, len(grid[0]))
	}
	for len(bfsQ) > 0 {
		current := bfsQ[0]
		if current.x != -1 && !finalVisited[current.x][current.y] {
			finalVisited[current.x][current.y] = true
			count++
		}
		bfsQ = bfsQ[1:]
		if visited[current] {
			continue
		}
		visited[current] = true
		for parent := range parents[current] {
			if !visited[parent] {
				bfsQ = append(bfsQ, parent)
			}
		}
	}
	return count
}

func getInput() []string {
	lines := []string{}
	file, err := os.Open("day16/input.txt")
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
