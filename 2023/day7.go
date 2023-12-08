package aoc_2023

import (
	"slices"
	"strconv"
	"strings"

	aoc "github.com/jcm5155/advent-of-code/common"
)

type d7_Type int

const (
	d7_HighCard d7_Type = iota
	d7_OnePair
	d7_TwoPair
	d7_ThreeOfAKind
	d7_FullHouse
	d7_FourOfAKind
	d7_FiveOfAKind
)

var d7_getCardValue = map[string]int{
	"A": 14, "K": 13, "Q": 12, "J": 11, "T": 10, "9": 9,
	"8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2,
}

func (h *Handler) Day7() (int, int) {
	pzl := aoc.ReadInput("2023", "7").StringLines("\n")

	var hands []*d7_Hand
	for _, row := range pzl {
		r := strings.Split(row, " ")
		cardStr := strings.Split(r[0], "")
		bet, _ := strconv.Atoi(r[1])
		hands = append(hands, d7_newHand(cardStr, bet))
	}

	slices.SortStableFunc(hands, d7_compareHands)

	var p1, p2 int

	for idx, hand := range hands {
		p1 += (idx + 1) * hand.Bet
	}

	return p1, p2
}

type d7_Hand struct {
	Cards []int
	Bet   int
	Type  d7_Type
}

func d7_newHand(cardStr []string, bet int) *d7_Hand {
	var handMap = make(map[string]int)
	var cards = make([]int, 5)
	for _, card := range cardStr {
		cardValue := d7_getCardValue[card]
		cards = append(cards, cardValue)
		handMap[card]++
	}

	// there has to be a better way
	var handType d7_Type
	if len(handMap) == 1 {
		handType = d7_FiveOfAKind
	} else if len(handMap) == 5 {
		handType = d7_HighCard
	} else {
		var hasThree, hasTwo, hasTwoPair bool
	Loop:
		for _, count := range handMap {
			switch count {
			case 4:
				handType = d7_FourOfAKind
				break Loop
			case 3:
				hasThree = true
			case 2:
				if hasTwo {
					hasTwoPair = true
				}
				hasTwo = true
			}
		}

		if hasThree && hasTwo {
			handType = d7_FullHouse
		} else if hasThree {
			handType = d7_ThreeOfAKind
		} else if hasTwoPair {
			handType = d7_TwoPair
		} else if hasTwo {
			handType = d7_OnePair
		}
	}

	return &d7_Hand{
		Cards: cards,
		Bet:   bet,
		Type:  handType,
	}
}

func d7_compareHands(hand *d7_Hand, other *d7_Hand) int {
	if hand.Type != other.Type {
		if hand.Type > other.Type {
			return 1
		}
		return -1
	}

	for idx, card := range hand.Cards {
		otherCard := other.Cards[idx]

		if card == otherCard {
			continue
		}

		if card > otherCard {
			return 1
		}
		break
	}
	return -1
}
