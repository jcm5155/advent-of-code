package solutions

import (
	"image"
	"math"
	"strconv"

	"github.com/jcm5155/advent-of-code/shared"
)

// Day12 solution
func (h *Handler) Day12() (float64, float64) {
	pzl := shared.ReadInput("2020", "12").StringLines("\n")
	head, heading := map[int]byte{0: 'N', 90: 'E', 180: 'S', 270: 'W'}, 90
	s1, s2, wp := image.Point{}, image.Point{}, image.Point{10, 1}
	var c = make(chan image.Point, 1)
	for _, row := range pzl {
		mov := row[0]
		amt, _ := strconv.Atoi(row[1:])
		if mov == 'F' {
			c <- wp
			s2 = s2.Add(wp.Mul(amt))
			mov = head[heading]
		}
		switch mov {
		case 'L':
			heading = (heading - amt + 360) % 360
			for i := 0; i < amt/90; i++ {
				wp.X, wp.Y = -wp.Y, wp.X
			}
		case 'R':
			heading = (heading + amt + 360) % 360
			for i := 0; i < amt/90; i++ {
				wp.X, wp.Y = wp.Y, -wp.X
			}
		case 'N':
			s1.Y += amt
			wp.Y += amt
		case 'E':
			s1.X += amt
			wp.X += amt
		case 'S':
			s1.Y -= amt
			wp.Y -= amt
		case 'W':
			s1.X -= amt
			wp.X -= amt
		}
		select {
		case wp = <-c:
		default:
		}
	}
	return math.Abs(float64(s1.X)) + math.Abs(float64(s1.Y)),
		math.Abs(float64(s2.X)) + math.Abs(float64(s2.Y))
}
