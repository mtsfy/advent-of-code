package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	row int
	col int
}

type Data struct {
	current Pair
	path    []Pair
}

var directions = []Pair{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkBounds(r, c, m, n int) bool {
	return r >= 0 && r < m && c >= 0 && c < n
}

func part1(r *bufio.Scanner, size, limit int) int {
	grid := make([][]string, size+1)
	for i := range grid {
		grid[i] = make([]string, size+1)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}
	count := 0
	for r.Scan() && count < limit {
		line := strings.Split(r.Text(), ",")
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		grid[y][x] = "#"
		count++
	}
	path := search(grid)

	/*
		for _, point := range path {
			grid[point.row][point.col] = "O"
		}
		for _, row := range grid {
			fmt.Println(row)
		}
	*/
	return len(path) - 1
}

func search(grid [][]string) []Pair {
	visited := make(map[string]bool)
	queue := list.New()
	start := Pair{0, 0}
	queue.PushBack(Data{start, []Pair{start}})

	for queue.Len() > 0 {
		data := queue.Remove(queue.Front()).(Data)
		cur := data.current
		path := data.path

		if !checkBounds(cur.row, cur.col, len(grid), len(grid[0])) {
			continue
		}
		if grid[cur.row][cur.col] == "#" {
			continue
		}

		key := fmt.Sprintf("%d_%d", cur.row, cur.col)
		if visited[key] {
			continue
		}
		visited[key] = true

		if cur.row == len(grid)-1 && cur.col == len(grid[0])-1 {
			return path
		}

		for _, dir := range directions {
			new := Pair{cur.row + dir.row, cur.col + dir.col}
			newKey := fmt.Sprintf("%d_%d", new.row, new.col)
			if !visited[newKey] {
				newPath := make([]Pair, len(path))
				copy(newPath, path)
				queue.PushBack(Data{new, append(newPath, new)})
			}
		}
	}
	return nil
}

func part2(r *bufio.Scanner, size int) string {
	grid := make([][]string, size+1)
	for i := range grid {
		grid[i] = make([]string, size+1)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	points := make([]Pair, 0)
	for r.Scan() {
		line := strings.Split(r.Text(), ",")
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		points = append(points, Pair{x, y})
	}

	l, h := 0, len(points)
	for l < h {
		m := (l + h) / 2

		for i := range grid {
			for j := range grid[i] {
				grid[i][j] = "."
			}
		}

		for _, point := range points[:m] {
			grid[point.row][point.col] = "#"
		}

		if !valid(grid, size) {
			h = m
		} else {
			l = m + 1
		}
	}
	return fmt.Sprintf("%d,%d", points[l-1].row, points[l-1].col)
}

func valid(grid [][]string, size int) bool {
	visited := make(map[string]bool)
	queue := list.New()
	start := Pair{0, 0}
	queue.PushBack(start)

	for queue.Len() > 0 {
		cur := queue.Remove(queue.Front()).(Pair)

		if !checkBounds(cur.row, cur.col, size+1, size+1) {
			continue
		}
		if grid[cur.row][cur.col] == "#" {
			continue
		}

		key := fmt.Sprintf("%d_%d", cur.row, cur.col)
		if visited[key] {
			continue
		}
		visited[key] = true

		if cur.row == size && cur.col == size {
			return true
		}

		for _, dir := range directions {
			new := Pair{cur.row + dir.row, cur.col + dir.col}
			newKey := fmt.Sprintf("%d_%d", new.row, new.col)
			if !visited[newKey] {
				queue.PushBack(new)
			}
		}
	}
	return false
}

func main() {
	file, err := os.Open("../input/18.txt")
	check(err)
	defer file.Close()

	s1 := bufio.NewScanner(file)
	fmt.Println("part1:", part1(s1, 70, 1024))

	file.Seek(0, 0)
	s2 := bufio.NewScanner(file)
	fmt.Println("part2:", part2(s2, 70))
}
