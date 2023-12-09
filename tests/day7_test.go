package tests

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/luxish/advent-of-code-2023/internal/day7"
)

const d7p1data string = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`

func TestDay7Part1(t *testing.T) {
	result := day7.Part1(bufio.NewReader(bytes.NewReader([]byte(d7p1data))))
	if result != 6440 {
		t.Errorf("The expected result is 6440, got %v", result)
	}
}

func TestDay7Part2(t *testing.T) {
	result := day7.Part2(bufio.NewReader(bytes.NewReader([]byte(d7p1data))))
	if result != 5905 {
		t.Errorf("The expected result is 5905, got %v", result)
	}
}
