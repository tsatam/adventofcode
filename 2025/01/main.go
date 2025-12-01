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

type performed struct {
	pos, timesZero int
}

func main() {
	fmt.Printf("Part 1: %d\n", handlePart1(input))
}

func handlePart1(input string) int {
	lines := readInput(input)
	rotations := fp.Map(lines, parseLine)
	perf := fp.Reduce(rotations, performed{pos: 50}, perform)
	return perf.timesZero
}

func readInput(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func parseLine(line string) int {
	var dir string
	var dist int
	if _, err := fmt.Sscanf(line, "%1s%d", &dir, &dist); err != nil {
		panic(err)
	}
	if dir == "L" {
		dist *= -1
	}
	return dist
}

func perform(curr performed, next int) performed {
	curr.pos += next
	if curr.pos < 0 {
		curr.pos += 100
	}
	curr.pos %= 100
	if curr.pos == 0 {
		curr.timesZero++
	}
	return curr
}
