package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/tsatam/adventofcode/common/fp"
)

var (
	//go:embed input
	input string
)

func main() {
	fmt.Printf("Part 1: %d\n", handlePart1(input))
}

func handlePart1(input string) int {
	banks := readInput(input)
	joltages := fp.Map(banks, joltageForBank)
	return fp.Sum(joltages)
}

func readInput(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func joltageForBank(bank string) int {

	largestOne := -1
	largestOneIdx := -1
	for i := 0; largestOne < 9 && i < len(bank)-1; i++ {
		curr := int(bank[i] - '0')
		if curr > largestOne {
			largestOne = curr
			largestOneIdx = i
		}
	}

	largestTwo := -1
	for i := largestOneIdx + 1; largestTwo < 9 && i < len(bank); i++ {
		curr := int(bank[i] - '0')
		if curr > largestTwo {
			largestTwo = curr
		}
	}

	return 10*largestOne + largestTwo
}
