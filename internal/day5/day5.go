package day5

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/luxish/advent-of-code-2023/internal/utils"
)

type Record struct {
	Dest   int
	Source int
	Range  int
}

func ToInt(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func ExtractSeeds(str string) []int {
	seedsStr := strings.Split(str, ":")
	if len(seedsStr) != 2 {
		log.Fatal("Invalid seeds string")
	}
	seeds := strings.Split(strings.Trim(seedsStr[1], " "), " ")
	return utils.Map[string, int](seeds, ToInt)
}

func ExtractRecord(str string) *Record {
	parts := strings.Split(str, " ")
	if len(parts) != 3 {
		log.Fatal("Invalid record " + str)
	}
	partsInt := utils.Map[string, int](parts, ToInt)
	return &Record{partsInt[0], partsInt[1], partsInt[2]}
}

func ExtractSection(scanner *bufio.Scanner) []Record {
	sectionRecords := make([]Record, 0)
	for scanner.Scan() {
		sectionLine := scanner.Text()
		if sectionLine == "" {
			break
		}
		record := ExtractRecord(sectionLine)
		sectionRecords = append(sectionRecords, *record)
	}
	return sectionRecords
}

func SeedToLocation(seed int, sections [][]Record) int {
	source := seed
	for idx := 0; idx < len(sections); idx++ {
		section := sections[idx]
		var rec *Record = nil
		for _, record := range section {
			if record.Source <= source && source < record.Source+record.Range {
				rec = &record
				break
			}
		}
		if rec != nil {
			source = rec.Dest + source - rec.Source
		}
	}
	return source
}

func Part1(reader io.Reader) int {
	result := -1
	sections := make([][]Record, 0)
	scanner := bufio.NewScanner(bufio.NewReader(reader))
	scanner.Scan()
	sources := ExtractSeeds(scanner.Text())
	scanner.Scan()
	for scanner.Scan() {
		sections = append(sections, ExtractSection(scanner))
	}
	for idxSource := 0; idxSource < len(sources); idxSource++ {
		loc := SeedToLocation(sources[idxSource], sections)
		if result == -1 || result > loc {
			result = loc
		}
	}
	return result
}

func Part2(reader io.Reader) int {
	result := -1
	sections := make([][]Record, 0)
	scanner := bufio.NewScanner(bufio.NewReader(reader))
	scanner.Scan()
	sources := ExtractSeeds(scanner.Text())
	scanner.Scan()
	for scanner.Scan() {
		sections = append(sections, ExtractSection(scanner))
	}
	for idxSource := 0; idxSource < len(sources); idxSource += 2 {
		for idxR := 0; idxR < sources[idxSource+1]; idxR++ {
			loc := SeedToLocation(sources[idxSource]+idxR, sections)
			if result == -1 || result > loc {
				result = loc
			}
		}
	}
	return result
}
