package aoc_2020

import (
	"regexp"
	"strconv"

	"github.com/jcm5155/advent-of-code/common"
)

type d7bags map[string]map[string]int
type d7rbags map[string][]string

// Day7 solution
func (h *Handler) Day7() (int, int) {
	pzl := common.ReadInput("2020", "7").StringLines("\n")
	bags, rbags := d7bags{}, d7rbags{}
	parentRe := regexp.MustCompile(`(.+) bags contain`)
	childrenRe := regexp.MustCompile(`(\d+) (.+?) bag`)

	for _, row := range pzl {
		parent := parentRe.FindStringSubmatch(row)[1]
		children := childrenRe.FindAllStringSubmatch(row, -1)
		bags[parent] = map[string]int{}
		for _, c := range children {
			child := c[2]
			childNum, _ := strconv.Atoi(c[1])
			bags[parent][child] = childNum
			rbags[child] = append(rbags[child], parent)
		}
	}
	p1 := d7countParents(rbags, "shiny gold", map[string]int{})
	p2 := d7countChildren(bags, "shiny gold")
	return p1, p2
}

func d7countParents(m d7rbags, find string, om map[string]int) int {
	om[find] = 1
	for _, newTarget := range m[find] {
		d7countParents(m, newTarget, om)
	}
	return len(om) - 1
}

func d7countChildren(m d7bags, find string) int {
	total := 0
	for child, childNum := range m[find] {
		total += childNum * (d7countChildren(m, child) + 1)
	}
	return total
}
