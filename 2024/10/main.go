package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	row int
	col int
}

var directions = []Pair{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parser(r *bufio.Scanner) ([][]int, []Pair) {
	grid := make([][]int, 0)
	zeroLocation := make([]Pair, 0)
	i := 0
	for r.Scan() {
		i++
		line := strings.Split(r.Text(), "")
		row := make([]int, 0)

		for j, cell := range line {
			if cell == "0" {
				zeroLocation = append(zeroLocation, Pair{i - 1, j})
			}
			num, err := strconv.Atoi(cell)
			check(err)
			row = append(row, num)
		}
		grid = append(grid, row)
	}
	return grid, zeroLocation
}

func part1(r *bufio.Scanner) int {
	res := 0
	grid, zeroLocations := parser(r)

	for _, start := range zeroLocations {
		visited := make(map[string]struct{})
		uniquePeaks := make(map[string]struct{})
		score := 0

		for _, dir := range directions {
			nextPos := Pair{
				row: start.row + dir.row,
				col: start.col + dir.col,
			}
			score += findScore(grid, nextPos, 0, &visited, &uniquePeaks)
		}
		res += len(uniquePeaks)
	}

	return res
}

func findScore(grid [][]int, current Pair, previousHeight int, visited *map[string]struct{}, uniquePeaks *map[string]struct{}) int {
	if checkOutOfBounds(current.row, current.col, len(grid), len(grid[0])) {
		return 0
	}

	currentHeight := grid[current.row][current.col]
	if currentHeight != previousHeight+1 {
		return 0
	}

	key := fmt.Sprintf("%d,%d", current.row, current.col)
	_, isVisited := (*visited)[key]
	if isVisited {
		return 0
	}

	(*visited)[key] = struct{}{}

	score := 0

	if currentHeight == 9 {
		if _, exists := (*uniquePeaks)[key]; !exists {
			(*uniquePeaks)[key] = struct{}{}
			score = 1
		}
	}

	for _, dir := range directions {
		next := Pair{
			row: current.row + dir.row,
			col: current.col + dir.col,
		}
		score += findScore(grid, next, currentHeight, visited, uniquePeaks)
	}

	delete(*visited, key)
	return score
}

func part2(r *bufio.Scanner) int {
	res := 0
	grid, zeroLocations := parser(r)

	for _, start := range zeroLocations {
		visited := make(map[string]struct{})
		uniquePaths := make(map[string]struct{})
		startPath := fmt.Sprintf("%d,%d", start.row, start.col)

		for _, dir := range directions {
			nextPos := Pair{
				row: start.row + dir.row,
				col: start.col + dir.col,
			}
			findPaths(grid, nextPos, 0, startPath, &visited, &uniquePaths)
		}

		res += len(uniquePaths)
	}
	return res
}

func findPaths(grid [][]int, current Pair, previousHeight int, currentPath string, visited *map[string]struct{}, uniquePaths *map[string]struct{}) {
	if checkOutOfBounds(current.row, current.col, len(grid), len(grid[0])) {
		return
	}

	currentHeight := grid[current.row][current.col]
	if currentHeight != previousHeight+1 {
		return
	}

	key := fmt.Sprintf("%d,%d", current.row, current.col)
	_, isVisited := (*visited)[key]
	if isVisited {
		return
	}

	(*visited)[key] = struct{}{}

	newPath := currentPath + ">" + key

	if currentHeight == 9 {
		(*uniquePaths)[newPath] = struct{}{}
		delete(*visited, key)
		return
	}

	for _, dir := range directions {
		nextPos := Pair{
			row: current.row + dir.row,
			col: current.col + dir.col,
		}
		findPaths(grid, nextPos, currentHeight, newPath, visited, uniquePaths)
	}

	delete(*visited, key)
}

func checkOutOfBounds(row, col, m, n int) bool {
	return row < 0 || row >= m || col < 0 || col >= n
}

func main() {
	file, err := os.Open("../input/10.txt")
	check(err)
	defer file.Close()

	scanner1 := bufio.NewScanner(file)
	fmt.Println("part1:", part1(scanner1))

	file.Seek(0, 0)
	scanner2 := bufio.NewScanner(file)
	fmt.Println("part2:", part2(scanner2))
}
