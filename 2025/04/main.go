package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/tsatam/adventofcode/common/cartesian"
	"github.com/tsatam/adventofcode/common/fp"
)

var (
	//go:embed input
	input string
)

type grid [][]int

func main() {
	fmt.Printf("Part 1: %d\n", handlePart1(input))
	fmt.Printf("Part 2: %d\n", handlePart2(input))
}

func handlePart1(input string) int {
	g := readInput(input)

	_, accessible := passthrough(g)
	return accessible
}

func handlePart2(input string) int {
	g := readInput(input)

	total := 0
	for g, removed := passthrough(g); removed > 0; g, removed = passthrough(g) {
		total += removed
	}
	return total
}

func passthrough(g grid) (grid, int) {
	removed := 0
	toRemove := []cartesian.Point{}
	for y := range len(g) {
		for x := range len(g[y]) {
			p := cartesian.Point{X: x, Y: y}
			if g.canAccessRoll(p) {
				removed++
				toRemove = append(toRemove, p)
			}
		}
	}

	for _, p := range toRemove {
		g.remove(p)
	}
	return g, removed
}

func readInput(input string) grid {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	return fp.Map(lines, func(in string) []int {
		return fp.Map([]rune(in), func(r rune) int {
			switch r {
			case '.':
				return 0
			case '@':
				return 1
			default:
				panic(fmt.Errorf("unexpected rune [%v]", r))
			}
		})
	})
}

func (g grid) remove(p cartesian.Point) {
	g[p.Y][p.X] = 0
}

func (g grid) canAccessRoll(p cartesian.Point) bool {
	if g.safeCheck(p) != 1 {
		return false
	}

	pos := []cartesian.Point{
		p.Move(cartesian.Up), p.Move(cartesian.Down), p.Move(cartesian.Left), p.Move(cartesian.Right),
		p.Move(cartesian.Up).Move(cartesian.Left), p.Move(cartesian.Up).Move(cartesian.Right),
		p.Move(cartesian.Down).Move(cartesian.Left), p.Move(cartesian.Down).Move(cartesian.Right),
	}
	checked := fp.Map(pos, g.safeCheck)
	sum := fp.Sum(checked)
	return sum < 4
}

func (g grid) safeCheck(p cartesian.Point) int {
	switch {
	case p.Y < 0,
		p.Y >= len(g),
		p.X < 0,
		p.X >= len(g[p.Y]):
		return 0
	default:
		return g[p.Y][p.X]
	}
}
