package common

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(integers ...int) int {
	a := integers[0]
	b := integers[1]

	result := a * b / GCD(a, b)

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
