package aoc_2021

import (
	"github.com/jcm5155/advent-of-code/common"
)

var d3ctx *d3_Context

type d3_Context struct {
	PuzzleInput []uint
	BitSize     int
}

func (h *Handler) Day3() (uint, uint) {
	p := common.ReadInput("2021", "3").StringLines("\n")
	d3ctx = &d3_Context{
		PuzzleInput: common.SliceAtoUint(p),
		BitSize:     len(p[0]),
	}
	return d3p1(), d3p2()
}

func d3p1() uint {
	gamma, epsilon := d3_getGreeks(d3ctx.PuzzleInput)
	return gamma * epsilon
}

func d3p2() uint {
	pzl := d3ctx.PuzzleInput
	gammaFinal, _, _ := d3_recurseFilterRows(pzl, 0, true)
	epsilonFinal, _, _ := d3_recurseFilterRows(pzl, 0, false)
	return gammaFinal[0] * epsilonFinal[0]
}

func d3_getGreeks(rows []uint) (uint, uint) {
	bitSize := d3ctx.BitSize
	bitSetCounts := make([]int, bitSize)
	for _, r := range rows {
		for pos := bitSize; pos >= 0; pos-- {
			if common.IsBitSet(r, pos) {
				bitSetCounts[bitSize-pos-1]++
			}
		}
	}

	majorityThreshold := len(rows) / 2
	var gamma, epsilon uint
	for i, n := range bitSetCounts {
		if n > majorityThreshold {
			gamma = common.SetBit(gamma, bitSize-i-1)
		} else {
			epsilon = common.SetBit(epsilon, bitSize-i-1)
		}
	}
	return gamma, epsilon
}

func d3_recurseFilterRows(rows []uint, idx int, isUseGamma bool) ([]uint, int, bool) {
	bitSize := d3ctx.BitSize
	gamma, epsilon := d3_getGreeks(rows)
	var relevantGreek uint
	if isUseGamma {
		relevantGreek = gamma
	} else {
		relevantGreek = epsilon
	}

	criteriaPosition := bitSize - idx - 1
	isCriteriaBitSet := common.IsBitSet(relevantGreek, criteriaPosition)
	var meetsCriteria, failsCriteria []uint
	for _, row := range rows {
		if common.IsBitSet(row, criteriaPosition) == isCriteriaBitSet {
			meetsCriteria = append(meetsCriteria, row)
		} else {
			failsCriteria = append(failsCriteria, row)
		}
	}

	nextRows := meetsCriteria
	needTieBreak := len(meetsCriteria) == len(failsCriteria)
	if needTieBreak && isUseGamma != common.IsBitSet(meetsCriteria[0], criteriaPosition) {
		nextRows = failsCriteria
	}

	if len(nextRows) > 1 {
		return d3_recurseFilterRows(nextRows, idx+1, isUseGamma)
	}

	return nextRows, 0, isUseGamma
}
