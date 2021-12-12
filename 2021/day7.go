package aoc_2021

import (
	"math"
	"strconv"

	aoc "github.com/jcm5155/advent-of-code/common"
)

func (h *Handler) Day7() (int, int) {
	pzl := aoc.ReadInput("2021", "7").StringLines(",")
	nums := make(map[int]int)
	for _, n := range pzl {
		num, _ := strconv.Atoi(n)
		nums[num]++
	}
	crabs := &d7_crabs{underlying: nums}
	return crabs.Solve()
}

type d7_crabs struct {
	underlying map[int]int
}

func (d *d7_crabs) Solve() (int, int) {
	constantBurn, nonConstantBurn := -1, -1
	var max int
	for i := range d.underlying {
		if i > max {
			max = i
		}
	}
	for position := 0; position <= max; position++ {
		curFuel := d.p1Burn(position)
		if constantBurn == -1 || curFuel < constantBurn {
			constantBurn = curFuel
		}
		curFuel = d.p2Burn(position)
		if nonConstantBurn == -1 || curFuel < nonConstantBurn {
			nonConstantBurn = curFuel
		}
	}
	return constantBurn, nonConstantBurn
}

func (d *d7_crabs) p1Burn(loc int) int {
	var fuel int
	for i, amount := range d.underlying {
		fuel += int(math.Abs(float64(loc-i))) * amount
	}
	return fuel
}

func (d *d7_crabs) p2Burn(loc int) int {
	var fuel int
	for position, amount := range d.underlying {
		diff := int(math.Abs(float64(loc - position)))
		fuel += (diff * (diff + 1)) / 2 * amount
	}
	return fuel
}
