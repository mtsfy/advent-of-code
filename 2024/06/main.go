package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pair struct {
	row int
	col int
}

var directionChange = map[Pair]Pair{
	{1, 0}:  {0, -1},
	{-1, 0}: {0, 1},
	{0, 1}:  {1, 0},
	{0, -1}: {-1, 0},
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1(r *bufio.Scanner) int {
	grid, starting := parser(r)
	res := searchPath(0, grid, starting, Pair{-1, 0})
	return res
}

func part2(r *bufio.Scanner) int {
	grid, starting := parser(r)
	count := 0

	visited := make(map[string]struct{})

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == starting.row && j == starting.col || grid[i][j] == "#" {
				continue
			}

			grid[i][j] = "#"

			for k := range visited {
				delete(visited, k)
			}

			if checkObstruction(grid, starting, Pair{-1, 0}, visited) {
				count++
			}

			grid[i][j] = "."
		}
	}
	return count
}

func checkObstruction(grid [][]string, position Pair, direction Pair, visited map[string]struct{}) bool {
	for {
		stateKey := fmt.Sprintf("%d,%d,%d,%d", position.row, position.col, direction.row, direction.col)
		_, exists := visited[stateKey]

		if exists {
			return true
		}

		visited[stateKey] = struct{}{}

		newRow := position.row + direction.row
		newCol := position.col + direction.col

		if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) {
			return false
		}

		if grid[newRow][newCol] == "#" {
			direction = directionChange[direction]
		} else {
			position = Pair{newRow, newCol}
		}
	}
}

func searchPath(steps int, grid [][]string, current Pair, direction Pair) int {
	currentRow, currentCol := current.row, current.col

	if grid[currentRow][currentCol] != "1" {
		grid[currentRow][currentCol] = "1"
		steps++
	}

	newRow := current.row + direction.row
	newCol := current.col + direction.col

	if newRow >= len(grid) || newRow < 0 || newCol >= len(grid[0]) || newCol < 0 {
		grid[currentRow][currentCol] = "1"
		return steps
	}

	if grid[newRow][newCol] == "#" {
		newDirection := directionChange[direction]
		return searchPath(steps, grid, current, newDirection)
	} else {
		return searchPath(steps, grid, Pair{newRow, newCol}, direction)
	}
}

func parser(r *bufio.Scanner) ([][]string, Pair) {
	grid := make([][]string, 0)
	var starting Pair

	for r.Scan() {
		line := strings.Split(r.Text(), "")

		for col, cell := range line {
			if cell == "^" {
				starting = Pair{len(grid), col}
			}
		}
		grid = append(grid, line)
	}

	return grid, starting
}

func main() {
	file, err := os.Open("../input/06.txt")
	check(err)
	defer file.Close()

	scanner1 := bufio.NewScanner(file)
	fmt.Println("part1:", part1(scanner1))

	file.Seek(0, 0)
	scanner2 := bufio.NewScanner(file)
	fmt.Println("part2:", part2(scanner2))
}
