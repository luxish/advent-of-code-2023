package tests

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/luxish/advent-of-code-2023/internal/day8"
)

const d8p1data string = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`

const d8p2data string = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`

func TestDay8Part1(t *testing.T) {
	result := day8.Part1(bufio.NewReader(bytes.NewReader([]byte(d8p1data))))
	if result != 6 {
		t.Errorf("The expected result is 6, got %v", result)
	}
}

func TestDay8Part2(t *testing.T) {
	result := day8.Part2(bufio.NewReader(bytes.NewReader([]byte(d8p2data))))
	if result != 6 {
		t.Errorf("The expected result is 6, got %v", result)
	}
}
