package aoc_2021

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/jcm5155/advent-of-code/common"
)

func (h *Handler) Day8() (int, int) {
	pzl := common.ReadInput("2021", "8").StringLines("\n")
	var p1, p2 int
	for _, line := range pzl {
		parts := strings.Split(line, " | ")
		left, right := strings.Fields(parts[0]), strings.Fields(parts[1])
		digits := d8_solveLeft(left)
		// p1
		var buf bytes.Buffer
		for _, outputValue := range right {
			ss := newSegmentSet(outputValue)
			switch ss.Len() {
			case 2, 4, 3, 7:
				p1++
			}

			for i, digit := range digits {
				if digit.Equals(ss) {
					buf.WriteString(strconv.Itoa(i))
					break
				}
			}
		}
		p, _ := strconv.Atoi(buf.String())
		p2 += p
	}
	return p1, p2
}

func newSegmentSet(s string) *d8_segmentSet {
	var underlying = make(map[int32]struct{})
	for _, c := range s {
		underlying[c] = struct{}{}
	}
	return &d8_segmentSet{
		underlying: underlying,
	}
}

func d8_solveLeft(left []string) [10]*d8_segmentSet {
	var digits [10]*d8_segmentSet
	var unknowns []*d8_segmentSet
	for _, rosetta := range left {
		rSet := newSegmentSet(rosetta)
		switch len(rosetta) {
		case 2:
			digits[1] = rSet
		case 4:
			digits[4] = rSet
		case 3:
			digits[7] = rSet
		case 7:
			digits[8] = rSet
		case 6:
			// ensure 6 length values are processed first
			unknowns = append([]*d8_segmentSet{rSet}, unknowns...)
		case 5:
			unknowns = append(unknowns, rSet)
		}
	}

	for _, u := range unknowns {
		switch u.Len() {
		case 6:
			// 0, 6, 9
			// upper right segment check
			if digits[1].IsIn(u) {
				// 0, 9
				// middle segment check
				if digits[4].IsIn(u) {
					digits[9] = u
				} else {
					digits[0] = u
				}
			} else {
				digits[6] = u
			}
		case 5:
			// 2, 3, 5
			// upper right + bottom right check
			if digits[1].IsIn(u) {
				digits[3] = u
			} else {
				// 2, 5
				// count shared segments
				if digits[6].SharedSegmentCount(u) == 5 {
					digits[5] = u
				} else {
					digits[2] = u
				}
			}
		}
	}
	return digits
}

type d8_segmentSet struct {
	underlying map[int32]struct{}
}

func (c *d8_segmentSet) Len() int {
	return len(c.underlying)
}

func (c *d8_segmentSet) SharedSegmentCount(c2 *d8_segmentSet) int {
	var output int
	for key := range c.underlying {
		if _, ok := c2.underlying[key]; ok {
			output++
		}
	}
	return output
}

func (c *d8_segmentSet) IsIn(c2 *d8_segmentSet) bool {
	return c.Len() <= c2.Len() && c.SharedSegmentCount(c2) == c.Len()
}

func (c *d8_segmentSet) Equals(c2 *d8_segmentSet) bool {
	return c.Len() == c2.Len() && c.IsIn(c2)
}
