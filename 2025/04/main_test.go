package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tsatam/adventofcode/common/cartesian"
)

func TestHandlePart1(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	want := 13
	got := handlePart1(input)

	assert.Equal(t, want, got)
}

func TestReadInput(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	want := grid{
		{0, 0, 1, 1, 0, 1, 1, 1, 1, 0},
		{1, 1, 1, 0, 1, 0, 1, 0, 1, 1},
		{1, 1, 1, 1, 1, 0, 1, 0, 1, 1},
		{1, 0, 1, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 1, 0, 1, 1},
		{0, 1, 1, 1, 1, 1, 1, 1, 0, 1},
		{0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
		{1, 0, 1, 1, 1, 0, 1, 1, 1, 1},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0, 1, 0},
	}

	got := readInput(input)

	assert.Equal(t, want, got)
}

func TestCanAccessRoll(t *testing.T) {
	g := grid{
		{0, 0, 1, 1, 0, 1, 1, 1, 1, 0},
		{1, 1, 1, 0, 1, 0, 1, 0, 1, 1},
		{1, 1, 1, 1, 1, 0, 1, 0, 1, 1},
		{1, 0, 1, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 1, 0, 1, 1},
		{0, 1, 1, 1, 1, 1, 1, 1, 0, 1},
		{0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
		{1, 0, 1, 1, 1, 0, 1, 1, 1, 1},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0, 1, 0},
	}

	for i, tt := range []struct {
		p    cartesian.Point
		want bool
	}{
		{p: cartesian.Point{X: 0, Y: 0}, want: false},
		{p: cartesian.Point{X: 2, Y: 0}, want: true},
		{p: cartesian.Point{X: 7, Y: 0}, want: false},
	} {
		t.Run(fmt.Sprintf("[%d] %v -> %v", i, tt.p, tt.want), func(t *testing.T) {
			got := g.canAccessRoll(tt.p)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSafeCheck(t *testing.T) {
	g := grid{
		{0, 1},
		{1, 0},
	}

	for i, tt := range []struct {
		p    cartesian.Point
		want int
	}{
		{p: cartesian.Point{X: -1, Y: 0}, want: 0},
		{p: cartesian.Point{X: 0, Y: -1}, want: 0},
		{p: cartesian.Point{X: 0, Y: 0}, want: 0},
		{p: cartesian.Point{X: 1, Y: 0}, want: 1},
	} {
		t.Run(fmt.Sprintf("[%d] %v -> %d", i, tt.p, tt.want), func(t *testing.T) {
			got := g.safeCheck(tt.p)

			assert.Equal(t, tt.want, got)
		})
	}
}
