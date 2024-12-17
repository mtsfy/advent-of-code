package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`
	scanner := bufio.NewScanner(strings.NewReader(input))
	expected := "4,6,3,5,6,3,5,2,1,0"
	assert.Equal(t, expected, part1(scanner))
}
