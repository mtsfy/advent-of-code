package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1(left []int, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	res := 0
	diff := 0

	for i := 0; i < len(left); i++ {
		diff = int(math.Abs(float64(right[i] - left[i])))
		res += diff
	}

	return res
}

func part2(left []int, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	res := 0
	counts := make(map[int]int)

	for _, num := range right {
		counts[num] = counts[num] + 1
	}

	for i := 0; i < len(left); i++ {
		res += (left[i] * counts[left[i]])
	}

	return res
}

func readInput(r *bufio.Scanner) ([]int, []int, error) {
	left := []int{}
	right := []int{}

	for r.Scan() {
		line := strings.Fields(r.Text())

		l, lerr := strconv.Atoi(line[0])
		if lerr != nil {
			return nil, nil, lerr
		}
		rVal, rerr := strconv.Atoi(line[1])
		if rerr != nil {
			return nil, nil, rerr
		}

		left = append(left, l)
		right = append(right, rVal)
	}

	if err := r.Err(); err != nil {
		return nil, nil, err
	}

	return left, right, nil
}

func main() {
	file, err := os.Open("../input/01.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	left, right, err := readInput(scanner)
	check(err)

	fmt.Println("part1:", part1(left, right))
	fmt.Println("part2:", part2(left, right))

}
