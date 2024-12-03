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
	res := 0

	for r.Scan() {
		text := r.Text()
		res = calulate(text)
	}
	return res
}

func part2(r *bufio.Scanner) int {
	res := 0
	re := regexp.MustCompile(`don't\(\).*?do\(\)|don't\(\).*`)
	for r.Scan() {
		text := r.Text()
		text = re.ReplaceAllString(text, "")
		res = calulate(text)
	}
	return res
}

func calulate(text string) int {
	res := 0
	mainRe := regexp.MustCompile(`mul\(\s*\d+\s*,\s*\d+\s*\)`)
	subRe := regexp.MustCompile(`mul\(\s*(\d+)\s*,\s*(\d+)\s*\)`)

	for _, match := range mainRe.FindAllStringSubmatch(text, -1) {
		nums := subRe.FindStringSubmatch(match[0])

		num1, err1 := strconv.ParseInt(nums[1], 10, 64)
		check(err1)
		num2, err2 := strconv.ParseInt(nums[2], 10, 64)
		check(err2)

		res += int(num1 * num2)
	}

	return res
}

func main() {
	file, err := os.Open("../input/03.txt")
	check(err)
	defer file.Close()

	scanner1 := bufio.NewScanner(file)
	fmt.Println("part1:", part1(scanner1))

	file.Seek(0, 0)
	scanner2 := bufio.NewScanner(file)
	fmt.Println("part2:", part2(scanner2))

}
