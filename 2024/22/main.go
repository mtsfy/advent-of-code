package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1(r *bufio.Scanner) int {
	res := 0
	for r.Scan() {
		num, err := strconv.Atoi(r.Text())
		check(err)
		for i := 0; i < 2000; i++ {
			num = calculate(num)
		}
		res += num
	}
	return res
}

func part2(r *bufio.Scanner) int {
	totals := make(map[[4]int]int)

	for r.Scan() {
		num, err := strconv.Atoi(r.Text())
		check(err)
		buyer := []int{num % 10}

		for i := 0; i < 2000; i++ {
			num = calculate(num)
			buyer = append(buyer, num%10)
		}

		seen := make(map[[4]int]bool)
		for i := 0; i < len(buyer)-4; i++ {
			sequence := [4]int{buyer[i+1] - buyer[i], buyer[i+2] - buyer[i+1], buyer[i+3] - buyer[i+2], buyer[i+4] - buyer[i+3]}
			if seen[sequence] {
				continue
			}

			seen[sequence] = true
			if _, exists := totals[sequence]; !exists {
				totals[sequence] = 0
			}
			totals[sequence] += buyer[i+4]
		}
	}

	max := 0
	for _, total := range totals {
		if total > max {
			max = total
		}
	}
	return max
}

func calculate(n int) int {
	n = (n ^ (n * 64)) % 16777216
	n = (n ^ (n / 32)) % 16777216
	n = (n ^ (n * 2048)) % 16777216
	return n
}

func main() {
	file, err := os.Open("../input/22.txt")
	check(err)
	defer file.Close()

	s1 := bufio.NewScanner(file)
	fmt.Println("part1:", part1(s1))

	file.Seek(0, 0)
	s2 := bufio.NewScanner(file)
	fmt.Println("part2:", part2(s2))
}
