package day2

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

const (
	MAX_RED   = 12
	MAX_GREEN = 13
	MAX_BLUE  = 14
)

type Game struct {
	Id    int
	Turns []map[string]int
}

func NewGame(data string) (*Game, error) {
	sections := strings.Split(data, ":")
	if len(sections) != 2 {
		return nil, errors.New(fmt.Sprintf("Invalid line %v", data))
	}
	// Parse header
	headerSections := strings.Split(sections[0], " ")
	if len(headerSections) != 2 {
		return nil, errors.New(fmt.Sprintf("Invalid head for line %v", data))
	}
	id, _ := strconv.Atoi(headerSections[1])
	// Parse content
	turns := make([]map[string]int, 0)
	gamesSections := strings.Split(sections[1], ";")
	for _, gameSection := range gamesSections {
		boxesMap := make(map[string]int)
		boxes := strings.Split(gameSection, ", ")
		for _, box := range boxes {
			entryData := strings.Split(strings.Trim(box, " "), " ")
			if len(entryData) != 2 {
				return nil, errors.New(fmt.Sprintf("Invalid entry %v in line %v", box, data))
			}
			val, _ := strconv.Atoi(entryData[0])
			boxesMap[entryData[1]] = val
		}
		turns = append(turns, boxesMap)
	}

	return &Game{Id: id, Turns: turns}, nil
}

func Day2Part1(reader io.Reader) int {
	scanner := bufio.NewScanner(bufio.NewReader(reader))
	result := 0
	for scanner.Scan() {
		game, err := NewGame(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		valid := true
		for _, turn := range game.Turns {
			if (turn["blue"] > MAX_BLUE) || (turn["green"] > MAX_GREEN) || (turn["red"] > MAX_RED) {
				valid = false
				break
			}
		}
		if valid {
			result += game.Id
		}
	}
	return result
}

func Day2Part2(reader io.Reader) int {
	scanner := bufio.NewScanner(bufio.NewReader(reader))
	result := 0
	for scanner.Scan() {
		game, err := NewGame(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		minBlue := 1
		minGreen := 1
		minRed := 1
		for _, turn := range game.Turns {
			if turn["blue"] > minBlue {
				minBlue = turn["blue"]
			}
			if turn["green"] > minGreen {
				minGreen = turn["green"]
			}
			if turn["red"] > minRed {
				minRed = turn["red"]
			}
		}
		result += minBlue * minGreen * minRed

	}
	return result
}
