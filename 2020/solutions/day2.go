package solutions

import (
	"strconv"
	"strings"

	"github.com/jcm5155/advent-of-code/shared"
)

// Day2 solution
func (h *Handler) Day2() (int, int) {
	pzl := shared.ReadInput("2020", "2").StringLines("\n")
	p1, p2 := 0, 0
	for _, i := range pzl {
		minVal, maxVal, val, pass := d2Process(i)
		valCount := strings.Count(pass, val)
		if minVal <= valCount && maxVal >= valCount {
			p1++
		}
		minVal--
		maxVal--
		exactlyOne := false
		if minVal <= len(pass) && pass[minVal] == val[0] {
			exactlyOne = !exactlyOne
		}
		if maxVal <= len(pass) && pass[maxVal] == val[0] {
			exactlyOne = !exactlyOne
		}
		if exactlyOne == true {
			p2++
		}

	}
	return p1, p2
}

func d2Process(i string) (int, int, string, string) {
	parts := strings.Fields(i)
	lims := strings.Split(parts[0], "-")
	minVal, _ := strconv.Atoi(lims[0])
	maxVal, _ := strconv.Atoi(lims[1])
	val := parts[1][:1]
	pass := parts[2]
	return minVal, maxVal, val, pass
}
