package shared

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

const (
	puzzleInputFileExtension = ".input"
)

func ReadInput(year, day string) *PuzzleInput {
	fullpath := fmt.Sprintf("./%v/inputs/day%v%v", year, day, puzzleInputFileExtension)
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
	return ArrAtoi(p.StringLines(sep))
}
