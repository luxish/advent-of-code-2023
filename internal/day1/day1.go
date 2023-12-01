package day1

import (
	"bufio"
	"io"
	"strings"
	"unicode"
)

const INIT_VAL = -1

var DMAP = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"zero":  0,
}

func Day1Part1(reader io.Reader) int {
	scanner := bufio.NewScanner(bufio.NewReader(reader))
	result := 0
	for scanner.Scan() {
		first := INIT_VAL
		last := INIT_VAL
		line := scanner.Text()
		for idx := 0; idx < len(line); idx++ {
			first, last = firstAndLastForRune(line[idx], first, last)
		}
		if first != -1 && last != -1 {
			result += first*10 + last
		}
	}
	return result
}

func Day1Part2(reader io.Reader) int {
	scanner := bufio.NewScanner(bufio.NewReader(reader))
	result := 0
	for scanner.Scan() {
		first := INIT_VAL
		last := INIT_VAL
		nameBuf := make([]byte, 0)
		line := scanner.Text()
		for idx := 0; idx < len(line); idx++ {
			// Digit as string
			// The max name length for a digit is 5. Update this buffer with every index change and
			// check if there is a digit name as suffix.
			if idx >= 5 {
				nameBuf = nameBuf[1:]
			}
			nameBuf = append(nameBuf, line[idx])
			first, last = firstAndLastForBuffer(nameBuf, first, last)

			// Digit as character
			first, last = firstAndLastForRune(line[idx], first, last)

		}
		if first != INIT_VAL && last != INIT_VAL {
			result += first*10 + last
		}
	}
	return result
}

func firstAndLastForRune(r byte, first int, last int) (newFirst int, newLast int) {
	val := rune(r)
	if unicode.IsDigit(val) {
		var digitVal = int(val - '0')
		if first == INIT_VAL {
			first = digitVal
		}
		last = digitVal
	}
	return first, last
}

func firstAndLastForBuffer(buffer []byte, first int, last int) (newFirst int, newLast int) {
	for k, v := range DMAP {
		if strings.HasSuffix(string(buffer), k) {
			if first == INIT_VAL {
				first = v
			}
			last = v
		}
	}
	return first, last
}
