package aoc_2023

import (
	aoc "github.com/jcm5155/advent-of-code/common"
	"strconv"
	"strings"
)

func (h *Handler) Day5() (int, int) {
	// this solution is extremely slow, but i ain't fixin it
	pzl := aoc.ReadInput("2023", "5").StringLines("\n\n")

	var seedsP1 []int
	for _, s := range strings.Fields(strings.Split(pzl[0], ": ")[1]) {
		seed, _ := strconv.Atoi(s)
		seedsP1 = append(seedsP1, seed)
	}

	var seedsP2 []int
	for i := 0; i < len(seedsP1)-1; i += 2 {
		for j := 0; j < seedsP1[i+1]; j++ {
			seedsP2 = append(seedsP2, seedsP1[i]+j)
		}
	}

	for _, chunk := range pzl[1:] {
		currentMap := d5_newMap(chunk)
		seedsP1 = d5_doTranslations(seedsP1, currentMap)
		seedsP2 = d5_doTranslations(seedsP2, currentMap)
	}

	p1 := seedsP1[0]
	for _, seed := range seedsP1[1:] {
		if seed < p1 {
			p1 = seed
		}
	}

	p2 := seedsP2[0]
	for _, seed := range seedsP2[1:] {
		if seed < p2 {
			p2 = seed
		}
	}

	return p1, p2
}

type d5_MapRow struct {
	Source      int
	Destination int
	Range       int
}

func (mr *d5_MapRow) Translate(seed int) (int, bool) {
	if mr.Source <= seed && mr.Source+mr.Range > seed {
		return mr.Destination + seed - mr.Source, true
	}
	return -1, false
}

func d5_newMap(chunk string) []d5_MapRow {
	var currentMap []d5_MapRow
	for _, row := range strings.Split(chunk, "\n")[1:] {
		allThree := strings.Fields(row)
		destination, _ := strconv.Atoi(allThree[0])
		source, _ := strconv.Atoi(allThree[1])
		rnge, _ := strconv.Atoi(allThree[2])

		currentMap = append(currentMap, d5_MapRow{
			Source:      source,
			Destination: destination,
			Range:       rnge,
		})
	}
	return currentMap
}

func d5_doTranslations(seeds []int, currentMap []d5_MapRow) []int {
	var isSeedMapped = make([]bool, len(seeds))
	for _, mr := range currentMap {
		for sidx, seed := range seeds {
			if isSeedMapped[sidx] {
				continue
			}

			if nextSeed, ok := mr.Translate(seed); ok {
				seeds[sidx] = nextSeed
				isSeedMapped[sidx] = true
			}
		}
	}

	return seeds
}
