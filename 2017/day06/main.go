package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type memory []int

func main() {
	mem := getInput()
	steps, cycle := mem.redistributeUntilRepeat()
	fmt.Printf("It takes %v redistribution steps to find a repeated state.\n", steps)
	fmt.Printf("The size of the infinite loop is %v.\n", cycle)
}

func (m memory) redistributeUntilRepeat() (steps, cycle int) {
	states := map[string]int{}
	states[m.getState()] = steps

	for {
		steps++
		pos, toDistribute := 0, 0
		for i, x := range m {
			if x > toDistribute {
				pos, toDistribute = i, x
			}
		}
		m[pos] = 0
		for toDistribute > 0 {
			pos++
			if pos == len(m) {
				pos = 0
			}
			m[pos]++
			toDistribute--
		}
		state := m.getState()
		if v, ok := states[state]; ok {
			cycle = steps - v
			return
		}

		states[state] = steps
	}
}

func (m *memory) getState() (state string) {
	for _, v := range *m {
		s := strconv.Itoa(v)
		state = state + "." + s
	}
	return
}

func getInput() (input memory) {
	in, _ := ioutil.ReadFile("input.txt")
	line := strings.Split(string(in), "\n")[0]
	for _, n := range strings.Split(line, "\t") {
		num, _ := strconv.Atoi(n)
		input = append(input, num)
	}
	return
}
