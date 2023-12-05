package aoc_2023

import (
	"regexp"
	"strings"

	aoc "github.com/jcm5155/advent-of-code/common"
)

func (h *Handler) Day4() (int, int) {
	pzl := aoc.ReadInput("2023", "4").StringLines("\n")
	gameRegex := regexp.MustCompile(`: (.*) \| (.*)`)
	var p1, p2 int
	var ticketCounts = make([]int, len(pzl))

	for idx, row := range pzl {
		for _, m := range gameRegex.FindAllStringSubmatch(row, -1) {
			var score int

			// dumb set (empty struct takes 0 bytes)
			var winnerSet = make(map[string]struct{})
			for _, f := range strings.Fields(m[1]) {
				winnerSet[f] = struct{}{}
			}

			// do this first to account for initial ticket
			ticketCounts[idx] += 1

			addTicketOffset := 1
			for _, possibleWinner := range strings.Fields(m[2]) {
				if _, ok := winnerSet[possibleWinner]; ok {
					if score == 0 {
						score++
					} else {
						score *= 2
					}

					ticketCounts[idx+addTicketOffset] += ticketCounts[idx]
					addTicketOffset++
				}
			}
			p1 += score
		}
	}

	for _, n := range ticketCounts {
		p2 += n
	}

	return p1, p2
}
