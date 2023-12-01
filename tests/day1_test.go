package tests

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/luxish/advent-of-code-2023/internal/day1"
)

const data1 string = `
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

const data2 string = `
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`

func TestDay1Part1(t *testing.T) {
	result := day1.Day1Part1(bufio.NewReader(bytes.NewReader([]byte(data1))))
	if result != 142 {
		t.Errorf("The expected result is 142, got %v", result)
	}
}

func TestDay1Part2(t *testing.T) {
	result := day1.Day1Part2(bufio.NewReader(bytes.NewReader([]byte(data2))))
	if result != 281 {
		t.Errorf("The expected result is 281, got %v", result)
	}
}
