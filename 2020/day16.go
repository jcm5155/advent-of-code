package aoc_2020

import (
	"regexp"
	"strings"

	"github.com/jcm5155/advent-of-code/common"
)

// Day16 solution
func (h *Handler) Day16() (int, int) {
	p := common.ReadInput("2020", "16").StringLines("\n\n")
	paramRegex := regexp.MustCompile(`(.*?): (\d+)-(\d+) or (\d+)-(\d+)`)
	var bounds []int
	departures := [6][]int{}
	for idx, row := range strings.Split(p[0], "\n") {
		matches := paramRegex.FindStringSubmatch(row)
		toCheck := common.SliceAtoi(matches[2:])
		if idx == 0 {
			bounds = toCheck
		} else {
			for i := range bounds {
				bounds[i] = d16check(bounds[i], toCheck[i], i%2 == 0)
			}
		}
		if strings.Contains(matches[1], "departure") {
			for _, i := range toCheck {
				departures[idx] = append(departures[idx], i)
			}
		}
	}

	p1 := 0
	validTickets := [][]int{}
	for _, row := range strings.Split(p[2], "\n")[1:] {
		valid := true
		nums := common.SliceAtoi(strings.Split(row, ","))
		for _, n := range nums {
			if n < bounds[0] || n > bounds[3] {
				p1 += n
				valid = false
			}
		}
		if valid {
			validTickets = append(validTickets, nums)
		}
	}

	myTicket := common.SliceAtoi(strings.Split(strings.Split(p[1], "\n")[1], ","))
	validTickets = append(validTickets, myTicket)

	p2 := 1

	// 1439429522627
	return p1, p2
}

func d16combinedCheck(v int, d []int) bool {
	return (v >= d[0] && v <= d[1]) || (v >= d[2] && v <= d[3])
}

func d16check(prev, curr int, checkLower bool) int {
	if checkLower {
		if prev > curr {
			return curr
		}
	} else {
		if prev < curr {
			return curr
		}
	}
	return prev
}
