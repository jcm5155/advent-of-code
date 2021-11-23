package util

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// ReadPuzzleInput - reads puzzle inputs, splits on sep
func ReadPuzzleInput(filenum, sep string) []string {
	pzl, err := ioutil.ReadFile("./inputs/day" + filenum + ".txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(pzl), sep)
}

// ReadRawPuzzleInput - reads puzzle inputs, no splits
func ReadRawPuzzleInput(filenum string) string {
	pzl, err := ioutil.ReadFile("./inputs/day" + filenum + ".txt")
	if err != nil {
		panic(err)
	}
	return string(pzl)
}

// ReadIntPuzzleInput - reads puzzle inputs as ints, sep on \n
func ReadIntPuzzleInput(filenum string) []int {
	p, err := ioutil.ReadFile("./inputs/day" + filenum + ".txt")
	if err != nil {
		panic(err)
	}
	pz := strings.Split(string(p), "\n")
	pzl := []int{}
	for _, i := range pz {
		a, _ := strconv.Atoi(i)
		pzl = append(pzl, a)
	}
	return pzl
}

// ArrAtoi - converts []string to []int
func ArrAtoi(s []string) []int {
	var output []int
	for _, i := range s {
		n, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		output = append(output, n)
	}
	return output
}
