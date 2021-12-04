package aoc_2020

import (
	"github.com/jcm5155/advent-of-code/common"
)

// Day6 solution
func (h *Handler) Day6() (int, int) {
	var nl = rune('\n')
	pzl := common.ReadInput("2020", "6").StringLines("\n\n")
	p1, p2 := 0, 0
	for _, group := range pzl {
		numPeople := 1
		m := make(map[rune]int)
		for _, c := range group {
			if c == nl {
				numPeople++
			} else {
				m[c]++
			}
		}
		for _, v := range m {
			p1++
			if v == numPeople {
				p2++
			}
		}
	}
	return p1, p2
}
