package aoc_2023

import (
	"regexp"
	"strconv"

	aoc "github.com/jcm5155/advent-of-code/common"
)

func (h *Handler) Day2() (int, int) {
	pzl := aoc.ReadInput("2023", "2")
	var p1, p2 int

	cubeCountsRegex := regexp.MustCompile(` (\d+) (red|blue|green)`)

	for idx, row := range pzl.StringLines("\n") {
		var currentGame d2_CubeGame

		for _, m := range cubeCountsRegex.FindAllStringSubmatch(row, -1) {
			amt, _ := strconv.Atoi(m[1])
			color := m[2]

			switch color {
			case "red":
				if amt > currentGame.Red {
					currentGame.Red = amt
				}
			case "green":
				if amt > currentGame.Green {
					currentGame.Green = amt
				}
			case "blue":
				if amt > currentGame.Blue {
					currentGame.Blue = amt
				}
			default:
				panic("unknown color " + color)
			}
		}

		if currentGame.IsValid(12, 13, 14) {
			p1 += idx + 1
		}

		p2 += currentGame.Power()
	}

	return p1, p2
}

type d2_CubeGame struct {
	Red   int
	Green int
	Blue  int
}

func (cg *d2_CubeGame) IsValid(redLimit, greenLimit, blueLimit int) bool {
	return cg.Red <= redLimit && cg.Green <= greenLimit && cg.Blue <= blueLimit
}

func (cg *d2_CubeGame) Power() int {
	return cg.Red * cg.Green * cg.Blue
}
