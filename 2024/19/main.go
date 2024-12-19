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

func parser(r *bufio.Scanner) ([]string, []string) {
	var designs, patterns []string
	var readDesigns bool

	for r.Scan() {
		line := r.Text()
		if len(strings.Split(line, "")) == 0 {
			readDesigns = true
			continue
		}
		if readDesigns {
			designs = append(designs, line)
			continue
		} else {
			patterns = append(patterns, strings.Split(line, ", ")...)
		}
	}

	return designs, patterns
}

func exists(design string, patterns []string) bool {
	for _, pattern := range patterns {
		if pattern == design {
			return true
		}
	}
	return false
}

func part1(r *bufio.Scanner) int {
	var res int
	designs, patterns := parser(r)

	var maxlen int
	for _, pattern := range patterns {
		if len(pattern) > maxlen {
			maxlen = len(pattern)
		}
	}
	cache := make(map[string]bool)
	for _, design := range designs {
		if valid(design, patterns, maxlen, &cache) {
			res++
		}
	}
	return res
}

func valid(design string, patterns []string, maxlen int, cache *map[string]bool) bool {
	if val, exists := (*cache)[design]; exists {
		return val
	}
	if design == "" {
		return true
	}
	for i := range min(maxlen, len(design)) + 1 {
		if exists(design[:i], patterns) && valid(design[i:], patterns, maxlen, cache) {
			(*cache)[design] = true
			return true
		}
	}
	(*cache)[design] = false
	return false
}

func part2(r *bufio.Scanner) int {
	var res int
	designs, patterns := parser(r)

	var maxlen int
	for _, pattern := range patterns {
		if len(pattern) > maxlen {
			maxlen = len(pattern)
		}
	}
	cache := make(map[string]int)
	for _, design := range designs {
		res += arrangments(design, patterns, maxlen, &cache)
	}
	return res
}

func arrangments(design string, patterns []string, maxlen int, cache *map[string]int) int {
	if count, exists := (*cache)[design]; exists {
		return count
	}

	if design == "" {
		return 1
	}
	var count int
	for i := range min(maxlen, len(design)) + 1 {
		if exists(design[:i], patterns) {
			count += arrangments(design[i:], patterns, maxlen, cache)
			(*cache)[design] = count
		}
	}
	(*cache)[design] = count
	return count
}

func main() {
	file, err := os.Open("../input/19.txt")
	check(err)
	defer file.Close()

	s1 := bufio.NewScanner(file)
	fmt.Println("part1:", part1(s1))

	file.Seek(0, 0)
	s2 := bufio.NewScanner(file)
	fmt.Println("part2:", part2(s2))
}
