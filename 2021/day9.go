package aoc_2021

import (
	"strconv"

	"github.com/jcm5155/advent-of-code/common"
)

func (h *Handler) Day9() (int, int) {
	var p1, p2 int
	pzl := common.ReadInput("2021", "9").StringLines("\n")
	for y, row := range pzl {
		for x, c := range []rune(row){
			pivotValue, _ := strconv.Atoi(string(c))

			if y > 0 {
				up, _ := strconv.Atoi(string(pzl[y-1][x]))
				if pivotValue >= up {
					continue
				}
			}

			if y < len(pzl)-1 {
				down, _ := strconv.Atoi(string(pzl[y+1][x]))
				if pivotValue >= down {
					continue
				}
			}

			if x < len(row)-1 {
				right, _ := strconv.Atoi(string(pzl[y][x+1]))
				if pivotValue >= right {
					continue
				}
			}

			if x > 0 {
				left, _ := strconv.Atoi(string(pzl[y][x-1]))
				if pivotValue >= left {
					continue
				}
			}

			p1 += pivotValue + 1
		}
	}
	return p1, p2
}