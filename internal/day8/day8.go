package day8

import (
	"bufio"
	"io"
	"log"
	"regexp"
	"strings"

	"github.com/luxish/advent-of-code-2023/internal/utils"
)

type Map struct {
	Input string
	Paths map[string][2]string
}

func (m *Map) RegisterPath(main string, left string, right string) {
	if m.Paths == nil {
		m.Paths = make(map[string][2]string)
	}
	m.Paths[main] = [2]string{left, right}
}

func ExtractMap(scanner *bufio.Scanner) *Map {
	m := Map{}
	// Input
	scanner.Scan()
	m.Input = scanner.Text()
	// Paths
	rexp, err := regexp.Compile("[0-9A-Z][0-9A-Z][0-9A-Z]")
	if err != nil {
		log.Fatal(err)
	}
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		labels := rexp.FindAllString(line, 3)
		m.RegisterPath(labels[0], labels[1], labels[2])
	}
	return &m
}

func Part1(reader io.Reader) int {
	result := 0
	scanner := bufio.NewScanner(bufio.NewReader(reader))
	m := ExtractMap(scanner)
	currentLabel := "AAA"
	destLabel := "ZZZ"
	for idx := 0; idx < len(m.Input) && currentLabel != destLabel; idx = (idx + 1) % len(m.Input) {
		action := string(m.Input[idx])
		switch action {
		case "R":
			currentLabel = m.Paths[currentLabel][1]
		case "L":
			currentLabel = m.Paths[currentLabel][0]
		}
		result++
	}
	return result
}

func Part2(reader io.Reader) int {
	result := 0
	scanner := bufio.NewScanner(bufio.NewReader(reader))
	m := ExtractMap(scanner)

	startLabels := make([]string, 0)
	for label := range m.Paths {
		if strings.HasSuffix(label, "A") {
			startLabels = append(startLabels, label)
		}
	}

	for _, label := range startLabels {
		movesToEnd := 0
		currentLabel := label
		for idx := 0; idx < len(m.Input); idx = (idx + 1) % len(m.Input) {
			action := string(m.Input[idx])
			switch action {
			case "L":
				currentLabel = m.Paths[currentLabel][0]
			case "R":
				currentLabel = m.Paths[currentLabel][1]
			}
			movesToEnd++
			if strings.HasSuffix(currentLabel, "Z") {
				break
			}
		}
		// Least common multiple
		if result == 0 {
			result = movesToEnd
		} else {
			result = result * movesToEnd / utils.GCD(result, movesToEnd)
		}

	}

	return result
}
