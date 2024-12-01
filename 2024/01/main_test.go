package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := `3   4
	4   3
	2   5
	1   3
	3   9
	3   3`

	scanner := bufio.NewScanner(strings.NewReader(input))

	left, right, err := readInput(scanner)
	assert.NoError(t, err, "Failed to read input")

	expected := 11
	assert.Equal(t, expected, part1(left, right))
}

func TestPart2(t *testing.T) {
	input := `3   4
	4   3
	2   5
	1   3
	3   9
	3   3`

	scanner := bufio.NewScanner(strings.NewReader(input))

	left, right, err := readInput(scanner)
	assert.NoError(t, err, "Failed to read input")

	expected := 31
	assert.Equal(t, expected, part2(left, right))
}
