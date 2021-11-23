package solutions

import (
	"image"

	"github.com/jcm5155/advent-of-code-2020/util"
)

type d11seats map[image.Point]rune

// Day11 solution
func (h *Handler) Day11() (int, int) {
	pzl := util.ReadPuzzleInput("11", "\n")
	var seats = make(d11seats)
	for y := range pzl {
		for x := range pzl[0] {
			seats[image.Point{x, y}] = rune(pzl[y][x])
		}
	}
	return d11loop(seats, 1), d11loop(seats, 2)
}

func d11loop(seats d11seats, partNum int) int {
	coords := []image.Point{{0, 1}, {1, 1}, {0, -1}, {1, 0}, {-1, -1}, {-1, 0}, {-1, 1}, {1, -1}}
	occupied, empty := '#', 'L'
	scour, maxAdj, occ := d11p1logic, 3, 0
	if partNum == 2 {
		scour, maxAdj = d11p2logic, 4
	}
	for diff := true; diff; {
		occ, diff = 0, false

		next := d11seats{}
		for yx, seat := range seats {
			occAdj := 0
			for _, direction := range coords {
				if seats[scour(seats, yx, direction)] == occupied {
					occAdj++
				}
			}
			if seat == occupied && occAdj > maxAdj {
				seat = empty
			} else if seat == empty && occAdj == 0 || seat == occupied {
				seat = occupied
				occ++
			}
			next[yx] = seat
			diff = diff || seat != seats[yx]
		}
		seats = next
	}
	return occ
}

func d11p1logic(_ d11seats, yx, direction image.Point) image.Point {
	return yx.Add(direction)
}

func d11p2logic(seats d11seats, yx, direction image.Point) image.Point {
	for seats[yx.Add(direction)] == '.' {
		yx = yx.Add(direction)
	}
	return yx.Add(direction)
}
