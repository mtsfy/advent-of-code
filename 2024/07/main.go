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

func part1(r *bufio.Scanner) int {
	res := 0
	equations := parser(r)

	for _, equation := range equations {
		target := equation[0]
		numbers := equation[1:]
		result := 0
		search(numbers, 1, numbers[0], target, &result)
		res += result
	}
	return res
}
func search(numbers []int, idx int, current int, target int, result *int) bool {
	if idx == len(numbers) {
		if current == target {
			*result += current
			return true
		}
		return false
	}

	if search(numbers, idx+1, current+numbers[idx], target, result) {
		return true
	}
	return search(numbers, idx+1, current*numbers[idx], target, result)
}

func part2(r *bufio.Scanner) int {
	res := 0
	equations := parser(r)

	for _, equation := range equations {
		target := equation[0]
		numbers := equation[1:]
		result := 0
		searchNew(numbers, 1, numbers[0], target, &result)
		res += result
	}
	return res
}

func searchNew(numbers []int, idx int, current int, target int, result *int) bool {
	if idx == len(numbers) {
		if current == target {
			*result += current
			return true
		}
		return false
	}

	if searchNew(numbers, idx+1, current+numbers[idx], target, result) {
		return true
	}
	if searchNew(numbers, idx+1, current*numbers[idx], target, result) {
		return true
	}
	return searchNew(numbers, idx+1, concatNums(current, numbers[idx]), target, result)
}

func concatNums(current int, next int) int {
	strNum := strconv.Itoa(current) + strconv.Itoa(next)
	num, err := strconv.ParseInt(strNum, 10, 64)
	check(err)
	return int(num)
}

func parser(r *bufio.Scanner) [][]int {
	equations := make([][]int, 0)

	for r.Scan() {
		line := strings.Split(r.Text(), " ")
		expected64, err := strconv.ParseInt(line[0][:len(line[0])-1], 10, 64)
		check(err)

		equation := make([]int, len(line))
		equation[0] = int(expected64)

		for i, operand := range line[1:] {
			num, err := strconv.ParseInt(operand, 10, 64)
			check(err)
			equation[i+1] = int(num)
		}
		equations = append(equations, equation)
	}

	return equations
}

func main() {
	file, err := os.Open("../input/07.txt")
	check(err)
	defer file.Close()

	scanner1 := bufio.NewScanner(file)
	fmt.Println("part1:", part1(scanner1))

	file.Seek(0, 0)
	scanner2 := bufio.NewScanner(file)
	fmt.Println("part2:", part2(scanner2))
}
