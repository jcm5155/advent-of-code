package main

import (
	"os"
	"strconv"

	"github.com/jcm5155/advent-of-code-2020/solutions"
)

func main() {
	if len(os.Args) == 1 {
		for i := 1; i < 26; i++ {
			solutions.Day(i).Print()
		}
	} else {
		for _, i := range os.Args[1:] {
			j, err := strconv.Atoi(i)
			if err == nil {
				solutions.Day(j).Print()
			}
		}
	}
}
