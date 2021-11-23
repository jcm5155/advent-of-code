package solutions

import (
	"fmt"
	"reflect"
)

var hd = Handler{}

// Handler - handler for all solutions (for reflection)
type Handler struct{}

// Solution - answer to a day's problems
type Solution struct {
	DayNum  int
	PartOne reflect.Value
	PartTwo reflect.Value
}

// Print - prints solution parts for a given Solution (or skips if no solutions are found)
func (s *Solution) Print() {
	if !reflect.ValueOf(s.PartOne).IsZero() || !reflect.ValueOf(s.PartTwo).IsZero() {
		fmt.Printf("*-=DAY %v=-*\np1: %v\np2: %v\n\n", s.DayNum, s.PartOne, s.PartTwo)
	}
}

// Day - returns the Solution for a given day
func Day(dayNum int) *Solution {
	p1, p2 := reflect.Value{}, reflect.Value{}
	m := reflect.ValueOf(&hd).MethodByName(fmt.Sprintf("Day%v", dayNum))
	if m.IsValid() {
		p := m.Call(nil)
		p1, p2 = p[0], p[1]
	}

	return &Solution{
		DayNum:  dayNum,
		PartOne: p1,
		PartTwo: p2,
	}
}
