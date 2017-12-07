package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type program struct {
	name     string
	weight   int
	children map[string]*program
}

func (p *program) getTotalWeight() (weight int) {
	weight += p.weight
	for c := range p.children {
		weight += p.children[c].getTotalWeight()
	}
	return
}

func main() {
	root := getRootProgram()
	fmt.Printf("The name of the bottom program is %v.\n", root.name)

	p, w := getIncorrectProgram(root)
	fmt.Printf("Program %v should have a weight of %v to obtain balance.\n", p.name, w)
}

func getRootProgram() (root *program) {

	// temporary flat list of all programs in the file
	list := map[string]*program{}

	// read the file and loop over al lines
	in, _ := ioutil.ReadFile("input.txt")
	for _, l := range strings.Split(string(in), "\n") {
		if len(l) > 0 {

			// create a program and extract its field from the text line
			p := program{}
			words := strings.Split(l, " ")
			p.name = words[0]
			p.weight, _ = strconv.Atoi(words[1][1 : len(words[1])-1])

			// set the program children name, with a nil pointer, these will be assigned later
			p.children = map[string]*program{}
			if len(words) > 3 {
				for _, c := range strings.Split(strings.Join(words[3:], ""), ",") {
					p.children[c] = nil
				}
			}

			// add the program to the list
			list[p.name] = &p

		}
	}

	// loop over all programs in the list, assign all the childres and find the root.
	for n, p := range list {

		// assign the children with the correct pointers
		for c := range p.children {
			list[p.name].children[c] = list[c]
		}

		// check if this program is the root, ie no others have it as a child.
		isRoot := true
		for _, p2 := range list {
			if _, exists := p2.children[n]; exists {
				isRoot = false
				break
			}
		}
		if isRoot {
			root = list[n]
		}

	}
	return
}

func getIncorrectProgram(tree *program) (p *program, correctedWeight int) {
	p = tree
	idealWeight := 0
	for {
		// assuming imbalance can only occur with 3 children or more. In case of 2
		// children multiple answers would be possible.
		wrongWeight, wrongName, rightWeight, rightName, totalWeight := 0, "", 0, "", 0

		for n := range p.children {
			childWeight := p.children[n].getTotalWeight()
			switch {
			case rightWeight == 0 || rightWeight == childWeight:
				rightWeight, rightName = childWeight, n
			case childWeight == wrongWeight:
				wrongWeight, rightWeight = rightWeight, wrongWeight
				wrongName, rightName = rightName, wrongName
			default:
				wrongWeight, wrongName = childWeight, n
			}
			totalWeight += childWeight
		}
		if wrongWeight == 0 {
			correctedWeight = idealWeight - totalWeight
			return
		}

		// descend to the child with the wrong weight, while keeping track of the
		// total weight it should have
		p = p.children[wrongName]
		idealWeight = rightWeight
	}
}
