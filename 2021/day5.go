package aoc_2021

import (
	"regexp"
	"strconv"

	aoc "github.com/jcm5155/advent-of-code/common"
)

func (h *Handler) Day5() (int, int) {
	pzl := aoc.ReadInput("2021", "5").StringLines("\n")
	var grid = map[int]map[int]int{}
	var intersections int

	// plot a point on the grid
	plot := func(x, y int) {
		if intersects, ok := grid[x]; !ok {
			grid[x] = make(map[int]int)
		} else if intersects[y] == 1 {
			intersections++
		}
		grid[x][y]++
	}

	var diagonals [][4]int
	var inputRowRegex = regexp.MustCompile(`(?P<x1>\d+),(?P<y1>\d+) -> (?P<x2>\d+),(?P<y2>\d+)`)
	for _, row := range pzl {
		coords := aoc.RegexpNamedMatches(inputRowRegex, row) // extracts regexp capturing group values to a map[string]string
		x1, _ := strconv.Atoi(coords["x1"])
		y1, _ := strconv.Atoi(coords["y1"])
		x2, _ := strconv.Atoi(coords["x2"])
		y2, _ := strconv.Atoi(coords["y2"])

		switch {
		// vertical lines
		case x1 == x2:
			higherValue, lowerValue := aoc.MaxInt(y1, y2), aoc.MinInt(y1, y2)
			for y := lowerValue; y <= higherValue; y++ {
				plot(x1, y)
			}
		// horizontal lines
		case y1 == y2:
			higherValue, lowerValue := aoc.MaxInt(x1, x2), aoc.MinInt(x1, x2)
			for x := lowerValue; x <= higherValue; x++ {
				plot(x, y1)
			}
		// diagonal lines
		default:
			diagonals = append(diagonals, [4]int{x1, y1, x2, y2})
		}
	}

	// cache number of intersections before plotting diagonals
	p1 := intersections

	for _, d := range diagonals {
		x1, y1, x2, y2 := d[0], d[1], d[2], d[3]
		xdir, ydir := 1, 1
		if x1 > x2 {
			xdir = -1
		}
		if y1 > y2 {
			ydir = -1
		}
		x, y := x1, y1
		for x != x2+xdir {
			plot(x, y)
			x += xdir
			y += ydir
		}
	}
	return p1, intersections
}
