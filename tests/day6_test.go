package tests

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/luxish/advent-of-code-2023/internal/day6"
)

const d6p1data string = `Time:      7  15   30
Distance:  9  40  200
`

func TestDay6Part1(t *testing.T) {
	result := day6.Part1(bufio.NewReader(bytes.NewReader([]byte(d6p1data))))
	if result != 288 {
		t.Errorf("The expected result is 288, got %v", result)
	}
}

func TestDay6Part2(t *testing.T) {
	result := day6.Part2(bufio.NewReader(bytes.NewReader([]byte(d6p1data))))
	if result != 71503 {
		t.Errorf("The expected result is 71503, got %v", result)
	}
}
