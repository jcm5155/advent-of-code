package shared

import (
	"fmt"
	"reflect"
)

func Solve(dayNum int, h interface{}) *Solution {
	p1, p2 := reflect.Value{}, reflect.Value{}
	m := reflect.ValueOf(h).MethodByName(fmt.Sprintf("Day%v", dayNum))
	if m.IsValid() {
		p := m.Call(nil)
		p1, p2 = p[0], p[1]
	}

	return &Solution{
		Day:     dayNum,
		PartOne: p1,
		PartTwo: p2,
	}
}

type Solution struct {
	Day     int
	PartOne reflect.Value
	PartTwo reflect.Value
}

func (s *Solution) Print() {
	if !reflect.ValueOf(s.PartOne).IsZero() || !reflect.ValueOf(s.PartTwo).IsZero() {
		fmt.Printf("*-=DAY %v=-*\np1: %v\np2: %v\n\n", s.Day, s.PartOne, s.PartTwo)
	}
}
