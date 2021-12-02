package main

import (
	"fmt"
	"os"
	"strconv"

	aoc_2020 "github.com/jcm5155/advent-of-code/2020/solutions"
	aoc_2021 "github.com/jcm5155/advent-of-code/2021/solutions"
	"github.com/jcm5155/advent-of-code/shared"
)

func main() {
	year := os.Args[1]

	var handler interface{}
	switch year {
	case "2021":
		handler = &aoc_2021.Handler{}
	case "2020":
		handler = &aoc_2020.Handler{}
	default:
		panic("unknown year: " + year)
	}

	for _, i := range os.Args[2:] {
		if j, err := strconv.Atoi(i); err != nil {
			fmt.Printf("Skipping invalid day: %v...\n", j)
		} else {
			shared.Solve(j, handler).Print()
		}
	}
}
