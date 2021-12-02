package solutions

import (
	"sort"
	"strconv"
	"strings"

	"github.com/jcm5155/advent-of-code/shared"
)

// Day5 solution
func (h *Handler) Day5() (int, int) {
	r := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")
	pzl := strings.Split(r.Replace(shared.ReadInput("2020", "5").String()), "\n")
	ids := []int{}
	for _, s := range pzl {
		sid, _ := strconv.ParseInt(s, 2, 16)
		ids = append(ids, int(sid))
	}
	sort.Ints(ids)
	for i, v := range ids {
		if ids[i+1]-v == 2 {
			return ids[len(ids)-1], v + 1
		}
	}
	panic("well this is awkward")
}
