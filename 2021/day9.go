package aoc_2021

import (
	"sort"
	"strconv"

	aoc "github.com/jcm5155/advent-of-code/common"
)

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

var (
	d9_points [][]int
	d9_lenY   int
	d9_lenX   int
)

func (h *Handler) Day9() (int, int) {
	p1, p2, pzl := 0, 0, aoc.ReadInput("2021", "9").StringLines("\n")
	d9_lenY, d9_lenX = len(pzl), len(pzl[0])
	d9_points = make([][]int, d9_lenY, d9_lenY)
	for y, row := range pzl {
		d9_points[y] = make([]int, d9_lenX, d9_lenX)
		for x, c := range []rune(row) {
			point, _ := strconv.Atoi(string(c))
			d9_points[y][x] = point
		}
	}
	var basinSizes = make(chan int, d9_lenY)
	var basinCount int
	for y, row := range d9_points {
		for x, point := range row {
			up, right, down, left := d9_areAdjacentsHigher(y, x)
			if !((up || y == 0) &&
				(right || x == d9_lenX-1) &&
				(down || y == d9_lenY-1) &&
				(left || x == 0)) {
				continue
			}
			// basin found
			p1 += point
			basinCount++
			go func(y_, x_ int) {
				var bsn8r = make(d9_basinator)
				bsn8r.RecurseVisit(y_, x_, -1)
				basinSizes <- len(bsn8r)
			}(y, x)
		}
	}
	var basins = make([]int, basinCount, basinCount)
	for i := range basins {
		basins[i] = <-basinSizes
	}
	sort.Ints(basins)
	p2 = basins[basinCount-1] * basins[basinCount-2] * basins[basinCount-3]
	return p1 + basinCount, p2
}

func d9_areAdjacentsHigher(y, x int) (bool, bool, bool, bool) {
	current := d9_points[y][x]
	up := y > 0 && current < d9_points[y-1][x]
	right := x < d9_lenX-1 && current < d9_points[y][x+1]
	down := y < d9_lenY-1 && current < d9_points[y+1][x]
	left := x > 0 && current < d9_points[y][x-1]
	return up, right, down, left
}

type d9_basinator map[int]struct{}

func (b *d9_basinator) RecurseVisit(y, x, fromDirection int) {
	point := d9_points[y][x]
	if point == 9 {
		return
	}
	b.Visit(y, x)
	up, right, down, left := d9_areAdjacentsHigher(y, x)
	if fromDirection != UP && up {
		b.RecurseVisit(y-1, x, DOWN)
	}
	if fromDirection != RIGHT && right {
		b.RecurseVisit(y, x+1, LEFT)
	}
	if fromDirection != DOWN && down {
		b.RecurseVisit(y+1, x, UP)
	}
	if fromDirection != LEFT && left {
		b.RecurseVisit(y, x-1, RIGHT)
	}
}

func (b d9_basinator) Visit(y, x int) {
	// https://en.wikipedia.org/wiki/Pairing_function#Cantor_pairing_function
	id := (y+x)*(y+x+1)/2 + y
	if _, ok := b[id]; !ok {
		b[id] = struct{}{}
	}
}
