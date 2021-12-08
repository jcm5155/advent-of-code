package aoc_2021

import (
	aoc "github.com/jcm5155/advent-of-code/common"
)

func (h *Handler) Day6() (int, int) {
	pzl := aoc.ReadInput("2021", "6").IntLines(",")
	var p1, p2 int

	// truths:
	//  - fish spawn exactly 1 new fish every 7 days (their "spawnday").
	//  - fish who share a birthday also have the same spawnday.
	//  - on any given day, all new fish share their spawnday with the fish
	//    that are currently due to spawn a new fish in 7 days.

	var fishBySpawnday = make([]int, 9)
	for _, fish := range pzl {
		fishBySpawnday[fish]++
	}
	for d := 0; d < 256; d++ {
		todaysSpawnGroup := d % 9
		spawnGroupInSevenDays := (d + 7) % 9
		// add today's fish to the fish in +7 days
		fishBySpawnday[spawnGroupInSevenDays] += fishBySpawnday[todaysSpawnGroup]
		if d == 79 {
			p1 = aoc.SliceSum(fishBySpawnday)
		}
	}
	p2 = aoc.SliceSum(fishBySpawnday)
	return p1, p2
}
