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

var dirs = map[string]Pair{
	">": {0, 1},
	"^": {-1, 0},
	"v": {1, 0},
	"<": {0, -1},
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkBounds(r, c, m, n int) bool {
	return r >= 0 && r < m && c >= 0 && c < n
}

func parser(r *bufio.Scanner) ([][]string, []string, Pair) {
	grid := make([][]string, 0)
	steps := make([]string, 0)

	readSteps := false
	var startPos Pair
	for r.Scan() {
		line := strings.Split(r.Text(), "")

		if len(line) == 0 {
			readSteps = true
		}
		if readSteps {
			steps = append(steps, line...)
		} else {
			grid = append(grid, line)
			for col, char := range line {
				if char == "@" {
					startPos = Pair{len(grid) - 1, col}
				}
			}
		}
	}
	return grid, steps, startPos
}

func part1(r *bufio.Scanner) int {
	res := 0
	grid, steps, robPos := parser(r)
	for _, step := range steps {
		move(&grid, &robPos, dirs[step])
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "O" {
				res += (100*i + j)
			}
		}
	}
	return res
}

func move(grid *[][]string, cur *Pair, dir Pair) {
	newRow, newCol := cur.row+dir.row, cur.col+dir.col
	if !checkBounds(newRow, newCol, len(*grid), len((*grid)[0])) {
		return
	}
	i, j := cur.row, cur.col
	for checkBounds(i, j, len(*grid), len((*grid)[0])) {
		i += dir.row
		j += dir.col

		if !checkBounds(i, j, len(*grid), len((*grid)[0])) || (*grid)[i][j] == "#" {
			break
		}
		if (*grid)[i][j] == "." {
			(*grid)[i][j] = "O"
			(*grid)[cur.row][cur.col] = "."
			cur.row += dir.row
			cur.col += dir.col
			(*grid)[cur.row][cur.col] = "@"
			break
		}
	}
}

func main() {
	file, err := os.Open("../input/15.txt")
	check(err)
	defer file.Close()

	scanner1 := bufio.NewScanner(file)
	fmt.Println("part1:", part1(scanner1))
	// TODO: finish part 2 later
}
