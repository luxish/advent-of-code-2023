package tests

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/luxish/advent-of-code-2023/internal/day3"
)

const d3p1data string = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`

func TestDay3Part1(t *testing.T) {
	result := day3.Day3Part1(bufio.NewReader(bytes.NewReader([]byte(d3p1data))))
	if result != 4361 {
		t.Errorf("The expected result is 4361, got %v", result)
	}
}

func TestDay3Part2(t *testing.T) {
	result := day3.Day3Part2(bufio.NewReader(bytes.NewReader([]byte(d3p1data))))
	if result != 467835 {
		t.Errorf("The expected result is 467835, got %v", result)
	}
}
