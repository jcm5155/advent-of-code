package aoc_2023

import (
	"regexp"
	"strconv"

	aoc "github.com/jcm5155/advent-of-code/common"
)

type d3_EnginePart struct {
	Value  int
	Symbol *d3_Symbol
}

type d3_Symbol struct {
	Value     string
	X         int
	Y         int
	Adjacents []*d3_EnginePart
}

func (h *Handler) Day3() (int, int) {
	pzl := aoc.ReadInput("2023", "3").StringLines("\n")

	enginePartRegex := regexp.MustCompile(`\d+`)
	symbolRegex := regexp.MustCompile(`[^\d\.]{1}`)
	minY, maxY := 0, len(pzl)
	minX, maxX := 0, len(pzl[0])
	var partMap = make([][]*d3_EnginePart, maxY)
	var partList []*d3_EnginePart
	var symbols []*d3_Symbol

	for y, row := range pzl {
		partMap[y] = make([]*d3_EnginePart, maxX)

		// build parts list and "map"
		for _, m := range enginePartRegex.FindAllStringSubmatchIndex(row, -1) {
			start, stop := m[0], m[1]
			val, _ := strconv.Atoi(row[start:stop])
			foundPart := &d3_EnginePart{Value: val}

			// fill in the same reference to foundPart in all slots that its number occupies
			for x := start; x < stop; x++ {
				partMap[y][x] = foundPart
			}

			partList = append(partList, foundPart)
		}

		// build symbols list
		for _, m := range symbolRegex.FindAllStringSubmatchIndex(row, -1) {
			start, stop := m[0], m[1]
			symbols = append(symbols, &d3_Symbol{
				Value:     row[start:stop],
				X:         start,
				Y:         y,
				Adjacents: []*d3_EnginePart{},
			})
		}
	}

	// check 1 space around each symbol
	for _, sym := range symbols {
		for y := -1; y < 2; y++ {
			for x := -1; x < 2; x++ {
				netX := x + sym.X
				netY := y + sym.Y

				// boundary checking (and skip urself)
				if (x == 0 && y == 0) || netX < minX || netX > maxX || netY < minY || netY > maxY {
					continue
				}

				// if a part is found, handle it
				if part := partMap[netY][netX]; part != nil {
					part.Symbol = sym

					var isAlreadyAdjacent bool
					for _, adj := range sym.Adjacents {
						if part == adj {
							isAlreadyAdjacent = true
							break
						}
					}
					if !isAlreadyAdjacent {
						sym.Adjacents = append(sym.Adjacents, part)
					}

				}
			}
		}
	}

	// calculate scores
	var p1, p2 int

	for _, part := range partList {
		if part.Symbol != nil {
			p1 += part.Value
		}
	}

	for _, symbol := range symbols {
		if symbol.Value == "*" && len(symbol.Adjacents) == 2 {
			p2 += symbol.Adjacents[0].Value * symbol.Adjacents[1].Value
		}
	}

	return p1, p2
}
