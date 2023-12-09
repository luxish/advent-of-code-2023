package day7

import (
	"bufio"
	"io"
	"log"
	"sort"
	"strconv"
	"strings"
)

var CMAP = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

type CardsHand struct {
	Str    string
	Points int
	Value  int
}

func (ch CardsHand) compare(other CardsHand) int {
	val := ch.Value - other.Value
	if val == 0 {
		for idx := 0; idx < 5; idx++ {
			v1 := CMAP[string(ch.Str[idx])]
			v2 := CMAP[string(other.Str[idx])]
			if v1 != v2 {
				return v1 - v2
			}
		}
	}
	return val
}

func LastTwoOccurences1(str string) (int, int) {
	tokens := make(map[byte]int)
	for idx := 0; idx < len(str); idx++ {
		tokens[str[idx]]++

	}
	maxOc1 := 0
	maxOc2 := 0
	for _, val := range tokens {
		if maxOc1 < val {
			maxOc2 = maxOc1
			maxOc1 = val
		} else if maxOc2 < val {
			maxOc2 = val
		}
	}
	return maxOc2, maxOc1
}

func LastTwoOccurences2(str string) (int, int) {
	tokens := make(map[byte]int)
	for idx := 0; idx < len(str); idx++ {
		tokens[str[idx]]++

	}

	//Reset wild in tokens
	wild := tokens['J']
	tokens['J'] = 0
	maxOc1 := 0
	maxOc2 := 0
	for _, val := range tokens {
		if maxOc1 < val {
			maxOc2 = maxOc1
			maxOc1 = val
		} else if maxOc2 < val {
			maxOc2 = val
		}
	}
	return maxOc2, maxOc1 + wild
}

func HandValue(maxOc1, maxOc2 int) int {
	value := 0
	if maxOc1 == 2 && maxOc2 == 1 {
		value = 1 //One pair
	} else if maxOc1 == 2 && maxOc2 == 2 {
		value = 2 //Two pair
	} else if maxOc1 == 3 && maxOc2 == 1 {
		value = 3 //Three of a kind
	} else if maxOc1 == 3 && maxOc2 == 2 {
		value = 4 //Full
	} else if maxOc1 == 4 && maxOc2 == 1 {
		value = 5 //Four of a kind
	} else if maxOc1 == 5 && maxOc2 == 0 {
		value = 6 //Five of a kind
	}
	return value
}

func NewCardHandR1(str string) *CardsHand {
	parts := strings.Split(str, " ")
	if len(parts) != 2 {
		log.Fatal("Invalid card hand " + str)
	}
	points, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal("Invalid card hand " + str)
	}
	// Calculate value
	maxOc2, maxOc1 := LastTwoOccurences1(parts[0])
	return &CardsHand{Str: parts[0], Points: points, Value: HandValue(maxOc1, maxOc2)}
}

func NewCardHandR2(str string) *CardsHand {
	parts := strings.Split(str, " ")
	if len(parts) != 2 {
		log.Fatal("Invalid card hand " + str)
	}
	points, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal("Invalid card hand " + str)
	}
	// Calculate value
	maxOc2, maxOc1 := LastTwoOccurences2(parts[0])
	return &CardsHand{Str: parts[0], Points: points, Value: HandValue(maxOc1, maxOc2)}
}

func Part1(reader io.Reader) int {
	result := 0
	hands := make([]CardsHand, 0)
	scanner := bufio.NewScanner(bufio.NewReader(reader))
	for scanner.Scan() {
		ch := NewCardHandR1(scanner.Text())
		hands = append(hands, *ch)
	}
	sort.Slice(hands, func(idx1, idx2 int) bool {
		first := hands[idx1]
		second := hands[idx2]
		result := first.Value - second.Value
		if result == 0 {
			for idx := 0; idx < 5; idx++ {
				v1 := CMAP[string(first.Str[idx])]
				v2 := CMAP[string(second.Str[idx])]
				if v1 != v2 {
					result = v1 - v2
					break
				}
			}
		}
		return result < 0
	})
	for idx := 0; idx < len(hands); idx++ {
		result += (idx + 1) * hands[idx].Points
	}
	return result
}

func Part2(reader io.Reader) int {
	result := 0
	hands := make([]CardsHand, 0)
	scanner := bufio.NewScanner(bufio.NewReader(reader))
	for scanner.Scan() {
		ch := NewCardHandR2(scanner.Text())
		hands = append(hands, *ch)
	}
	sort.Slice(hands, func(idx1, idx2 int) bool {
		first := hands[idx1]
		second := hands[idx2]
		result := first.Value - second.Value
		if result == 0 {
			for idx := 0; idx < 5; idx++ {
				v1 := CMAP[string(first.Str[idx])]
				v2 := CMAP[string(second.Str[idx])]
				if v1 != v2 {
					if v1 == 11 {
						result = -1
						break
					}
					if v2 == 11 {
						result = 1
						break
					}
					result = v1 - v2
					break
				}
			}
		}
		return result < 0
	})
	for idx := 0; idx < len(hands); idx++ {
		result += (idx + 1) * hands[idx].Points
	}
	return result
}
