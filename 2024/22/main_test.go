package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := `1
10
100
2024`
	scanner := bufio.NewScanner(strings.NewReader(input))
	expected := 37327623
	assert.Equal(t, expected, part1(scanner))
}

func TestPart2(t *testing.T) {
	input := `1
2
3
2024`
	scanner := bufio.NewScanner(strings.NewReader(input))
	expected := 23
	assert.Equal(t, expected, part2(scanner))
}
