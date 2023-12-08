package aoc_2023

import (
	"cmp"
	"fmt"
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

	var handsP1, handsP2 []*d7_Hand
	for _, row := range pzl {
		r := strings.Split(row, " ")
		cardStr := strings.Split(r[0], "")
		bet, _ := strconv.Atoi(r[1])
		handsP1 = append(handsP1, d7_newHand(cardStr, bet, false))
		handsP2 = append(handsP2, d7_newHand(cardStr, bet, true))
	}

	slices.SortStableFunc(handsP1, d7_compareHands)
	slices.SortStableFunc(handsP2, d7_compareHands)

	var p1, p2 int
	for idx, hand := range handsP1 {
		p1 += (idx + 1) * hand.Bet
		p2 += (idx + 1) * handsP2[idx].Bet
	}
	return p1, p2
}

type d7_Hand struct {
	Cards []int
	Bet   int
	Type  d7_Type
}

func d7_newHand(cardStr []string, bet int, isPartTwo bool) *d7_Hand {
	var handMap = make(map[string]int)
	var cards = make([]int, 5)
	for _, card := range cardStr {
		cardValue := d7_getCardValue[card]
		if isPartTwo && card == "J" {
			cardValue = 1
		}
		cards = append(cards, cardValue)
		handMap[card]++
	}

	return &d7_Hand{
		Cards: cards,
		Bet:   bet,
		Type:  d7_getHandType(handMap, isPartTwo),
	}
}

func d7_getHandType(handMap map[string]int, isPartTwo bool) d7_Type {
	if len(handMap) == 1 {
		return d7_FiveOfAKind
	}

	if len(handMap) == 5 {
		if isPartTwo {
			if _, ok := handMap["J"]; ok {
				return d7_OnePair
			}
		}
		return d7_HighCard
	}

	var jokerCount int
	var hasFour, hasThree, hasTwo, hasTwoPair bool
	for card, count := range handMap {
		if isPartTwo && card == "J" {
			jokerCount = count
			continue
		}

		switch count {
		case 4:
			hasFour = true
		case 3:
			hasThree = true
		case 2:
			if hasTwo {
				hasTwoPair = true
			}
			hasTwo = true
		}
	}

	// get handtype
	var handType d7_Type
	if hasFour {
		handType = d7_FourOfAKind
	} else if hasThree && hasTwo {
		handType = d7_FullHouse
	} else if hasThree {
		handType = d7_ThreeOfAKind
	} else if hasTwoPair {
		handType = d7_TwoPair
	} else if hasTwo {
		handType = d7_OnePair
	}

	// apply p2 translations
	if isPartTwo {
	JokerLoop:
		for i := 0; i < jokerCount; i++ {
			switch handType {
			case d7_FiveOfAKind:
				break JokerLoop
			case d7_FourOfAKind:
				handType = d7_FiveOfAKind
			case d7_FullHouse:
				handType = d7_FourOfAKind
			case d7_ThreeOfAKind:
				handType = d7_FourOfAKind
			case d7_TwoPair:
				handType = d7_FullHouse
			case d7_OnePair:
				handType = d7_ThreeOfAKind
			case d7_HighCard:
				handType = d7_OnePair
			default:
				fmt.Printf("didnt match a case: handtype=%v\n", handType)
			}
		}
	}

	return handType
}

func d7_compareHands(hand *d7_Hand, other *d7_Hand) int {
	if hand.Type != other.Type {
		return cmp.Compare(hand.Type, other.Type)
	}

	for idx, card := range hand.Cards {
		otherCard := other.Cards[idx]
		if card != otherCard {
			return cmp.Compare(card, otherCard)
		}
	}
	return -1
}
