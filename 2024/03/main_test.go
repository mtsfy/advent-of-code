package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

	scanner := bufio.NewScanner(strings.NewReader(input))
	expected := 161
	assert.Equal(t, expected, part1(scanner))
}

func TestPart2(t *testing.T) {
	input := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

	scanner := bufio.NewScanner(strings.NewReader(input))
	expected := 48
	assert.Equal(t, expected, part2(scanner))
}
