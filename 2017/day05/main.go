package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	list := getList()

	fmt.Printf("It takes %v steps to reach the exit.\n", getStepsToExit(list, false))
	fmt.Printf("With the stranger jumps, it takes %v steps to reach the exit.\n", getStepsToExit(list, true))
}

func getStepsToExit(instructions []int, strangeJumps bool) (steps int) {
	pos := 0
	instr := make([]int, len(instructions))
	copy(instr, instructions)
	for {
		steps++
		jump := instr[pos]
		if pos+jump > len(instr)-1 || pos+jump < 0 {
			return
		}

		if strangeJumps && jump > 2 {
			instr[pos]--
		} else {
			instr[pos]++
		}
		pos = pos + jump
	}
}

func getList() (list []int) {
	in, _ := ioutil.ReadFile("input.txt")
	for _, c := range strings.Split(string(in), "\n") {
		if len(c) > 0 {
			num, _ := strconv.Atoi(string(c))
			list = append(list, num)
		}
	}
	return
}
