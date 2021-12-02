package solutions

import (
	"sort"

	"github.com/jcm5155/advent-of-code/shared"
)

// Day9 solution
func (h *Handler) Day9() (int, int) {
	pzl := shared.ReadInput("2020", "9").IntLines("\n")
	var c = make(chan int, 2)
	p1 := d9p1(pzl, c)

	for i := range pzl {
		select {
		case p2 := <-c:
			return p1, p2
		default:
			break
		}
		go d9p2(pzl, i, p1, c)
	}
	return p1, <-c
}

func d9p1(pzl []int, c chan int) int {
	for idx := 25; idx < len(pzl); idx++ {
		go d9checkSlice(pzl[idx-25:idx], pzl[idx], c)
		select {
		case p1 := <-c:
			return p1
		default:
			break
		}
	}
	return <-c
}

func d9checkSlice(checkSlice []int, find int, c chan int) {
	for i, val := range checkSlice {
		if val < find {
			for j := i + 1; j < len(checkSlice); j++ {
				if val+checkSlice[j] == find {
					return
				}
			}
		}
	}
	c <- find
}

func d9p2(pzl []int, start, find int, c chan int) {
	accum := 0
	for i := start; i < len(pzl); i++ {
		switch {
		case accum < find:
			accum += pzl[i]
		case accum > find:
			return
		case accum == find:
			goodSlice := pzl[start:i]
			if len(goodSlice) > 1 {
				sort.Ints(goodSlice)
				c <- goodSlice[0] + goodSlice[len(goodSlice)-1]
			}
			return
		}
	}
}
