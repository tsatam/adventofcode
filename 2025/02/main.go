package main

import (
	_ "embed"
	"fmt"
	"strconv"
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
	ranges := readInput(input)
	invalidIds := fp.Map(ranges, sumInvalidIds)
	return fp.Sum(invalidIds)
}

func readInput(input string) []string {
	return strings.Split(strings.TrimSpace(input), ",")
}

func sumInvalidIds(idRange string) int {
	split := strings.Split(idRange, "-")
	left, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}
	right, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}
	sum := 0
	for id := left; id <= right; id++ {
		sum += invalidId(id)
	}
	return sum
}

func invalidId(id int) int {
	idStr := strconv.Itoa(id)
	switch {
	case len(idStr)%2 != 0,
		idStr[len(idStr)/2:] != idStr[:len(idStr)/2]:
		return 0
	default:
		return id
	}
}
