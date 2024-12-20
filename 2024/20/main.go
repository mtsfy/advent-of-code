package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Pair struct {
	row int
	col int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkBounds(r, c, m, n int) bool {
	return r >= 0 && r < m && c >= 0 && c < n
}

func parser(r *bufio.Scanner) ([][]string, Pair, Pair) {
	grid := make([][]string, 0)
	var start Pair
	var end Pair
	for r.Scan() {
		row := strings.Split(r.Text(), "")
		for col, cell := range row {
			if cell == "S" {
				start.row = len(grid)
				start.col = col
				break
			} else if cell == "E" {
				end.row = len(grid)
				end.col = col
				break
			}
		}
		grid = append(grid, row)
	}
	return grid, start, end
}

func part1(r *bufio.Scanner) int {
	res := 0
	grid, start, _ := parser(r)
	distances := calculateDistances(grid, start)
	var cheatDirections = []Pair{
		{2, 0},
		{1, 1},
		{0, 2},
		{-1, 1},
		{-2, 0},
		{-1, -1},
		{0, -2},
		{1, -1},
	}
	for i := range distances {
		for j := range distances {
			if distances[i][j] != -1 {
				for _, dir := range cheatDirections {
					newRow, newCol := i+dir.row, j+dir.col
					if !checkBounds(newRow, newCol, len(distances), len(distances[0])) {
						continue
					}

					if distances[newRow][newCol] == -1 {
						continue
					}

					if distances[i][j]-distances[newRow][newCol] >= 102 {
						res++
					}

				}
			}
		}
	}

	return res
}

func part2(r *bufio.Scanner) int {
	res := 0
	grid, start, end := parser(r)

	path := []Pair{start}
	dirs := []Pair{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	for path[len(path)-1] != end {
		curRow, curCol := path[len(path)-1].row, path[len(path)-1].col
		moved := false
		for _, dir := range dirs {
			newRow, newCol := curRow+dir.row, curCol+dir.col
			if !checkBounds(newRow, newCol, len(grid), len(grid[0])) {
				continue
			}

			if len(path) > 1 && newRow == path[len(path)-2].row && newCol == path[len(path)-2].col {
				continue
			}

			if grid[newRow][newCol] == "#" {
				continue
			}
			path = append(path, Pair{newRow, newCol})
			moved = true
			break
		}
		if !moved {
			break
		}
	}

	old := len(path) - 1
	distances := make(map[Pair]int)
	for i, pos := range path {
		distances[pos] = old - i
	}

	for i, cur := range path {
		curRow, curCol := cur.row, cur.col
		for newRow := curRow - 20; newRow <= curRow+20; newRow++ {
			for newCol := curCol - 20; newCol <= curCol+20; newCol++ {
				if !checkBounds(newRow, newCol, len(grid), len(grid[0])) {
					continue
				}
				if grid[newRow][newCol] == "#" {
					continue
				}

				distance := int(math.Abs(float64(newRow-curRow)) + math.Abs(float64(newCol-curCol)))
				if distance > 20 {
					continue
				}

				remainDist := distances[Pair{newRow, newCol}]
				if old-(i+remainDist+distance) >= 100 {
					res++
				}
			}
		}
	}
	return res
}

func calculateDistances(grid [][]string, start Pair) [][]int {
	distances := make([][]int, len(grid))
	for i := range grid {
		distances[i] = make([]int, len(grid[0]))
		for j := range grid[0] {
			distances[i][j] = -1
		}
	}

	distances[start.row][start.col] = 0
	curRow, curCol := start.row, start.col

	var dirs = []Pair{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	for grid[curRow][curCol] != "E" {
		for _, d := range dirs {
			nr, nc := curRow+d.row, curCol+d.col
			if !checkBounds(nr, nc, len(grid), len(grid[0])) {
				continue
			}

			if grid[nr][nc] == "#" || distances[nr][nc] != -1 {
				continue
			}
			distances[nr][nc] = distances[curRow][curCol] + 1
			curRow, curCol = nr, nc
		}
	}

	return distances
}

func main() {
	file, err := os.Open("../input/20.txt")
	check(err)
	defer file.Close()

	s1 := bufio.NewScanner(file)
	fmt.Println("part1:", part1(s1))

	file.Seek(0, 0)
	s2 := bufio.NewScanner(file)
	fmt.Println("part2:", part2(s2))
}
