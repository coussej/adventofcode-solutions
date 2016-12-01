package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type output struct {
	result      uint16
	resultknown bool
	source      []string
}

type circuit map[string]*output

func buildCircuit(instructions string) circuit {
	lines := strings.Split(instructions, "\n")
	c := circuit{}
	for _, v := range lines {
		words := strings.Split(v, " ")
		out := words[len(words)-1]
		in := words[:len(words)-2]
		c[out] = &output{0, false, in}
	}
	return c
}

func (c circuit) getOutputValue(o string) uint16 {
	if !c[o].resultknown {
		// parse source input
		switch len(c[o].source) {

		case 1:
			// direct assignment. Check if other output or value
			var v1 uint16
			v, err := strconv.ParseUint(c[o].source[0], 10, 16)
			if err != nil {
				// is reference to another output. Recurse!
				v1 = c.getOutputValue(c[o].source[0])
			} else {
				v1 = uint16(v)
			}
			c[o].result = v1

		case 2:
			// the only instruction with 2 words is the NOT instruction.
			var v1 uint16
			v, err := strconv.ParseUint(c[o].source[1], 10, 16)
			if err != nil {
				// is reference to another output. Recurse!
				v1 = c.getOutputValue(c[o].source[1])
			} else {
				v1 = uint16(v)
			}
			// perform NOT on val for this output.
			c[o].result = v1 ^ 65535

		case 3:
			var v1, v2 uint16
			v, err := strconv.ParseUint(c[o].source[0], 10, 16)
			if err != nil {
				// is reference to another output. Recurse!
				v1 = c.getOutputValue(c[o].source[0])
			} else {
				v1 = uint16(v)
			}
			v, err = strconv.ParseUint(c[o].source[2], 10, 16)
			if err != nil {
				// is reference to another output. Recurse!
				v2 = c.getOutputValue(c[o].source[2])
			} else {
				v2 = uint16(v)
			}
			switch c[o].source[1] {
			case "AND":
				c[o].result = v1 & v2
			case "OR":
				c[o].result = v1 | v2
			case "RSHIFT":
				c[o].result = v1 >> v2
			case "LSHIFT":
				c[o].result = v1 << v2
			}
		}
	}
	c[o].resultknown = true
	return c[o].result
}

func main() {
	in, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	c := buildCircuit(string(in))

	// part 1: get the output on wire a.
	a := c.getOutputValue("a")
	fmt.Println("The signal provided to wire a is", a)

	// part 2: create a new circuit, set source for wire b to a of previous excercise and find a again.
	c = buildCircuit(string(in))
	c["b"].source = []string{strconv.Itoa(int(a))}
	fmt.Println("When overriding b with the signal of a in the previous setup, a is now", c.getOutputValue("a"))
}
