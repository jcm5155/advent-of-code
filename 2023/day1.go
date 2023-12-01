package aoc_2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	aoc "github.com/jcm5155/advent-of-code/common"
)

func (h *Handler) Day1() (int, int) {
	pzl := aoc.ReadInput("2023", "1")
	p1 := d1p1(pzl)
	p2 := d1p2(pzl)
	return p1, p2
}

func d1p1(pzl *aoc.PuzzleInput) int {
	var output int

	for _, row := range pzl.StringLines("\n") {
		var first, last rune
		for _, r := range row {
			if unicode.IsDigit(r) {
				if first == 0 {
					first = r
					last = r
				} else {
					last = r
				}
			}
		}

		num, err := strconv.Atoi(fmt.Sprintf("%c%c", first, last))
		if err != nil {
			panic(err)
		}

		output += num
	}
	return output
}

func d1p2(pzl *aoc.PuzzleInput) int {
	var output int

	wordToDigit := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	// is there a better way to do this in Go?
	regexParts := make([]string, 0, len(wordToDigit)*2)
	for k, v := range wordToDigit {
		regexParts = append(regexParts, k)
		regexParts = append(regexParts, v)
	}
	pattern := regexp.MustCompile(strings.Join(regexParts, "|"))

	// lol
	overlapsLmao := strings.NewReplacer(
		"zerone", "zeroone",
		"oneight", "oneeight",
		"twone", "twoone",
		"threeight", "threeeight",
		"sevenine", "sevennine",
		"eightwo", "eighttwo",
		"eighthree", "eightthree",
		"nineight", "nineeight",
	)

	for _, r := range pzl.StringLines("\n") {
		row := overlapsLmao.Replace(r)
		matches := pattern.FindAllStringIndex(row, -1)

		firstMatch := matches[0]
		first := row[firstMatch[0]:firstMatch[1]]
		if len(first) != 1 {
			first = wordToDigit[first]
		}

		lastMatch := matches[len(matches)-1]
		last := row[lastMatch[0]:lastMatch[1]]
		if len(last) != 1 {
			last = wordToDigit[last]
		}

		num, err := strconv.Atoi(first + last)
		if err != nil {
			panic(err)
		}

		output += num
	}
	return output
}
