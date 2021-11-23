package solutions

import (
	"strconv"

	"github.com/jcm5155/advent-of-code-2020/util"
)

// Day1 solution
func (h *Handler) Day1() (int, int) {
	p := util.ReadPuzzleInput("1", "\n")
	pzl := []int{}
	for _, i := range p {
		j, _ := strconv.Atoi(i)
		pzl = append(pzl, j)
	}

	p1, p2 := 0, 0
	p1Found, p2Found := false, false
	for i, ival := range pzl {
		for j := i + 1; j < len(pzl); j++ {
			switch c := ival + pzl[j]; {
			case c == 2020:
				p1 = ival * pzl[j]
				p1Found = true
			case c < 2020:
				for k := j + 1; k < len(pzl); k++ {
					if c+pzl[k] == 2020 {
						p2 = ival * pzl[j] * pzl[k]
						p2Found = true
						break
					}
				}
			default:
				break
			}
			if p1Found && p2Found {
				return p1, p2
			}
		}
	}
	panic("i swear this usually never happens!")
}
