package day6

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

func ExtractNumbersString(str string) string {
	parts := strings.Split(str, ":")
	if len(parts) != 2 {
		log.Fatal("Invalid timings for " + str)
	}
	return parts[1]
}

func ExtractInts(str string) []int {
	nStr := ExtractNumbersString(str)
	result := make([]int, 0)
	tokensStr := strings.Split(nStr, " ")
	for _, token := range tokensStr {
		t, err := strconv.Atoi(token)
		if err == nil {
			result = append(result, t)
		}
	}
	return result
}

func Part1(reader io.Reader) int {
	result := 1
	scanner := bufio.NewScanner(bufio.NewReader(reader))
	scanner.Scan()
	timings := ExtractInts(scanner.Text())
	scanner.Scan()
	distances := ExtractInts(scanner.Text())
	for idx := 0; idx < len(timings); idx++ {
		time := timings[idx]
		dist := distances[idx]
		wins := 0
		for speed := 1; speed < time; speed++ {
			if dist/speed < time-speed {
				wins++
			}
		}
		result *= wins
	}

	return result
}

func Part2(reader io.Reader) int {
	scanner := bufio.NewScanner(bufio.NewReader(reader))
	scanner.Scan()
	timingStr := strings.ReplaceAll(ExtractNumbersString(scanner.Text()), " ", "")
	time, err := strconv.Atoi(timingStr)
	if err != nil {
		log.Fatal("Invalid timing " + timingStr)
	}
	scanner.Scan()
	distStr := strings.ReplaceAll(ExtractNumbersString(scanner.Text()), " ", "")
	dist, err := strconv.Atoi(distStr)
	if err != nil {
		log.Fatal("Invalid timing " + distStr)
	}
	log.Print(time, dist)
	wins := 0
	for speed := 1; speed < time; speed++ {
		if dist/speed < time-speed {
			wins++
		}
	}
	return wins
}
