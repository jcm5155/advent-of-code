package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadInput(year, day string) *PuzzleInput {
	fullpath := fmt.Sprintf("./%v/inputs/day%v.input", year, day)
	pzl, err := os.ReadFile(filepath.FromSlash(fullpath))
	if err != nil {
		panic(err)
	}
	return &PuzzleInput{bytes: pzl}
}

type PuzzleInput struct {
	bytes []byte
}

func (p *PuzzleInput) String() string {
	return strings.TrimSuffix(string(p.bytes), "\n")
}

func (p *PuzzleInput) StringLines(sep string) []string {
	return strings.Split(p.String(), sep)
}

func (p *PuzzleInput) IntLines(sep string) []int {
	return SliceAtoi(p.StringLines(sep))
}

func (p *PuzzleInput) UintLines(sep string) []uint {
	return SliceAtoUint(p.StringLines(sep))
}

func (p *PuzzleInput) Float64Lines(sep string) []float64 {
	return SliceAtoFloat64(p.StringLines(sep))
}
