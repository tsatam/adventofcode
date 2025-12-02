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

func TestSumInvalidIds(t *testing.T) {
	for i, tt := range []struct {
		idRange string
		want    int
	}{
		{idRange: "11-22", want: 33},
		{idRange: "95-115", want: 99},
		{idRange: "998-1012", want: 1010},
		{idRange: "1188511880-1188511890", want: 1188511885},
		{idRange: "222220-222224", want: 222222},
		{idRange: "1698522-1698528"},
		{idRange: "446443-446449", want: 446446},
		{idRange: "38593856-38593862", want: 38593859},
		{idRange: "565653-565659"},
		{idRange: "824824821-824824827"},
		{idRange: "2121212118-2121212124"},
	} {
		t.Run(fmt.Sprintf("[%d] %s -> %d", i, tt.idRange, tt.want), func(t *testing.T) {
			got := sumInvalidIds(tt.idRange)

			assert.Equal(t, tt.want, got)
		})
	}

}
