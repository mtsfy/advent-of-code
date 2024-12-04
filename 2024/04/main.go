package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func exists(temp string, re *regexp.Regexp) int {
	if re.MatchString(temp) {
		return 1
	}
	return 0
}

func parser(r *bufio.Scanner) [][]string {
	words := make([][]string, 0)
	for r.Scan() {
		line := strings.Split(r.Text(), "")
		words = append(words, line)
	}
	return words
}

func part1(r *bufio.Scanner) int {
	res := 0
	words := parser(r)

	var directions = [][]int{
		{0, 1},
		{1, 0},
		{1, 1},
		{1, -1},
		{0, -1},
		{-1, 0},
		{-1, -1},
		{-1, 1},
	}

	re := regexp.MustCompile("XMAS")

	temp := ""

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {

			for _, dir := range directions {
				if i+3*dir[0] >= 0 && i+3*dir[0] < len(words) && j+3*dir[1] >= 0 && j+3*dir[1] < len(words[0]) {

					for k := 0; k < 4; k++ {
						temp += words[i+k*dir[0]][j+k*dir[1]]
					}

					res += exists(temp, re)
					temp = ""
				}
			}

		}
	}
	return res
}

func part2(r *bufio.Scanner) int {
	count := 0
	words := parser(r)

	re := regexp.MustCompile("MAS|SAM")

	temp1, temp2 := "", ""

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			if i-1 >= 0 && i+1 < len(words) && j-1 >= 0 && j+1 < len(words[0]) {
				if words[i][j] == "A" {
					temp1 += words[i-1][j+1]
					temp1 += words[i][j]
					temp1 += words[i+1][j-1]

					temp2 += words[i-1][j-1]
					temp2 += words[i][j]
					temp2 += words[i+1][j+1]

					if exists(temp1, re) == 1 && exists(temp2, re) == 1 {
						count++
					}
				}
				temp1, temp2 = "", ""
			}
		}
	}
	return count
}

func main() {
	file, err := os.Open("../input/04.txt")
	check(err)
	defer file.Close()

	scanner1 := bufio.NewScanner(file)
	fmt.Println("part1:", part1(scanner1))

	file.Seek(0, 0)
	scanner2 := bufio.NewScanner(file)
	fmt.Println("part2:", part2(scanner2))
}
