package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getInstructions() (instr []string) {
	in, _ := ioutil.ReadFile("input.txt")
	for _, line := range strings.Split(string(in), "\n") {
		if line != "" {
			instr = append(instr, line)
		}
	}
	return
}

type registry map[string]int

func (r registry) runInstructions(instr []string) {
	n := 0
	for n < len(instr) {
		if n < 0 {
			n = 0
		}
		words := strings.Split(instr[n], " ")
		switch words[0] {
		case "cpy":
			x, err := strconv.Atoi(words[1])
			if err != nil {
				x = r[words[1]]
			}
			r[words[2]] = x
		case "inc":
			r[words[1]]++
		case "dec":
			r[words[1]]--
		case "jnz":
			x, err := strconv.Atoi(words[1])
			if err != nil {
				x = r[words[1]]
			}
			y, err := strconv.Atoi(words[2])
			if err != nil {
				y = r[words[2]]
			}
			if x != 0 {
				n += y - 1 // because n++ at the end
			}
		}
		n++
	}
	return
}

func main() {
	instr := getInstructions()

	r := registry{}
	r.runInstructions(instr)
	fmt.Printf("The program results in a value of %v in registry a.\n", r["a"])

	r = registry{"c": 1}
	r.runInstructions(instr)
	fmt.Printf("With the position of the ignition key initialized in c, the program results in a value of %v in registry a.\n", r["a"])

}
