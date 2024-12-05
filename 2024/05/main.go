package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func containsValue(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func part1(r *bufio.Scanner) int {
	res := 0
	contains := false
	isCorrect := true

	rules, updates := parser(r)

	for _, update := range updates {
		for i := 0; i < len(update); i++ {
			for j := i + 1; j < len(update); j++ {
				contains = containsValue(rules[update[i]], update[j])
				if !contains {
					isCorrect = false
					break
				}
			}
		}

		if isCorrect {
			midVal, err := strconv.ParseInt(update[int(len(update)/2)], 10, 64)
			check(err)
			res += int(midVal)
		}
		isCorrect = true
	}

	return res
}

func correctify(rules map[string][]string, update []string) []string {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(update)-1; i++ {
			pageNums, exists := rules[update[i+1]]
			if exists {
				if containsValue(pageNums, update[i]) {
					update[i], update[i+1] = update[i+1], update[i]
					swapped = true
					continue
				}
			}
			pageNums, exists = rules[update[i]]
			if exists {
				if !containsValue(pageNums, update[i+1]) {
					update[i], update[i+1] = update[i+1], update[i]
					swapped = true
				}
			}
		}
	}
	return update
}
func part2(r *bufio.Scanner) int {
	res := 0
	contains := false
	isCorrect := true

	rules, updates := parser(r)

	for _, update := range updates {
		for i := 0; i < len(update); i++ {
			for j := i + 1; j < len(update); j++ {
				contains = containsValue(rules[update[i]], update[j])
				if !contains {
					isCorrect = false
					break
				}
			}
		}
		if !isCorrect {
			correctUpdate := correctify(rules, update)
			midVal, err := strconv.ParseInt(correctUpdate[int(len(correctUpdate)/2)], 10, 64)
			check(err)
			res += int(midVal)
		}
		isCorrect = true
	}

	return res
}

func split(r rune) bool {
	return r == '|' || r == ','
}

func parser(r *bufio.Scanner) (map[string][]string, [][]string) {
	rules := make(map[string][]string)
	updates := make([][]string, 0)
	isReadingUpdates := false

	for r.Scan() {
		line := strings.FieldsFunc(r.Text(), split)
		if len(line) == 0 {
			isReadingUpdates = true
			continue
		}

		if isReadingUpdates {
			updates = append(updates, line)
		} else {
			rules[line[0]] = append(rules[line[0]], line[1])
		}

	}
	return rules, updates
}

func main() {
	file, err := os.Open("../input/05.txt")
	check(err)
	defer file.Close()

	scanner1 := bufio.NewScanner(file)
	fmt.Println("part1:", part1(scanner1))

	file.Seek(0, 0)
	scanner2 := bufio.NewScanner(file)
	fmt.Println("part2:", part2(scanner2))

}
