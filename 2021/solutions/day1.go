package solutions

import (
	"github.com/jcm5155/advent-of-code/shared"
)

func (h *Handler) Day1() (int, int) {
	pzl := shared.ReadInput("2021", "1").IntLines("\n")
	var p1, p2 int

	p2bounds := len(pzl) - 3
	last := pzl[0]
	for i, current := range pzl {
		if current > last {
			p1++
		}
		if i < p2bounds && pzl[i+3] > current {
			p2++
		}
		last = current
	}
	return p1, p2
}
