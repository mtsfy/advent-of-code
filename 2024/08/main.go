package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parser(r *bufio.Scanner) ([][]string, map[string][][]int) {
	grid := [][]string{}
	antennas := make(map[string][][]int)
	row := 0
	for r.Scan() {
		line := strings.Split(r.Text(), "")
		grid = append(grid, line)
		for col, val := range line {
			if val != "." {
				antennas[val] = append(antennas[val], []int{row, col})
			}
		}
		row++
	}
	return grid, antennas
}

func part1(r *bufio.Scanner) int {
	grid, antennas := parser(r)
	n := len(grid)
	antinodeSet := make(map[string]struct{})

	for _, points := range antennas {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				antinodes := getAntinodes(points[i], points[j], n)
				for _, antinode := range antinodes {
					key := fmt.Sprintf("%d,%d", antinode[0], antinode[1])
					antinodeSet[key] = struct{}{}
				}
			}
		}
	}
	return len(antinodeSet)
}

func part2(r *bufio.Scanner) int {
	grid, antennas := parser(r)
	n := len(grid)
	antinodeSet := make(map[string]struct{})

	for _, points := range antennas {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				antinodes := getAllAntiondes(points[i], points[j], n)
				for _, antinode := range antinodes {
					key := fmt.Sprintf("%d,%d", antinode[0], antinode[1])
					antinodeSet[key] = struct{}{}
				}
			}
		}
	}
	return len(antinodeSet)
}

func getAntinodes(p1, p2 []int, n int) [][]int {
	x1, y1 := p1[0], p1[1]
	x2, y2 := p2[0], p2[1]

	x3, y3 := x1-(x2-x1), y1-(y2-y1)
	x4, y4 := x2+(x2-x1), y2+(y2-y1)

	antinodes := [][]int{}

	if x3 >= 0 && x3 < n && y3 >= 0 && y3 < n {
		antinodes = append(antinodes, []int{x3, y3})
	}
	if x4 >= 0 && x4 < n && y4 >= 0 && y4 < n {
		antinodes = append(antinodes, []int{x4, y4})
	}
	return antinodes
}

func getAllAntiondes(p1, p2 []int, n int) [][]int {
	x1, y1 := p1[0], p1[1]
	x2, y2 := p2[0], p2[1]

	antinodes := [][]int{}
	i := 0
	for {
		x3, y3 := x1-(x2-x1)*i, y1-(y2-y1)*i
		if x3 >= 0 && x3 < n && y3 >= 0 && y3 < n {
			antinodes = append(antinodes, []int{x3, y3})
		} else {
			break
		}
		i++
	}

	i = 0
	for {
		x4, y4 := x2+(x2-x1)*i, y2+(y2-y1)*i
		if x4 >= 0 && x4 < n && y4 >= 0 && y4 < n {
			antinodes = append(antinodes, []int{x4, y4})
		} else {
			break
		}
		i++
	}
	return antinodes
}

func main() {
	file, err := os.Open("../input/08.txt")
	check(err)
	defer file.Close()

	scanner1 := bufio.NewScanner(file)
	fmt.Println("part1:", part1(scanner1))

	file.Seek(0, 0)
	scanner2 := bufio.NewScanner(file)
	fmt.Println("part2:", part2(scanner2))
}
