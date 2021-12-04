package aoc_2021

import (
	"strconv"
	"strings"

	"github.com/jcm5155/advent-of-code/common"
)

func (h *Handler) Day2() (int, int) {
	pzl := common.ReadInput("2021", "2").StringLines("\n")
	var p1H, p1D, p2H, p2D, aim int
	for _, row := range pzl {
		fields := strings.Fields(row)
		direction := fields[0]
		distance, _ := strconv.Atoi(fields[1])

		switch direction {
		case "forward":
			p1H += distance
			p2H += distance
			p2D += aim * distance
		case "down":
			p1D += distance
			aim += distance
		case "up":
			p1D -= distance
			aim -= distance
		default:
			panic("unknown command")
		}
	}
	return p1H * p1D, p2H * p2D
}
