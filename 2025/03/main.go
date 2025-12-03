package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"

	"github.com/tsatam/adventofcode/common/fp"
)

var (
	//go:embed input
	input string
)

func main() {
	fmt.Printf("Part 1: %d\n", handlePart1(input))
	fmt.Printf("Part 2: %d\n", handlePart2(input))
}

func handlePart1(input string) int {
	banks := readInput(input)
	joltages := fp.Map(banks, joltageForBank(2))
	return fp.Sum(joltages)
}

func handlePart2(input string) int {
	banks := readInput(input)
	joltages := fp.Map(banks, joltageForBank(12))
	return fp.Sum(joltages)
}

func readInput(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func joltageForBank(digits int) func(bank string) int {
	return func(bank string) int {
		largests := make([]int, digits)
		largestsIdx := make([]int, digits)
		for i := range largestsIdx {
			largestsIdx[i] = -1
		}

		for digit := digits - 1; digit >= 0; digit-- {
			start := 0
			if digit < digits-1 && largestsIdx[digit+1] != -1 {
				start = largestsIdx[digit+1] + 1
			}

			for i := start; largests[digit] < 9 && i < len(bank)-digit; i++ {
				curr := int(bank[i] - '0')
				if curr > largests[digit] {
					largests[digit] = curr
					largestsIdx[digit] = i
				}
			}
		}

		value := 0
		for digit := 0; digit < digits; digit = digit + 1 {
			value += largests[digit] * int(math.Pow10(digit))
		}
		return value
	}
}
