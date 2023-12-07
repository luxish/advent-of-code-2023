package tests

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/luxish/advent-of-code-2023/internal/day5"
)

const d5p1data string = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`

func TestDay5Part1(t *testing.T) {
	result := day5.Part1(bufio.NewReader(bytes.NewReader([]byte(d5p1data))))
	if result != 35 {
		t.Errorf("The expected result is 35, got %v", result)
	}
}

func TestDay5Part2(t *testing.T) {
	result := day5.Part2(bufio.NewReader(bytes.NewReader([]byte(d5p1data))))
	if result != 46 {
		t.Errorf("The expected result is 46, got %v", result)
	}
}
