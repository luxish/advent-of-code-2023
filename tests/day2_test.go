package tests

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/luxish/advent-of-code-2023/internal/day2"
)

const d2p1data string = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`

func TestDay2Part1(t *testing.T) {
	result := day2.Day2Part1(bufio.NewReader(bytes.NewReader([]byte(d2p1data))))
	if result != 8 {
		t.Errorf("The expected result is 8, got %v", result)
	}
}

func TestDay2Part2(t *testing.T) {
	result := day2.Day2Part2(bufio.NewReader(bytes.NewReader([]byte(d2p1data))))
	if result != 2286 {
		t.Errorf("The expected result is 2286, got %v", result)
	}
}
