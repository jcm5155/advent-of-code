package aoc_2023

import (
	"log"
	"regexp"
	"strings"

	aoc "github.com/jcm5155/advent-of-code/common"
)

type d8_Node struct {
	Left   string
	Right  string
	Symbol string
}

var d8_nodeMap map[string]*d8_Node

func (h *Handler) Day8() (int, int) {
	pzl := aoc.ReadInput("2023", "8").StringLines("\n\n")
	nodeRegex := regexp.MustCompile(`(.{3}) = \((.{3}), (.{3})\)`)
	d8_nodeMap = make(map[string]*d8_Node)
	instructions := pzl[0]

	for _, row := range strings.Split(pzl[1], "\n") {
		match := nodeRegex.FindAllStringSubmatch(row, -1)[0]
		parent, left, right := match[1], match[2], match[3]

		d8_nodeMap[parent] = &d8_Node{
			Left:   left,
			Right:  right,
			Symbol: parent,
		}
	}

	return d8p1(instructions), d8p2(instructions)
}

func d8p1(instructions string) int {
	return d8_solvePath(d8_nodeMap["AAA"], instructions, false)
}

func d8p2(instructions string) int {
	var paths []*d8_Node
	for symbol, node := range d8_nodeMap {
		if symbol[2] == 'A' {
			paths = append(paths, node)
		}
	}

	var stepCounts = make([]int, len(paths))
	for idx, p := range paths {
		currentNode := p
		stepCounts[idx] = d8_solvePath(currentNode, instructions, true)
	}

	return aoc.LCM(stepCounts...)
}

func d8_solvePath(start *d8_Node, instructions string, isPartTwo bool) int {
	var stepCount int
	currentNode := start
	for {
		for _, step := range instructions {
			switch step {
			case 'R':
				currentNode = d8_nodeMap[currentNode.Right]
			case 'L':
				currentNode = d8_nodeMap[currentNode.Left]
			default:
				log.Panicf("unknown instruction: %c", step)
			}

			stepCount++

			if currentNode.Symbol == "ZZZ" || (isPartTwo && currentNode.Symbol[2] == 'Z') {
				return stepCount
			}
		}
	}
}
