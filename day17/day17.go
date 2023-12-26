package day17

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func Day17() {
	fmt.Println("--- Day 17: Clumsy Crucible ---")
	lines := getInput()
	part1Answer := getPart1Answer(lines)
	fmt.Println("Part1:", part1Answer)
	part2Answer := getPart2Answer(lines)
	fmt.Println("Part2:", part2Answer)
}

func getPart1Answer(lines []string) int {
	grid := getGrid(lines)
	return runDjikstra(grid, 1, 3)
}

type state struct {
	x          int
	y          int
	dx         int
	dy         int
	moves      int
	distance   int
	totalMoves int
}

func (s state) getStateKey() string {
	return fmt.Sprintf("%d,%d,%d,%d,%d", s.x, s.y, s.dx, s.dy, s.moves)
}

type minheap []state

func (h minheap) Len() int           { return len(h) }
func (h minheap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h minheap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minheap) Push(x interface{}) {
	*h = append(*h, x.(state))
}
func (h *minheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func runDjikstra(grid [][]int, minMoves, maxMoves int) int {
	n := len(grid)
	m := len(grid[0])
	visited := make(map[string]bool)
	pq := &minheap{state{x: 0, y: 0, dx: 0, dy: 0, moves: 0, totalMoves: 0}}
	heap.Init(pq)
	for len(*pq) > 0 {
		current := heap.Pop(pq).(state)
		if current.x == n-1 && current.y == m-1 && current.moves >= minMoves && current.moves <= maxMoves {
			return current.distance
		}
		currentDistance := current.distance
		key := current.getStateKey()
		if _, ok := visited[key]; ok {
			continue
		}
		visited[key] = true
		if (current.dx != 0 || current.dy != 0) && current.moves <= maxMoves-1 {
			nx := current.x + current.dx
			ny := current.y + current.dy
			if nx >= 0 && nx < n && ny >= 0 && ny < m {
				heap.Push(pq, state{
					x:          nx,
					y:          ny,
					dx:         current.dx,
					dy:         current.dy,
					moves:      current.moves + 1,
					distance:   currentDistance + grid[nx][ny],
					totalMoves: current.totalMoves + 1,
				})
			}
		}
		if current.totalMoves != 0 && current.moves < minMoves {
			continue
		}
		for _, d := range [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			if d[0] == current.dx && d[1] == current.dy {
				continue
			}
			if d[0] == -current.dx && d[1] == -current.dy {
				continue
			}
			nx := current.x + d[0]
			ny := current.y + d[1]
			if nx >= 0 && nx < n && ny >= 0 && ny < m {
				heap.Push(pq, state{
					x:          nx,
					y:          ny,
					dx:         d[0],
					dy:         d[1],
					moves:      1,
					distance:   currentDistance + grid[nx][ny],
					totalMoves: current.totalMoves + 1,
				})
			}
		}
	}
	return 0
}

func getGrid(lines []string) [][]int {
	n := len(lines)
	m := len(lines[0])
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			grid[i] = append(grid[i], int(lines[i][j]-'0'))
		}
	}
	return grid
}

func getPart2Answer(lines []string) int {
	grid := getGrid(lines)
	return runDjikstra(grid, 4, 10)
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
