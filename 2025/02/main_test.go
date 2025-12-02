package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlePart1(t *testing.T) {
	input := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

	want := 1227775554
	got := handlePart1(input)

	assert.Equal(t, want, got)
}

func TestHandlePart2(t *testing.T) {
	input := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

	want := 4174379265
	got := handlePart2(input)

	assert.Equal(t, want, got)
}

func TestReadInput(t *testing.T) {
	input := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

	want := []string{
		"11-22",
		"95-115",
		"998-1012",
		"1188511880-1188511890",
		"222220-222224",
		"1698522-1698528",
		"446443-446449",
		"38593856-38593862",
		"565653-565659",
		"824824821-824824827",
		"2121212118-2121212124",
	}
	got := readInput(input)

	assert.Equal(t, want, got)
}

func TestRangeToIds(t *testing.T) {
	for i, tt := range []struct {
		idRange string
		want    []int
	}{
		{idRange: "11-22", want: []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}},
		{idRange: "95-115", want: []int{95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115}},
	} {
		t.Run(fmt.Sprintf("[%d] %s -> %v", i, tt.idRange, tt.want), func(t *testing.T) {
			got := rangeToIds(tt.idRange)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInvalidIdPart1(t *testing.T) {
	for _, tt := range []struct {
		id   int
		want bool
	}{
		{id: 11, want: true},
		{id: 22, want: true},
		{id: 95},
		{id: 99, want: true},
		{id: 111},
	} {
		t.Run(fmt.Sprintf("%d -> %v", tt.id, tt.want), func(t *testing.T) {
			got := invalidIdPart1(tt.id)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInvalidIdPart2(t *testing.T) {
	for _, tt := range []struct {
		id   int
		want bool
	}{
		{id: 11, want: true},
		{id: 22, want: true},
		{id: 95},
		{id: 99, want: true},
		{id: 111, want: true},
	} {
		t.Run(fmt.Sprintf("%d -> %v", tt.id, tt.want), func(t *testing.T) {
			got := invalidIdPart2(tt.id)

			assert.Equal(t, tt.want, got)
		})
	}

}
