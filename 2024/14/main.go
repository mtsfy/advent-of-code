package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Robot struct {
	pX int
	pY int
	vX int
	vY int
}
type Pos struct {
	x int
	y int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1(r *bufio.Scanner, s int, w int, h int) int {
	res := 1
	robots := parser(r)
	midX, midY := int(w/2), int(h/2)
	quadrants := []int{0, 0, 0, 0}

	for _, robot := range robots {
		finalPos := calculateFinalPosition(robot, s, w, h)
		if finalPos.x == midX || finalPos.y == midY {
			continue
		}
		if finalPos.x > midX && finalPos.y > midY {
			quadrants[0]++
		} else if finalPos.x < midX && finalPos.y > midY {
			quadrants[1]++
		} else if finalPos.x < midX && finalPos.y < midY {
			quadrants[2]++
		} else if finalPos.x > midX && finalPos.y < midY {
			quadrants[3]++
		}
	}
	for _, count := range quadrants {
		res *= count
	}
	return res
}

func part2(r *bufio.Scanner, w int, h int) int {
	res := 1
	robots := parser(r)
	midX, midY := int(w/2), int(h/2)
	min := int(^uint(0) >> 1)

	for s := 0; s < w*h; s++ {
		temp := 1
		quadrants := []int{0, 0, 0, 0}
		for _, robot := range robots {
			finalPos := calculateFinalPosition(robot, s, w, h)
			if finalPos.x == midX || finalPos.y == midY {
				continue
			}
			if finalPos.x > midX && finalPos.y > midY {
				quadrants[0]++
			} else if finalPos.x < midX && finalPos.y > midY {
				quadrants[1]++
			} else if finalPos.x < midX && finalPos.y < midY {
				quadrants[2]++
			} else if finalPos.x > midX && finalPos.y < midY {
				quadrants[3]++
			}
		}
		for _, count := range quadrants {
			temp *= count
		}
		if temp < min {
			res = s
			min = temp
		}
	}
	return res
}

func calculateFinalPosition(r Robot, s int, w int, h int) Pos {
	xD := r.pX + s*r.vX
	yD := r.pY + s*r.vY

	xF := ((xD % w) + w) % w
	yF := ((yD % h) + h) % h

	return Pos{xF, yF}
}

func parser(r *bufio.Scanner) []Robot {
	robots := make([]Robot, 0)
	pattern := regexp.MustCompile(`-?\d+`)
	for r.Scan() {
		line := r.Text()
		matches := pattern.FindAllString(line, -1)
		pX, _ := strconv.Atoi(matches[0])
		pY, _ := strconv.Atoi(matches[1])
		vX, _ := strconv.Atoi(matches[2])
		vY, _ := strconv.Atoi(matches[3])
		robots = append(robots, Robot{pX, pY, vX, vY})
	}
	return robots
}

func main() {
	file, err := os.Open("../input/14.txt")
	check(err)
	defer file.Close()

	scanner1 := bufio.NewScanner(file)
	fmt.Println("part1:", part1(scanner1, 100, 101, 103))

	file.Seek(0, 0)
	scanner2 := bufio.NewScanner(file)
	fmt.Println("part2:", part2(scanner2, 101, 103))
}
