package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := strings.Split("125 17", " ")
	expected := 22
	assert.Equal(t, expected, part1(input, 6))
}

func TestPart2(t *testing.T) {
	input := strings.Split("125 17", " ")
	expected := 55312
	assert.Equal(t, expected, part2(input, 25))
}
