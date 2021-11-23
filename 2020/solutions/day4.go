package solutions

import (
	"regexp"
	"strings"

	"github.com/jcm5155/advent-of-code-2020/util"
)

var d4Fields = map[string]*regexp.Regexp{
	"byr": regexp.MustCompile(`^19[^01]\d$|^200[0-2]$`),                    // four digits; at least 1920 and at most 2002
	"iyr": regexp.MustCompile(`^20(1\d|20)$`),                              // four digits; at least 2010 and at most 2020.
	"eyr": regexp.MustCompile(`^20(2\d|30)$`),                              // four digits; at least 2020 and at most 2030.
	"hgt": regexp.MustCompile(`^1([5-8]\d|9[0-3])cm$|^(59|6\d|7[0-6])in$`), // a number followed by either cm or in: (and additional logic, lol)
	"hcl": regexp.MustCompile(`^#[\d|a-f]{6}$`),                            // a # followed by exactly six characters 0-9 or a-f.
	"ecl": regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`),           // exactly one of: amb blu brn gry grn hzl oth.
	"pid": regexp.MustCompile(`^\d{9}$`),                                   // a nine-digit number, including leading zeroes.
}

// Day4 solution
func (h *Handler) Day4() (int, int) {
	pzl := util.ReadPuzzleInput("4", "\n\n")
	p1, p2 := 0, 0
	for _, passport := range pzl {
		fullPP := strings.Fields(passport)
		if len(fullPP) < 7 {
			continue
		}
		m := make(map[string]string)
		for _, i := range fullPP {
			parts := strings.Split(i, ":")
			m[parts[0]] = parts[1]
		}
		if p1Validate(m) {
			p1++
			if p2Validate(m) {
				p2++
			}
		}
	}
	return p1, p2
}

func p1Validate(m map[string]string) bool {
	for k := range d4Fields {
		if m[k] == "" {
			return false
		}
	}
	return true
}

func p2Validate(m map[string]string) bool {
	for k, v := range d4Fields {
		if v.FindString(m[k]) == "" {
			return false
		}
	}
	return true
}
