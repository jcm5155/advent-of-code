package common

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func ReadInput(year, day string) *PuzzleInput {
	fullpath := fmt.Sprintf("./%v/inputs/day%v.input", year, day)
	pzl, err := ioutil.ReadFile(filepath.FromSlash(fullpath))
	if err != nil {
		panic(err)
	}
	return &PuzzleInput{bytes: pzl}
}

type PuzzleInput struct {
	bytes []byte
}

func (p *PuzzleInput) String() string {
	return string(p.bytes)
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
