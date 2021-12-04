package common

import (
	"reflect"
)

type Solution struct {
	Day   string
	Part1 reflect.Value
	Part2 reflect.Value
}

func Solve(h interface{}, day string) *Solution {
	p1, p2 := reflect.Value{}, reflect.Value{}
	m := reflect.ValueOf(h).MethodByName("Day" + day)
	if m.IsValid() {
		p := m.Call(nil)
		p1, p2 = p[0], p[1]
	}
	return &Solution{
		Day:   day,
		Part1: handleZeroValues(p1),
		Part2: handleZeroValues(p2),
	}
}

func handleZeroValues(v reflect.Value) reflect.Value {
	if reflect.ValueOf(v).IsZero() {
		return reflect.ValueOf("?")
	}
	return v
}
