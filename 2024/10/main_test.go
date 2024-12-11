package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`
	scanner := bufio.NewScanner(strings.NewReader(input))
	expected := 36
	assert.Equal(t, expected, part1(scanner))
}

func TestPart2(t *testing.T) {
	input := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`
	scanner := bufio.NewScanner(strings.NewReader(input))
	expected := 81
	assert.Equal(t, expected, part2(scanner))
}
