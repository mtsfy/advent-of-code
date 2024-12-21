package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Pos struct {
	row int
	col int
}

type Pair struct {
	x any
	y any
}

type Item struct {
	pos  Pos
	path string
}

var numbers = [][]string{
	{"7", "8", "9"},
	{"4", "5", "6"},
	{"1", "2", "3"},
	{"X", "0", "A"},
}

var directions = [][]string{
	{"X", "^", "A"},
	{"<", "v", ">"},
}

var moves = []Item{
	{Pos{-1, 0}, "^"},
	{Pos{1, 0}, "v"},
	{Pos{0, -1}, "<"},
	{Pos{0, 1}, ">"},
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkBounds(r, c, m, n int) bool {
	return r >= 0 && r < m && c >= 0 && c < n
}

func part1(r *bufio.Scanner) int {
	res := 0
	numSequences := sequences(numbers)
	dirSequences := sequences(directions)
	pattern := regexp.MustCompile(`-?\d+`)

	for r.Scan() {
		code := r.Text()
		num, err := strconv.Atoi(pattern.FindAllString(code, -1)[0])
		check(err)

		inital := presses(code, numSequences)
		next := inital

		for i := 0; i < 2; i++ {
			var choices []string
			for _, move := range next {
				choices = append(choices, presses(move, dirSequences)...)
			}

			minLen := math.MaxInt
			for i := range choices {
				if len(choices[i]) < minLen {
					minLen = len(choices[i])
				}
			}

			var filtered []string
			for _, move := range choices {
				if len(move) == minLen {
					filtered = append(filtered, move)
				}
			}
			next = filtered
		}
		res += len(next[0]) * num
	}

	return res
}

func sequences(keypad [][]string) map[Pair][]string {
	positions := make(map[string]Pos)
	for i := range keypad {
		for j := range keypad[i] {
			if keypad[i][j] != "X" {
				positions[keypad[i][j]] = Pos{i, j}
			}
		}
	}

	sequences := make(map[Pair][]string)
	for x := range positions {
		for y := range positions {
			if x == y {
				sequences[Pair{x, y}] = []string{"A"}
				continue
			}

			var choices []string
			queue := list.New()
			shortest := math.MaxInt
			findingPaths := true

			queue.PushBack(Item{positions[x], ""})

			for queue.Len() > 0 && findingPaths {
				cur := queue.Remove(queue.Front()).(Item)
				curRow, curCol := cur.pos.row, cur.pos.col
				curPath := cur.path

				for _, move := range moves {
					newRow, newCol := curRow+move.pos.row, curCol+move.pos.col
					newPath := curPath + move.path

					if !checkBounds(newRow, newCol, len(keypad), len(keypad[0])) {
						continue
					}
					if keypad[newRow][newCol] == "X" {
						continue
					}

					if keypad[newRow][newCol] == y {
						if shortest < len(newPath) {
							findingPaths = false
							break
						}
						shortest = len(newPath)
						choices = append(choices, newPath+"A")
					} else {
						queue.PushBack(Item{Pos{newRow, newCol}, newPath})
					}
				}
			}
			sequences[Pair{x, y}] = choices
		}
	}
	return sequences
}
func product(lists [][]string) []string {
	if len(lists) == 0 {
		return []string{""}
	}
	result := []string{}
	subProduct := product(lists[1:])
	for _, item := range lists[0] {
		for _, p := range subProduct {
			result = append(result, item+p)
		}
	}
	return result
}
func presses(code string, sequences map[Pair][]string) []string {
	options := make([][]string, len(code))

	for i := 0; i < len(code); i++ {
		prev := "A"
		if i > 0 {
			prev = string(code[i-1])
		}
		curr := string(code[i])
		options[i] = sequences[Pair{prev, curr}]
	}

	return product(options)
}

func main() {
	file, err := os.Open("../input/21.txt")
	check(err)
	defer file.Close()

	s1 := bufio.NewScanner(file)
	fmt.Println("part1:", part1(s1))
}
