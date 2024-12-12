package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

type Pair struct {
	row float64
	col float64
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

func part1(r *bufio.Scanner) int {
	res := 0
	grid := parser(r)
	regions := findRegions(grid)
	area := area(regions)
	perimeter := perimeter(regions)

	for i := 0; i < len(regions); i++ {
		res += area[i] * perimeter[i]
	}
	return res
}

func part2(r *bufio.Scanner) int {
	res := 0
	grid := parser(r)
	regions := findRegions(grid)
	area := area(regions)
	sides := sides(regions)

	for i := 0; i < len(regions); i++ {
		res += area[i] * sides[i]
	}
	return res
}

func findRegions(grid [][]string) []map[Pair]struct{} {
	m := len(grid)
	n := len(grid[0])
	visited := make(map[Pair]bool)
	regions := make([]map[Pair]struct{}, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[Pair{float64(i), float64(j)}] {
				continue
			}

			currentPlant := grid[i][j]
			region := make(map[Pair]struct{})

			q := list.New()
			q.PushBack(Pair{float64(i), float64(j)})

			visited[Pair{float64(i), float64(j)}] = true
			region[Pair{float64(i), float64(j)}] = struct{}{}

			for q.Len() > 0 {
				pair := q.Front().Value.(Pair)
				q.Remove(q.Front())
				row, col := pair.row, pair.col

				for _, dir := range directions {
					newRow, newCol := row+dir.row, col+dir.col
					if checkOutOfBounds(int(newRow), int(newCol), m, n) ||
						grid[int(newRow)][int(newCol)] != currentPlant ||
						visited[Pair{newRow, newCol}] {
						continue
					}

					visited[Pair{newRow, newCol}] = true
					region[Pair{newRow, newCol}] = struct{}{}
					q.PushBack(Pair{newRow, newCol})
				}
			}

			if len(region) > 0 {
				regions = append(regions, region)
			}
		}
	}
	return regions
}

func area(regions []map[Pair]struct{}) []int {
	areas := make([]int, 0)
	for _, region := range regions {
		areas = append(areas, len(region))
	}
	return areas
}

func perimeter(regions []map[Pair]struct{}) []int {
	perimeters := make([]int, 0)
	for _, region := range regions {
		perimeter := 0
		for pos := range region {
			for _, dir := range directions {
				newRow, newCol := pos.row+dir.row, pos.col+dir.col
				if _, exists := region[Pair{newRow, newCol}]; !exists {
					perimeter++
				}
			}
		}
		perimeters = append(perimeters, perimeter)
	}
	return perimeters
}

func sides(regions []map[Pair]struct{}) []int {
	sides := make([]int, len(regions))
	for i, region := range regions {
		tempCorners := make(map[Pair]struct{})
		for pos := range region {
			offsets := []Pair{
				{pos.row - 0.5, pos.col - 0.5},
				{pos.row + 0.5, pos.col - 0.5},
				{pos.row + 0.5, pos.col + 0.5},
				{pos.row - 0.5, pos.col + 0.5},
			}
			for _, offset := range offsets {
				tempCorners[offset] = struct{}{}
			}
		}

		corners := 0
		for corner := range tempCorners {
			occupancy := make([]bool, 4)
			cornerChecks := []Pair{
				{corner.row - 0.5, corner.col - 0.5},
				{corner.row + 0.5, corner.col - 0.5},
				{corner.row + 0.5, corner.col + 0.5},
				{corner.row - 0.5, corner.col + 0.5},
			}

			for j, check := range cornerChecks {
				_, occupancy[j] = region[check]
			}

			sum := 0
			for _, occupied := range occupancy {
				if occupied {
					sum++
				}
			}

			if sum == 1 {
				corners++
			} else if sum == 2 {
				if (occupancy[0] && occupancy[2] && !occupancy[1] && !occupancy[3]) ||
					(!occupancy[0] && !occupancy[2] && occupancy[1] && occupancy[3]) {
					corners += 2
				}
			} else if sum == 3 {
				corners++
			}
		}
		sides[i] = corners
	}
	return sides
}
func checkOutOfBounds(row, col, m, n int) bool {
	return row < 0 || row >= m || col < 0 || col >= n
}

func parser(r *bufio.Scanner) [][]string {
	grid := make([][]string, 0)
	for r.Scan() {
		row := strings.Split(r.Text(), "")
		grid = append(grid, row)
	}
	return grid
}

func main() {
	file, err := os.Open("../input/12.txt")
	check(err)
	defer file.Close()

	scanner1 := bufio.NewScanner(file)
	fmt.Println("part1:", part1(scanner1))

	file.Seek(0, 0)
	scanner2 := bufio.NewScanner(file)
	fmt.Println("part2:", part2(scanner2))

}
