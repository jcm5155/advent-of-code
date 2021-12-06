package aoc_2021

import (
	"github.com/jcm5155/advent-of-code/common"
	"strings"
	"sync"
)

func (h *Handler) Day4() (int, int) {
	pzl := common.ReadInput("2021", "4").StringLines("\n\n")
	drawOrder := common.SliceAtoi(strings.Split(pzl[0], ","))
	bingoBoards := pzl[1:]
	numBoards := len(bingoBoards)

	// a spot on a bingo board
	type bingoSpot struct {
		Value    int
		X        int
		Y        int
		IsMarked bool
	}

	// a bingo bingo'd
	type bingoResult struct {
		WinningValue int
		WinningTurn  int
		BoardsIdx    int
	}

	var boards = make([]map[int]bingoSpot, numBoards)
	var results = make([]bingoResult, numBoards)
	var wg sync.WaitGroup

	for idx, board := range bingoBoards {
		wg.Add(1)
		go func(boardListIdx int, board_ string) {
			defer wg.Done()

			// map out the board
			boardMap := make(map[int]bingoSpot, 100) // bingo board contains up to 2-digit numbers only
			for y, row := range strings.Split(board_, "\n") {
				for x, spot := range common.SliceAtoi(strings.Fields(row)) {
					boardMap[spot] = bingoSpot{
						Value: spot,
						X:     x,
						Y:     y,
					}
				}
			}
			boards[boardListIdx] = boardMap

			// simulate the bingo game until...
			marksInColumn, marksInRow := [5]int{}, [5]int{}
			for turnNumber, drawn := range drawOrder {
				if spot, isOnBoard := boardMap[drawn]; isOnBoard {
					spot.IsMarked = true
					boardMap[drawn] = spot
					marksInColumn[spot.X]++
					marksInRow[spot.Y]++

					// https://www.youtube.com/watch?v=HNGXsgLRkXU
					if marksInColumn[spot.X] == 5 || marksInRow[spot.Y] == 5 {
						bingo := bingoResult{BoardsIdx: boardListIdx, WinningValue: drawn, WinningTurn: turnNumber}
						results[boardListIdx] = bingo
						break
					}
				}
			}
		}(idx, board)
	}
	wg.Wait()

	// find the winningest and losingest boards
	p1, p2 := results[0], results[0]
	for _, result := range results[1:] {
		if result.WinningTurn < p1.WinningTurn {
			p1 = result
		}
		if result.WinningTurn > p2.WinningTurn {
			p2 = result
		}
	}

	// get final scores
	scoreOf := func(board map[int]bingoSpot, multiplier int) int {
		var output int
		for _, spot := range board {
			if !spot.IsMarked {
				output += spot.Value
			}
		}
		return output * multiplier
	}

	return scoreOf(boards[p1.BoardsIdx], p1.WinningValue), scoreOf(boards[p2.BoardsIdx], p2.WinningValue)
}
