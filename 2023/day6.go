package aoc_2023

import (
	"strconv"
	"strings"

	aoc "github.com/jcm5155/advent-of-code/common"
)

func (h *Handler) Day6() (int, int) {
	pzl := aoc.ReadInput("2023", "6").StringLines("\n")
	return d6p1(pzl), d6p2(pzl)
}

func d6p1(pzl []string) int {
	var times []int
	for _, t := range strings.Fields(pzl[0])[1:] {
		time, _ := strconv.Atoi(t)
		times = append(times, time)
	}

	var distances []int
	for _, d := range strings.Fields(pzl[1])[1:] {
		distance, _ := strconv.Atoi(d)
		distances = append(distances, distance)
	}

	var waysToWin = make([]int, len(times))
	for idx, t := range times {
		waysToWin[idx] = d6_solveRace(t, distances[idx])
	}

	p1 := 1
	for _, w := range waysToWin {
		p1 *= w
	}
	return p1
}

func d6p2(pzl []string) int {
	ts := strings.Join(strings.Fields(pzl[0])[1:], "")
	ds := strings.Join(strings.Fields(pzl[1])[1:], "")
	time, _ := strconv.Atoi(ts)
	recordDistance, _ := strconv.Atoi(ds)
	return d6_solveRace(time, recordDistance)
}

func d6_solveRace(time int, recordDistance int) int {
	var output int
	for i := 1; i < time; i++ {
		distanceTraveled := i * (time - i)
		if distanceTraveled > recordDistance {
			output++
		}
	}
	return output
}
