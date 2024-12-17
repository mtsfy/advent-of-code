package main

import (
	"bufio"
	"container/heap"
	"container/list"
	"fmt"
	"math"
	"os"
	"strings"
)

type Pair struct {
	row int
	col int
}

type Item struct {
	cost   int
	row    int
	col    int
	dirRow int
	dirCol int
	index  int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].cost != pq[j].cost {
		return pq[i].cost < pq[j].cost
	}

	if pq[i].row != pq[j].row {
		return pq[i].row < pq[j].row
	}

	if pq[i].col != pq[j].col {
		return pq[i].col < pq[j].col
	}

	if pq[i].dirRow != pq[j].dirRow {
		return pq[i].dirRow < pq[j].dirRow
	}

	return pq[i].dirCol < pq[j].dirCol
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parser(r *bufio.Scanner) ([][]string, Pair) {
	grid := make([][]string, 0)
	var start Pair
	for r.Scan() {
		row := strings.Split(r.Text(), "")
		for col, cell := range row {
			if cell == "S" {
				start.row = len(grid)
				start.col = col
				break
			}
		}
		grid = append(grid, row)
	}
	return grid, start
}

func part1(r *bufio.Scanner) int {
	grid, start := parser(r)

	pq := make(PriorityQueue, 0)
	visited := make(map[string]bool)

	heap.Push(&pq, &Item{
		cost:   0,
		row:    start.row,
		col:    start.col,
		dirRow: 0,
		dirCol: 1,
		index:  0,
	})

	key := fmt.Sprintf("%d_%d_%d_%d", start.row, start.col, 0, 1)
	visited[key] = true

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)

		key := fmt.Sprintf("%d_%d_%d_%d", item.row, item.col, item.dirRow, item.dirCol)
		visited[key] = true

		if grid[item.row][item.col] == "E" {
			return item.cost
		}

		nextPositions := []struct {
			cost   int
			row    int
			col    int
			dirRow int
			dirCol int
		}{
			{
				cost:   item.cost + 1,
				row:    item.row + item.dirRow,
				col:    item.col + item.dirCol,
				dirRow: item.dirRow,
				dirCol: item.dirCol,
			},
			{
				cost:   item.cost + 1000,
				row:    item.row,
				col:    item.col,
				dirRow: item.dirCol,
				dirCol: -item.dirRow,
			},
			{
				cost:   item.cost + 1000,
				row:    item.row,
				col:    item.col,
				dirRow: -item.dirCol,
				dirCol: item.dirRow,
			},
		}

		for _, nextPos := range nextPositions {
			if grid[nextPos.row][nextPos.col] == "#" {
				continue
			}
			key := fmt.Sprintf("%d_%d_%d_%d", nextPos.row, nextPos.col, nextPos.dirRow, nextPos.dirCol)
			if visited[key] {
				continue
			}
			heap.Push(&pq, &Item{
				cost:   nextPos.cost,
				row:    nextPos.row,
				col:    nextPos.col,
				dirRow: nextPos.dirRow,
				dirCol: nextPos.dirCol,
			})
		}
	}
	return -1
}

func part2(r *bufio.Scanner) int {
	grid, start := parser(r)

	pq := make(PriorityQueue, 0)
	lowestCost := make(map[string]int)
	backtrack := make(map[string]map[string]bool)
	endStates := make(map[string]bool)
	bestCost := math.MaxInt32

	heap.Push(&pq, &Item{
		cost:   0,
		row:    start.row,
		col:    start.col,
		dirRow: 0,
		dirCol: 1,
		index:  0,
	})

	lowestCost[fmt.Sprintf("%d_%d_%d_%d", start.row, start.col, 0, 1)] = 0

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)

		key := fmt.Sprintf("%d_%d_%d_%d", item.row, item.col, item.dirRow, item.dirCol)
		if storedCost, exists := lowestCost[key]; exists && item.cost > storedCost {
			continue
		}

		if grid[item.row][item.col] == "E" {
			if item.cost > bestCost {
				break
			}
			bestCost = item.cost
			endStateKey := fmt.Sprintf("%d_%d_%d_%d", item.row, item.col, item.dirRow, item.dirCol)
			endStates[endStateKey] = true
		}

		nextPositions := []struct {
			cost   int
			row    int
			col    int
			dirRow int
			dirCol int
		}{
			{
				cost:   item.cost + 1,
				row:    item.row + item.dirRow,
				col:    item.col + item.dirCol,
				dirRow: item.dirRow,
				dirCol: item.dirCol,
			},
			{
				cost:   item.cost + 1000,
				row:    item.row,
				col:    item.col,
				dirRow: item.dirCol,
				dirCol: -item.dirRow,
			},
			{
				cost:   item.cost + 1000,
				row:    item.row,
				col:    item.col,
				dirRow: -item.dirCol,
				dirCol: item.dirRow,
			},
		}

		for _, nextPos := range nextPositions {
			if grid[nextPos.row][nextPos.col] == "#" {
				continue
			}

			nextKey := fmt.Sprintf("%d_%d_%d_%d", nextPos.row, nextPos.col, nextPos.dirRow, nextPos.dirCol)
			currentKey := fmt.Sprintf("%d_%d_%d_%d", item.row, item.col, item.dirRow, item.dirCol)

			lowest := lowestCost[nextKey]
			if nextPos.cost > lowest && lowest != 0 {
				continue
			}

			if nextPos.cost < lowest || lowest == 0 {
				backtrack[nextKey] = make(map[string]bool)
				lowestCost[nextKey] = nextPos.cost
			}

			if backtrack[nextKey] == nil {
				backtrack[nextKey] = make(map[string]bool)
			}
			backtrack[nextKey][currentKey] = true

			heap.Push(&pq, &Item{
				cost:   nextPos.cost,
				row:    nextPos.row,
				col:    nextPos.col,
				dirRow: nextPos.dirRow,
				dirCol: nextPos.dirCol,
			})
		}
	}

	seen := make(map[string]bool)
	states := list.New()
	for state := range endStates {
		states.PushBack(state)
		seen[state] = true
	}

	for states.Len() > 0 {
		elem := states.Front()
		key := elem.Value.(string)
		states.Remove(elem)

		for lastState := range backtrack[key] {
			if !seen[lastState] {
				seen[lastState] = true
				states.PushBack(lastState)
			}
		}
	}

	uniquePositions := make(map[string]bool)
	for state := range seen {
		parts := strings.Split(state, "_")
		posKey := fmt.Sprintf("%s_%s", parts[0], parts[1])
		uniquePositions[posKey] = true
	}

	return len(uniquePositions)
}

func main() {
	file, err := os.Open("../input/16.txt")
	check(err)
	defer file.Close()

	scanner1 := bufio.NewScanner(file)
	fmt.Println("part1:", part1(scanner1))
	file.Seek(0, 0)
	scanner2 := bufio.NewScanner(file)
	fmt.Println("part2:", part2(scanner2))
}
