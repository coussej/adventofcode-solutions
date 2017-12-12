package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type list []string
type village map[string]list

func main() {

	village := getInput()

	// part 1

	count := 0
	for id := range village {
		if village.ProgramsCanConnect(id, "0", list{}) {
			count++
		}
	}
	fmt.Printf("The village has %v programs that can communicate with program 0.\n", count)

	// part 2

	groups := []list{}
	groups = append(groups, list{"0"})

prg:
	for p := range village {
		for g := range groups {
			if village.ProgramsCanConnect(p, groups[g][0], list{}) {
				groups[g] = append(groups[g], p)
				continue prg
			}
		}
		groups = append(groups, list{p})
	}
	fmt.Printf("There are %v separate groups in the villange.\n", len(groups))
}

func (v village) ProgramsCanConnect(p1, p2 string, visited list) bool {
	if p1 == p2 {
		return true
	}
	for _, p := range v[p1] {
		if visited.Has(p) {
			continue
		}
		visited = append(visited, p)
		if p1 != p && v.ProgramsCanConnect(p, p2, visited) {
			return true
		}
	}
	return false
}

func (l list) Has(element string) (has bool) {
	for _, el := range l {
		if el == element {
			has = true
		}
	}
	return
}

func getInput() (v village) {
	v = make(map[string]list)
	in, _ := ioutil.ReadFile("input.txt")
	for _, line := range strings.Split(string(in), "\n") {
		if len(line) > 0 {
			words := strings.Split(strings.Replace(line, " ", "", -1), "<->")
			v[words[0]] = strings.Split(words[1], ",")
		}
	}
	return
}
