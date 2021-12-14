package aoc_2021

import (
	"math"
	"sort"

	aoc "github.com/jcm5155/advent-of-code/common"
)

func (h *Handler) Day10() (int, int) {
	pzl := aoc.ReadInput("2021", "10").StringLines("\n")
	complementOf := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	scoreOf := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}

	var p1 int
	var p2 []int
outer:
	for _, row := range pzl {
		var stack = make(stackRune, 0, len(row))
		for _, currentBrace := range row {
			switch currentBrace {
			case ')', ']', '>', '}':
				if currentBrace != complementOf[stack.Pop()] {
					p1 += scoreOf[currentBrace]
					continue outer
				}
			default:
				stack.Push(currentBrace)
			}
		}
		var score int
		for popped := stack.Pop(); popped != math.MinInt32; popped = stack.Pop() {
			score = score*5 + scoreOf[popped]
		}
		p2 = append(p2, score)
	}
	sort.Ints(p2)
	return p1, p2[len(p2)/2]
}

type stackRune []rune

func (s *stackRune) Len() int {
	return len(*s)
}

func (s *stackRune) Push(r ...rune) {
	*s = append(*s, r...)
}

func (s *stackRune) Pop() rune {
	lenU := len(*s)
	if lenU == 0 {
		return math.MinInt32
	}
	o := (*s)[lenU-1]
	*s = (*s)[:lenU-1]
	return o
}
