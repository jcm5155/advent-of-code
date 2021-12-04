package aoc_2020

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/jcm5155/advent-of-code/common"
)

// Day13 solution
func (h *Handler) Day13() (int, int) {
	pzl := common.ReadInput("2020", "13").StringLines("\n")
	target, _ := strconv.ParseFloat(pzl[0], 64)
	busIDs := strings.Split(pzl[1], ",")
	var c = make(chan int)
	go d13p2(busIDs, c)
	return d13p1(target, busIDs), <-c
}

func d13p1(target float64, busIDs []string) int {
	var earliest, diff, time float64
	for _, busID := range busIDs {
		if busID != "x" {
			n, _ := strconv.ParseFloat(busID, 64)
			curr := (math.Floor(target/n) + 1) * n
			if earliest == 0 || curr < earliest {
				earliest = curr
				diff = curr - target
				time = n
			}
		}
	}
	return int(time * diff)
}

func d13p2(busIDs []string, c chan int) {
	query := "0="
	for offset, busID := range busIDs {
		if busID != "x" {
			query += fmt.Sprintf("+%%3D((n+%%2B+%v)+mod+%v)", offset, busID)
		}
	}
	// https://www.youtube.com/watch?v=oavMtUWDBTM
	r, _ := http.Get(fmt.Sprintf("http://api.wolframalpha.com/v2/query?input=%v&appid=%v", query, os.Getenv("WOLFRAMALPHA_API_KEY")))
	b, _ := ioutil.ReadAll(r.Body)
	pattern := regexp.MustCompile(`n = (?:\d+) m \+ (\d+), m element Z`)
	p2, _ := strconv.Atoi(string(pattern.FindSubmatch(b)[1]))
	c <- p2
}
