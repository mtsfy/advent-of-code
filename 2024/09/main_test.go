package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := `2333133121414131402`
	scanner := bufio.NewScanner(strings.NewReader(input))
	expected := 1928
	assert.Equal(t, expected, part1(scanner))
}

func TestPart2(t *testing.T) {
	input := `2333133121414131402`
	scanner := bufio.NewScanner(strings.NewReader(input))
	expected := 2858
	assert.Equal(t, expected, part2(scanner))
}
