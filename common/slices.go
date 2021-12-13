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

func SliceAtoFloat64(s []string) []float64 {
	var output []float64
	for _, i := range s {
		n, err := strconv.ParseFloat(i, 64)
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

func SliceSum(s []int) int {
	var output int
	for _, item := range s {
		output += item
	}
	return output
}

func SliceAny(s []bool) bool {
	for _, item := range s {
		if item {
			return true
		}
	}
	return false
}

func SliceAll(s []bool) bool {
	for _, item := range s {
		if !item {
			return false
		}
	}
	return true
}

func SliceNone(s []bool) bool {
	for _, item := range s {
		if item {
			return false
		}
	}
	return true
}
