package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func part1(input []string, blinkCount int) int {
	return calculateBlinks(input, blinkCount)
}

// very slow & inefficient
func calculateBlinks(nums []string, blinkCount int) int {
	stones := make([][]string, 0)
	stones = append(stones, nums)

	for i := 0; i < blinkCount; i++ {
		temp := make([]string, 0)
		for j := 0; j < len(stones[0]); j++ {
			currentStones := stones[0]

			currentStone := currentStones[j]
			numDigits := len(currentStones[j])

			if currentStone == "0" {
				temp = append(temp, "1")
			} else if numDigits%2 == 0 {
				temp = append(temp, cleanNum(currentStone[0:numDigits/2]))
				temp = append(temp, cleanNum(currentStone[numDigits/2:]))
			} else {
				num, _ := strconv.Atoi(currentStone)
				temp = append(temp, cleanNum(strconv.Itoa(num*2024)))
			}
		}
		stones = stones[:0]
		stones = append(stones, temp)
	}
	return len(stones[0])
}

func part2(input []string, blinkCount int) int {
	res := 0
	cache := make(map[string]int)
	for _, num := range input {
		res += calculateBlinksFast(num, blinkCount, &cache)
	}
	return res
}

func calculateBlinksFast(num string, blinkCount int, cache *map[string]int) int {
	key := fmt.Sprintf("%s-%d", num, blinkCount)
	if value, exists := (*cache)[key]; exists {
		return value
	}

	var length int
	numDigits := len(num)

	if blinkCount == 0 {
		length = 1
	} else if num == "0" {
		length = calculateBlinksFast("1", blinkCount-1, cache)
	} else if numDigits%2 == 0 {
		left := cleanNum(num[:numDigits/2])
		right := cleanNum(num[numDigits/2:])
		length = calculateBlinksFast(left, blinkCount-1, cache) + calculateBlinksFast(right, blinkCount-1, cache)
	} else {
		numBig, _ := new(big.Int).SetString(num, 10)
		newVal := new(big.Int).Mul(numBig, big.NewInt(2024))
		length = calculateBlinksFast(newVal.String(), blinkCount-1, cache)
	}
	(*cache)[key] = length
	return length
}

func cleanNum(num string) string {
	if num == strings.Repeat("0", len(num)) {
		return "0"
	}
	return strings.TrimLeft(num, "0")
}

func main() {
	input := strings.Split("41078 18 7 0 4785508 535256 8154 447", " ")
	fmt.Println("part1:", part1(input, 25))
	fmt.Println("part2:", part2(input, 75))
}
