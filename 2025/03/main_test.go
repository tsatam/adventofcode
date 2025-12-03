package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlePart1(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`

	want := 357
	got := handlePart1(input)

	assert.Equal(t, want, got)
}

func TestReadInput(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`

	want := []string{
		"987654321111111",
		"811111111111119",
		"234234234234278",
		"818181911112111",
	}
	got := readInput(input)

	assert.Equal(t, want, got)
}

func TestJoltageForBank(t *testing.T) {
	for i, tt := range []struct {
		bank string
		want int
	}{
		{bank: "987654321111111", want: 98},
		{bank: "811111111111119", want: 89},
		{bank: "234234234234278", want: 78},
		{bank: "818181911112111", want: 92},
	} {
		t.Run(fmt.Sprintf("[%d] %s -> %d", i, tt.bank, tt.want), func(t *testing.T) {
			got := joltageForBank(tt.bank)

			assert.Equal(t, tt.want, got)
		})
	}
}
