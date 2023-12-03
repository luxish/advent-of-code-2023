package day3

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"unicode"
)

// Number structure used to enacapsulate data about the number
type NumberMeta struct {
	numberName string
	line       int
	idx        int
}

func (n NumberMeta) ToNumber() int {
	val, _ := strconv.Atoi(n.numberName)
	return val
}

// Symbol structure use to store data abgout the position of a symbol
type Symbol struct {
	line int
	idx  int
}

// Engine map keep coordinates of all numbers and symbols.
type EngineMap struct {
	numbers   map[int][]NumberMeta
	symbols   []Symbol
	lineCount int
}

func (e *EngineMap) AddNumber(nm NumberMeta) {
	if e.numbers[nm.line] == nil {
		e.numbers[nm.line] = make([]NumberMeta, 0)
	}
	e.numbers[nm.line] = append(e.numbers[nm.line], nm)
}

func (e *EngineMap) AddSymbol(s Symbol) {
	e.symbols = append(e.symbols, s)
}

func NewEngineMap() *EngineMap {
	engine := &EngineMap{
		symbols: make([]Symbol, 0),
		numbers: make(map[int][]NumberMeta),
	}
	return engine
}

func EngineMapFromReader(reader io.Reader) *EngineMap {
	scanner := bufio.NewScanner(bufio.NewReader(reader))
	engine := NewEngineMap()
	lineIdx := 0
	for scanner.Scan() {
		line := scanner.Text()
		for idx := 0; idx < len(line); idx++ {
			numStr := ""
			for idx < len(line) && !IsFillerSymbol(line[idx]) && IsDigit(line[idx]) {
				numStr += string(line[idx])
				idx++

			}
			if numStr != "" {
				engine.AddNumber(NumberMeta{numStr, lineIdx, idx - len(numStr)})
			}
			if idx < len(line) && !IsFillerSymbol(line[idx]) && !IsDigit(line[idx]) {
				engine.AddSymbol(Symbol{lineIdx, idx})
			}
		}
		lineIdx++
	}
	engine.lineCount = lineIdx
	return engine
}

func IsFillerSymbol(data byte) bool {
	return string(data) == "."
}

func IsDigit(data byte) bool {
	return unicode.IsDigit(rune(data))
}

func Day3Part1(reader io.Reader) int {
	result := 0
	engine := EngineMapFromReader(reader)
	for _, s := range engine.symbols {
		// Top
		if s.line != 0 {
			for nidx := 0; nidx < len(engine.numbers[s.line-1]); nidx++ {
				nm := engine.numbers[s.line-1][nidx]
				if s.idx-nm.idx <= len(nm.numberName) && s.idx-nm.idx >= -1 {
					log.Printf("Found top of %v %v", s, nm.numberName)
					result += nm.ToNumber()
				}
			}
		}
		// Bottom
		if s.line != engine.lineCount {
			for nidx := 0; nidx < len(engine.numbers[s.line+1]); nidx++ {
				nm := engine.numbers[s.line+1][nidx]
				if s.idx-nm.idx <= len(nm.numberName) && s.idx-nm.idx >= -1 {
					log.Printf("Found bottom of %v %v", s, nm.numberName)
					result += nm.ToNumber()
				}
			}
		}
		// Sides
		for nidx := 0; nidx < len(engine.numbers[s.line]); nidx++ {
			nm := engine.numbers[s.line][nidx]
			if s.idx-nm.idx == len(nm.numberName) || s.idx-nm.idx == -1 {
				log.Printf("Found side of %v %v", s, nm.numberName)
				result += nm.ToNumber()
			}
		}
	}
	return result
}

func Day3Part2(reader io.Reader) int {
	result := 0
	engine := EngineMapFromReader(reader)
	for _, s := range engine.symbols {
		adjNumbers := make([]int, 0)
		// Top
		if s.line != 0 {
			for nidx := 0; nidx < len(engine.numbers[s.line-1]); nidx++ {
				nm := engine.numbers[s.line-1][nidx]
				if s.idx-nm.idx <= len(nm.numberName) && s.idx-nm.idx >= -1 {
					log.Printf("Found top of %v %v", s, nm.numberName)
					adjNumbers = append(adjNumbers, nm.ToNumber())
				}
			}
		}
		// Bottom
		if s.line != engine.lineCount {
			for nidx := 0; nidx < len(engine.numbers[s.line+1]); nidx++ {
				nm := engine.numbers[s.line+1][nidx]
				if s.idx-nm.idx <= len(nm.numberName) && s.idx-nm.idx >= -1 {
					log.Printf("Found bottom of %v %v", s, nm.numberName)
					adjNumbers = append(adjNumbers, nm.ToNumber())
				}
			}
		}
		// Sides
		for nidx := 0; nidx < len(engine.numbers[s.line]); nidx++ {
			nm := engine.numbers[s.line][nidx]
			if s.idx-nm.idx == len(nm.numberName) || s.idx-nm.idx == -1 {
				log.Printf("Found side of %v %v", s, nm.numberName)
				adjNumbers = append(adjNumbers, nm.ToNumber())
			}
		}
		if len(adjNumbers) == 2 {
			result += adjNumbers[0] * adjNumbers[1]
		}
	}
	return result
}
