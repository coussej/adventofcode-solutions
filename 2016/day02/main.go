package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type numpad struct {
	keys        [][]string
	keysPressed []string
	curRow      int
	curCol      int
}

func newNormalNumpad() numpad {
	return numpad{
		[][]string{
			[]string{"_", "_", "_", "_", "_"},
			[]string{"_", "1", "2", "3", "_"},
			[]string{"_", "4", "5", "6", "_"},
			[]string{"_", "7", "8", "9", "_"},
			[]string{"_", "_", "_", "_", "_"},
		},
		[]string{},
		2,
		2,
	}
}

func newWeirdNumpad() numpad {
	return numpad{
		[][]string{
			[]string{"_", "_", "_", "_", "_", "_", "_"},
			[]string{"_", "_", "_", "1", "_", "_", "_"},
			[]string{"_", "_", "2", "3", "4", "_", "_"},
			[]string{"_", "5", "6", "7", "8", "9", "_"},
			[]string{"_", "_", "A", "B", "C", "_", "_"},
			[]string{"_", "_", "_", "D", "_", "_", "_"},
			[]string{"_", "_", "_", "_", "_", "_", "_"},
		},
		[]string{},
		3,
		3,
	}
}

func (n *numpad) move(direction string) {
	switch direction {
	case "U":
		if n.keys[n.curRow-1][n.curCol] != "_" {
			n.curRow--
		}
	case "D":
		if n.keys[n.curRow+1][n.curCol] != "_" {
			n.curRow++
		}
	case "L":
		if n.keys[n.curRow][n.curCol-1] != "_" {
			n.curCol--
		}
	case "R":
		if n.keys[n.curRow][n.curCol+1] != "_" {
			n.curCol++
		}
	}
}

func (n *numpad) pressKey() {
	n.keysPressed = append(n.keysPressed, n.keys[n.curRow][n.curCol])
}

func getInstructions() (instructions [][]string) {
	in, _ := ioutil.ReadFile("input.txt")
	for _, v := range strings.Split(string(in), "\n") {
		if len(v) > 0 {
			instructions = append(instructions, strings.Split(v, ""))
		}
	}
	return
}

func main() {
	numpad1 := newNormalNumpad()
	numpad2 := newWeirdNumpad()
	for _, instr := range getInstructions() {
		for _, dir := range instr {
			numpad1.move(dir)
			numpad2.move(dir)
		}
		numpad1.pressKey()
		numpad2.pressKey()
	}
	fmt.Println("The bathroom code is", numpad1.keysPressed)
	fmt.Println("But with the weird numpad layout, the bathroom code is", numpad2.keysPressed)
}
