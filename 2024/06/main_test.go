package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`
	scanner := bufio.NewScanner(strings.NewReader(input))
	expected := 41
	assert.Equal(t, expected, part1(scanner))
}

func TestPart2(t *testing.T) {
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`
	scanner := bufio.NewScanner(strings.NewReader(input))
	expected := 6
	assert.Equal(t, expected, part2(scanner))
}