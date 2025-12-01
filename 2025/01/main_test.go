package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlePart1(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	want := 3
	got := handlePart1(input)

	assert.Equal(t, want, got)
}

func TestHandlePart2(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	want := 6
	got := handlePart2(input)

	assert.Equal(t, want, got)
}

func TestParseLine(t *testing.T) {
	for i, tt := range []struct {
		line string
		want int
	}{
		{line: "L68", want: -68},
		{line: "L30", want: -30},
		{line: "R48", want: 48},
		{line: "L5", want: -5},
		{line: "R60", want: 60},
		{line: "L55", want: -55},
		{line: "L1", want: -1},
		{line: "L99", want: -99},
		{line: "R14", want: 14},
		{line: "L82", want: -82},
	} {
		t.Run(fmt.Sprintf("[%d] %s -> %d", i, tt.line, tt.want), func(t *testing.T) {
			got := parseLine(tt.line)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPerform(t *testing.T) {
	for i, tt := range []struct {
		curr performed
		next int
		want performed
	}{
		{curr: performed{pos: 50}, next: -68, want: performed{pos: 82}},
		{curr: performed{pos: 82}, next: -30, want: performed{pos: 52}},
		{curr: performed{pos: 52}, next: 48, want: performed{pos: 0, timesZero: 1}},
		{curr: performed{pos: 0, timesZero: 1}, next: -5, want: performed{pos: 95, timesZero: 1}},
		{curr: performed{pos: 95, timesZero: 1}, next: 60, want: performed{pos: 55, timesZero: 1}},
		{curr: performed{pos: 55, timesZero: 1}, next: -55, want: performed{pos: 0, timesZero: 2}},
		{curr: performed{pos: 0, timesZero: 2}, next: -1, want: performed{pos: 99, timesZero: 2}},
		{curr: performed{pos: 99, timesZero: 2}, next: -99, want: performed{pos: 0, timesZero: 3}},
		{curr: performed{pos: 0, timesZero: 3}, next: 14, want: performed{pos: 14, timesZero: 3}},
		{curr: performed{pos: 14, timesZero: 3}, next: -82, want: performed{pos: 32, timesZero: 3}},
	} {
		t.Run(fmt.Sprintf("[%d] %v + %d -> %v", i, tt.curr, tt.next, tt.want), func(t *testing.T) {
			got := perform(tt.curr, tt.next)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPerformCLICK(t *testing.T) {
	for i, tt := range []struct {
		curr performed
		next int
		want performed
	}{
		{curr: performed{pos: 50}, next: -68, want: performed{pos: 82, timesZero: 1}},
		{curr: performed{pos: 82, timesZero: 1}, next: -30, want: performed{pos: 52, timesZero: 1}},
		{curr: performed{pos: 52, timesZero: 1}, next: 48, want: performed{pos: 0, timesZero: 2}},
		{curr: performed{pos: 0, timesZero: 2}, next: -5, want: performed{pos: 95, timesZero: 2}},
		{curr: performed{pos: 95, timesZero: 2}, next: 60, want: performed{pos: 55, timesZero: 3}},
		{curr: performed{pos: 55, timesZero: 3}, next: -55, want: performed{pos: 0, timesZero: 4}},
		{curr: performed{pos: 0, timesZero: 4}, next: -1, want: performed{pos: 99, timesZero: 4}},
		{curr: performed{pos: 99, timesZero: 4}, next: -99, want: performed{pos: 0, timesZero: 5}},
		{curr: performed{pos: 0, timesZero: 5}, next: 14, want: performed{pos: 14, timesZero: 5}},
		{curr: performed{pos: 14, timesZero: 5}, next: -82, want: performed{pos: 32, timesZero: 6}},

		{curr: performed{pos: 0, timesZero: 0}, next: -200, want: performed{pos: 0, timesZero: 2}},
		{curr: performed{pos: 0, timesZero: 0}, next: -201, want: performed{pos: 99, timesZero: 2}},
	} {
		t.Run(fmt.Sprintf("[%d] %v + %d -> %v", i, tt.curr, tt.next, tt.want), func(t *testing.T) {
			got := performCLICK(tt.curr, tt.next)
			assert.Equal(t, tt.want, got)
		})
	}
}
