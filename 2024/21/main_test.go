package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := `029A
980A
179A
456A
379A`
	scanner := bufio.NewScanner(strings.NewReader(input))
	expected := 126384
	assert.Equal(t, expected, part1(scanner))
}
