package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`
	scanner := bufio.NewScanner(strings.NewReader(input))
	expected := 22
	assert.Equal(t, expected, part1(scanner, 6, 12))
}

func TestPart2(t *testing.T) {
	input := `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`
	scanner := bufio.NewScanner(strings.NewReader(input))
	expected := "6,1"
	assert.Equal(t, expected, part2(scanner, 6))
}
