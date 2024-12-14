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

func part1(r *bufio.Scanner) int {
	machines := parser(r)
	return calculateTokens(machines, 0)
}

func part2(r *bufio.Scanner) int {
	machines := parser(r)
	return calculateTokens(machines, 10000000000000)
}

func calculateTokens(machines [][]int, add int) int {
	totalTokens := 0
	for _, machine := range machines {
		aX, aY := machine[0], machine[1]
		bX, bY := machine[2], machine[3]
		pX, pY := machine[4]+add, machine[5]+add

		countA := float64(pX*bY-pY*bX) / float64(aX*bY-aY*bX)
		countB := (float64(pX) - float64(aX)*countA) / float64(bX)

		if countA == float64(int(countA)) && countB == float64(int(countB)) {
			if countA >= 0 && countB >= 0 {
				totalTokens += int(countA*3 + countB)
			}
		}
	}
	return totalTokens
}

func parser(r *bufio.Scanner) [][]int {
	machines := make([][]int, 0)
	machine := make([]int, 0)
	pattern := regexp.MustCompile(`-?\d+`)

	for r.Scan() {
		line := r.Text()
		if line == "" {
			if len(machine) > 0 {
				machines = append(machines, machine)
				machine = make([]int, 0)
			}
			continue
		}

		matches := pattern.FindAllString(line, -1)
		for _, match := range matches {
			num, _ := strconv.Atoi(match)
			machine = append(machine, num)
		}
	}
	if len(machine) > 0 {
		machines = append(machines, machine)
	}
	return machines
}
func main() {
	file, err := os.Open("../input/13.txt")
	check(err)
	defer file.Close()

	scanner1 := bufio.NewScanner(file)
	fmt.Println("part1:", part1(scanner1))

	file.Seek(0, 0)
	scanner2 := bufio.NewScanner(file)
	fmt.Println("part2:", part2(scanner2))
}
