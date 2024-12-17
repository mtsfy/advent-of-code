package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var (
	A int
	B int
	C int
)

func parser(r *bufio.Scanner) [][]int {
	pattern := regexp.MustCompile(`-?\d+`)
	info := make([][]int, 0)
	for r.Scan() {
		line := r.Text()
		matches := pattern.FindAllString(line, -1)
		if len(matches) > 0 {
			numbers := make([]int, 0)
			for _, match := range matches {
				num, _ := strconv.Atoi(match)
				numbers = append(numbers, num)
			}
			info = append(info, numbers)
		}
	}
	return info
}

func comboOperand(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return A
	case 5:
		return B
	case 6:
		return C
	default:
		panic(fmt.Sprintf("Invalid combOperand: %d", operand))
	}
}

func part1(r *bufio.Scanner) string {
	info := parser(r)
	A, B, C = info[0][0], info[1][0], info[2][0]
	program := info[len(info)-1]

	pointer := 0
	results := make([]int, 0)

	for pointer < len(program) {
		instr := program[pointer]
		operand := program[pointer+1]

		switch instr {
		case 0:
			A = A >> comboOperand(operand)
		case 1:
			B = B ^ operand
		case 2:
			B = comboOperand(operand) % 8
		case 3:
			if A != 0 {
				pointer = operand
				continue
			}
		case 4:
			B = B ^ C
		case 5:
			results = append(results, comboOperand(operand)%8)
		case 6:
			B = A >> comboOperand(operand)
		case 7:
			C = A >> comboOperand(operand)
		}

		pointer += 2
	}
	res := ""
	for i, n := range results {
		if i > 0 {
			res += ","
		}
		res += fmt.Sprint(n)
	}
	return res
}

func main() {
	file, err := os.Open("../input/17.txt")
	check(err)
	defer file.Close()

	scanner1 := bufio.NewScanner(file)
	fmt.Println("part1:", part1(scanner1))
}
