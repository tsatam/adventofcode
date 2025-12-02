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
	fmt.Printf("Part 2: %d\n", handlePart2(input))
}

func handlePart1(input string) int {
	rawRanges := readInput(input)
	ranges := fp.Map(rawRanges, rangeToIds)
	ids := fp.Flatten(ranges)
	invalidIds := fp.Filter(ids, invalidIdPart1)
	return fp.Sum(invalidIds)
}

func handlePart2(input string) int {
	rawRanges := readInput(input)
	ranges := fp.Map(rawRanges, rangeToIds)
	ids := fp.Flatten(ranges)
	invalidIds := fp.Filter(ids, invalidIdPart2)
	return fp.Sum(invalidIds)
}

func readInput(input string) []string {
	return strings.Split(strings.TrimSpace(input), ",")
}

func rangeToIds(idRange string) []int {
	split := strings.Split(idRange, "-")
	left, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}
	right, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}

	ids := make([]int, right-left+1)
	for i := range ids {
		ids[i] = left + i
	}

	return ids
}

func invalidIdPart1(id int) bool {
	idStr := strconv.Itoa(id)
	return len(idStr)%2 == 0 &&
		idStr[len(idStr)/2:] == idStr[:len(idStr)/2]
}

func invalidIdPart2(id int) bool {
	idStr := strconv.Itoa(id)
	for patternLength := 1; patternLength <= len(idStr)/2; patternLength++ {
		if len(idStr)%patternLength == 0 {
			if checkRepeating(idStr, patternLength) {
				return true
			}
		}
	}
	return false
}

func checkRepeating(id string, patternLength int) bool {
	toMatch := id[:patternLength]
	for j := patternLength; j < len(id); j += patternLength {
		if id[j:j+patternLength] != toMatch {
			return false
		}
	}
	return true
}
