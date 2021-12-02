package shared

import "strconv"

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
