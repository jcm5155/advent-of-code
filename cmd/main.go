package main

import (
	"os"
	"sync"

	aoc_2020 "github.com/jcm5155/advent-of-code/2020"
	aoc_2021 "github.com/jcm5155/advent-of-code/2021"
	"github.com/jcm5155/advent-of-code/common"
)

func main() {
	year := os.Args[1]
	days := os.Args[2:]

	var solver interface{}
	switch year {
	case "2021":
		solver = &aoc_2021.Handler{}
	case "2020":
		solver = &aoc_2020.Handler{}
	default:
		panic("unknown year: " + year)
	}

	var solutions = make([]*common.Solution, len(days))
	var wg sync.WaitGroup
	for idx, day := range days {
		wg.Add(1)
		go func(i int, d string) {
			defer wg.Done()
			solutions[i] = common.Solve(solver, d)
		}(idx, day)
	}
	wg.Wait()

	common.DisplaySolutions(year, solutions)
}
