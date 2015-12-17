package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type AuntSues map[int]AuntProps
type AuntProps map[string]int

func main() {
	a := getAuntSues()
	r := AuntProps{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
	fmt.Println("According to the first interpretation the aunt is Sue #", a.FindMatch(r))
	fmt.Println("According to the second interpretation the aunt is Sue #", a.FindMatch2(r))
}

func getAuntSues() AuntSues {
	in, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	aunts := make(AuntSues)
	for _, line := range strings.Split(string(in), "\n") {
		words := strings.Split(line, " ")
		s, _ := strconv.Atoi(strings.Replace(words[1], ":", "", -1))
		p := make(AuntProps)
		for _, pr := range strings.Split(strings.Join(words[2:], ""), ",") {
			v := strings.Split(pr, ":")
			n, _ := strconv.Atoi(v[1])
			p[v[0]] = n
		}
		aunts[s] = p
	}
	return aunts
}

func (a AuntSues) FindMatch(ap AuntProps) int {
	for k, p := range a {
		match := true
		for prop, number := range p {
			if ap[prop] != number {
				match = false
				break
			}
		}
		if match {
			return k
		}
	}
	return -1
}

func (a AuntSues) FindMatch2(ap AuntProps) int {
	for k, p := range a {
		match := true
	propertyLoop:
		for prop, number := range p {
			switch {
			case prop == "cats" || prop == "trees":
				if ap[prop] >= number {
					match = false
					break propertyLoop
				}
			case prop == "pomeranians" || prop == "goldfish":
				if ap[prop] <= number {
					match = false
					break propertyLoop
				}
			default:
				if ap[prop] != number {
					match = false
					break
				}
			}
		}
		if match {
			return k
		}
	}
	return -1
}
