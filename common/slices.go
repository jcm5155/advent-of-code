package common

import (
	"strconv"
)

func SliceAtoi(s []string) []int {
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

func SliceAtoUint(s []string) []uint {
	var output []uint
	bitSize := len(s[0])
	for _, binaryLiteral := range s {
		o, _ := strconv.ParseUint(binaryLiteral, 2, bitSize)
		output = append(output, uint(o))
	}
	return output
}
