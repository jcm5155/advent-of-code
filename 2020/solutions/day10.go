package solutions

import (
	"sort"

	"github.com/jcm5155/advent-of-code-2020/util"
)

// Day10 solution
func (h *Handler) Day10() (int, int) {
	pzl := util.ReadIntPuzzleInput("10")
	sort.Ints(pzl)
	target := pzl[len(pzl)-1] + 3
	pzl = append(pzl, target)
	jolt, jolts, waysTo := 0, [4]int{}, map[int]int{0: 1}
	for _, v := range pzl {
		waysTo[v] = waysTo[v-1] + waysTo[v-2] + waysTo[v-3]
		jolts[v-jolt]++
		jolt = v
	}
	return jolts[1] * jolts[3], waysTo[target]
}
