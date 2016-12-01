package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type happinessmap map[string]map[string]int

func main() {
	in, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	hmap := getHapinessMapFromInput(string(in))

	smax, max := hmap.getOptimalSeating()
	fmt.Println("Most optimal seating:", smax, "| Hapiness:", max)

	// part 2: include yourself
	hmap["Me"] = make(map[string]int)
	for p, _ := range hmap {
		hmap["Me"][p] = 0
		hmap[p]["Me"] = 0
	}

	smax, max = hmap.getOptimalSeating()
	fmt.Println("Most optimal seating including me:", smax, "| Hapiness:", max)
}

func getHapinessMapFromInput(in string) happinessmap {
	hmap := happinessmap{}
	for _, line := range strings.Split(string(in), "\n") {
		word := strings.Split(line, " ")
		if hmap[word[0]] == nil {
			hmap[word[0]] = make(map[string]int)
		}
		hap, _ := strconv.Atoi(word[3])
		if word[2] == "lose" {
			hap = 0 - hap
		}
		hmap[word[0]][word[10][:len(word[10])-1]] = hap
	}
	return hmap
}

func (h happinessmap) getOptimalSeating() ([]string, int) {
	var seatings [][]string
	for p, _ := range h {
		seatings = append(h.getRemainingPossibleSeatings([]string{p}))
	}
	var max int
	var smax []string
	for _, s := range seatings {
		h := h.TotalHappinessOfSeating(s)
		if h > max || max == 0 {
			max = h
			smax = s
		}
	}
	return smax, max
}

func (h happinessmap) getRemainingPossibleSeatings(seating []string) [][]string {
	var s [][]string
	if len(h) == len(seating) {
		s = append(s, append(seating, seating[0]))
	} else {
		for person, _ := range h {
			if !sliceContains(seating, person) {
				r := h.getRemainingPossibleSeatings(append(seating, person))
				s = append(s, r...)
			}
		}
	}
	return s
}

func (h happinessmap) TotalHappinessOfSeating(seating []string) int {
	var total int
	for i := 0; i < len(seating)-1; i++ {
		total += h[seating[i]][seating[i+1]]
		total += h[seating[i+1]][seating[i]]
	}
	return total
}

func sliceContains(s []string, v string) bool {
	for _, val := range s {
		if val == v {
			return true
		}
	}
	return false
}
