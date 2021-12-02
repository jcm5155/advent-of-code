package solutions

import (
	"github.com/jcm5155/advent-of-code/shared"
)

// Day3 solution
func (h *Handler) Day3() (int, int) {
	pzl := shared.ReadInput("2020", "3").StringLines("\n")
	return d3(pzl)
}

func d3(pzl []string) (int, int) {
	tree := byte('#')
	yMax, xMax := len(pzl), len(pzl[0])
	slopes, xPos, treeCounts := [5]int{1, 3, 5, 7, 1}, [5]int{}, [5]int{}

	for y := 0; y < yMax; y++ {
		for i, x := range xPos {
			if i == 4 && y%2 != 0 {
				continue
			}
			if pzl[y][x] == tree {
				treeCounts[i]++
			}
			xPos[i] = (x + slopes[i]) % xMax
		}
	}

	p1 := treeCounts[1]
	p2 := 1
	for _, i := range treeCounts {
		p2 *= i
	}
	return p1, p2
}
