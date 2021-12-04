package aoc_2020

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jcm5155/advent-of-code/common"
)

// Day14 solution
func (h *Handler) Day14() (int64, int) {
	pzl := common.ReadInput("2020", "14").StringLines("\n")
	r := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)
	var mem = make(map[int][]byte)
	var mask string
	for _, row := range pzl {
		if row[:4] == "mask" {
			mask = strings.Split(row, " = ")[1]
		} else {
			parts := r.FindStringSubmatch(row)
			addr, _ := strconv.Atoi(parts[1])
			i, _ := strconv.Atoi(parts[2])
			val := []byte(fmt.Sprintf("%036b", i))

			for idx, maskVal := range mask {
				if maskVal != 'X' {
					val[idx] = byte(maskVal)
				}
			}
			mem[addr] = val
		}
	}
	var p1 int64
	for _, v := range mem {
		i, _ := strconv.ParseInt(string(v), 2, 64)
		p1 += i
	}

	// TODO: put p2 back in lel
	return p1, 0
}
