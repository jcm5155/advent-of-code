package solutions

import (
	"strconv"

	"github.com/jcm5155/advent-of-code-2020/util"
)

type d8visited map[int]bool

// Day8 solution
func (h *Handler) Day8() (int, int) {
	pzl := util.ReadPuzzleInput("8", "\n")
	c := make(chan int, 1)
	p1 := d8resolveTimeline(pzl, c, 0, 0, d8visited{}, true, pzl[0])
	p2 := <-c
	return p1, p2
}

func d8resolveTimeline(pzl []string, c chan int, idx, acc int, visited d8visited, canBranch bool, override string) int {
	firstIteration := true
	for {
		if idx == len(pzl)-1 {
			c <- acc
			return 0
		}
		if visited[idx] {
			return acc
		}

		instruction := pzl[idx]
		if firstIteration {
			instruction = override
			firstIteration = false
		}
		cmd := instruction[:3]
		amt, _ := strconv.Atoi(instruction[4:])

		if cmd == "acc" {
			visited[idx] = true
			acc += amt
			idx++
		} else {
			idxRef := idx
			if cmd == "nop" {
				override = "jmp" + instruction[3:]
				idx++
			} else {
				override = "nop" + instruction[3:]
				idx += amt
			}
			if canBranch {
				visitedCopy := d8visited{}
				for k, v := range visited {
					visitedCopy[k] = v
				}
				go d8resolveTimeline(pzl, c, idxRef, acc, visitedCopy, false, override)
			}
			visited[idxRef] = true
		}

	}
}
