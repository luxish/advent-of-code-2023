package main

import (
	"log"
	"os"

	"github.com/luxish/advent-of-code-2023/internal/day8"
	"github.com/luxish/advent-of-code-2023/internal/utils"
)

func main() {
	log.Println("Advent of code")
	if len(os.Args) < 2 {
		log.Fatal("Provide a file as input")
	}
	fileName := os.Args[1]

	log.Println("Using file " + fileName)
	// log.Printf("Day 1 part 1 - Result %v \n", day1.Day1Part1(utils.OpenFile(fileName)))
	// log.Printf("Day 1 part 2 - Result %v \n", day1.Day1Part2(utils.OpenFile(fileName)))
	// log.Printf("Day 2 part 1 - Result %v \n", day2.Day2Part1(utils.OpenFile(fileName)))
	// log.Printf("Day 2 part 2 - Result %v \n", day2.Day2Part2(utils.OpenFile(fileName)))
	// log.Printf("Day 3 part 1 - Result %v \n", day3.Day3Part1(utils.OpenFile(fileName)))
	// log.Printf("Day 3 part 2 - Result %v \n", day3.Day3Part2(utils.OpenFile(fileName)))
	// log.Printf("Day 4 part 1 - Result %v \n", day4.Day4Part1(utils.OpenFile(fileName)))
	// log.Printf("Day 4 part 2 - Result %v \n", day4.Day4Part2(utils.OpenFile(fileName)))
	// log.Printf("Day 5 part 1 - Result %v \n", day5.Part1(utils.OpenFile(fileName)))
	// log.Printf("Day 5 part 2 - Result %v \n", day5.Part2(utils.OpenFile(fileName)))
	// log.Printf("Day 6 part 1 - Result %v \n", day6.Part1(utils.OpenFile(fileName)))
	// log.Printf("Day 6 part 2 - Result %v \n", day6.Part2(utils.OpenFile(fileName)))
	// log.Printf("Day 7 part 1 - Result %v \n", day7.Part1(utils.OpenFile(fileName)))
	// log.Printf("Day 7 part 2 - Result %v \n", day7.Part2(utils.OpenFile(fileName)))
	// log.Printf("Day 8 part 1 - Result %v \n", day8.Part1(utils.OpenFile(fileName)))
	log.Printf("Day 8 part 2 - Result %v \n", day8.Part2(utils.OpenFile(fileName)))
}
