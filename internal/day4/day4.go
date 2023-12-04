package day4

import (
	"bufio"
	"errors"
	"io"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/luxish/advent-of-code-2023/internal/utils"
)

type Game struct {
	Id         int
	WinningSet []int
	OwnSet     []int
}

func NewGameFromLine(line string) (*Game, error) {
	lineParts := strings.Split(line, ":")
	if len(lineParts) != 2 {
		return nil, errors.New("Invalid line " + line)
	}
	idParts := strings.Split(lineParts[0], " ")
	numberParts := strings.Split(lineParts[1], "|")
	if len(numberParts) != 2 {
		return nil, errors.New("Invalid line " + line)
	}
	toInt := func(s string) int {
		val, err := strconv.Atoi(strings.Trim(s, " "))
		if err != nil {
			return -1
		}
		return val
	}
	validNumbers := func(n int) bool {
		return n >= 0
	}
	wNum := utils.Map[string, int](strings.Split(strings.Trim(numberParts[0], " "), " "), toInt)
	oNum := utils.Map[string, int](strings.Split(strings.Trim(numberParts[1], " "), " "), toInt)
	return &Game{toInt(idParts[len(idParts)-1]), utils.Filer[int](wNum, validNumbers), utils.Filer[int](oNum, validNumbers)}, nil
}

func Day4Part1(reader io.Reader) int {
	result := 0
	scanner := bufio.NewScanner(bufio.NewReader(reader))
	for scanner.Scan() {
		line := scanner.Text()
		game, err := NewGameFromLine(line)
		if err != nil {
			log.Fatal(err)
		}
		counter := 0
		for oIdx := 0; oIdx < len(game.OwnSet); oIdx++ {
			for wIdx := 0; wIdx < len(game.WinningSet); wIdx++ {
				if game.OwnSet[oIdx] == game.WinningSet[wIdx] {
					counter++
				}
			}
		}
		if counter > 0 {
			result += int(math.Pow(2, float64(counter-1)))
		}
	}
	return result
}

func Day4Part2(reader io.Reader) int {
	result := 0
	gamesMap := make(map[int]int)
	scanner := bufio.NewScanner(bufio.NewReader(reader))
	for scanner.Scan() {
		line := scanner.Text()
		game, err := NewGameFromLine(line)
		if err != nil {
			log.Fatal(err)
		}
		counter := 0
		for oIdx := 0; oIdx < len(game.OwnSet); oIdx++ {
			for wIdx := 0; wIdx < len(game.WinningSet); wIdx++ {
				if game.OwnSet[oIdx] == game.WinningSet[wIdx] {
					counter++
				}
			}
		}
		// The original
		gamesMap[game.Id]++

		if counter > 0 {
			for gIdx := 0; gIdx < counter; gIdx++ {
				gamesMap[game.Id+gIdx+1] += gamesMap[game.Id]
			}
		}
	}
	for _, count := range gamesMap {
		result += count
	}

	return result
}
