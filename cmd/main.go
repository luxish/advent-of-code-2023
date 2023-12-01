package main

import (
	"log"
	"os"

	"github.com/luxish/advent-of-code-2023/internal/day1"
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
	log.Printf("Day 1 part 2 - Result %v \n", day1.Day1Part2(utils.OpenFile(fileName)))
}
