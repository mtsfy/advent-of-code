package main

import (
	"bufio"
	"fmt"
	"math"
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

	for r.Scan() {
		line := strings.Fields(r.Text())
		level := make([]int64, len(line))

		for idx, str := range line {
			num, err := strconv.ParseInt(str, 10, 64)
			check(err)
			level[idx] = num
		}

		if checker(level) {
			res++
		}
	}
	return res
}

func part2(r *bufio.Scanner) int {
	res := 0

	for r.Scan() {
		line := strings.Fields(r.Text())
		level := make([]int64, len(line))

		for idx, str := range line {
			num, err := strconv.ParseInt(str, 10, 64)
			check(err)
			level[idx] = num
		}

		if checker(level) || remover(level) {
			res++
		}
	}
	return res
}
func checker(arr []int64) bool {
	if len(arr) <= 1 {
		return true
	}

	ascending := arr[0] < arr[1]

	for i := 0; i < len(arr)-1; i++ {
		diff := math.Abs(float64(arr[i] - arr[i+1]))
		if diff > 3 {
			return false
		}

		if ascending {
			if arr[i] >= arr[i+1] {
				return false
			}
		} else {
			if arr[i] <= arr[i+1] {
				return false
			}
		}
	}
	return true
}

func remover(arr []int64) bool {
	for i := 0; i < len(arr); i++ {
		temp := make([]int64, len(arr))
		copy(temp, arr)

		temp = append(temp[:i], temp[i+1:]...)

		if checker(temp) {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("../input/02.txt")
	check(err)
	defer file.Close()

	scanner1 := bufio.NewScanner(file)
	fmt.Println("part1:", part1(scanner1))

	file.Seek(0, 0)
	scanner2 := bufio.NewScanner(file)
	fmt.Println("part2:", part2(scanner2))
}
