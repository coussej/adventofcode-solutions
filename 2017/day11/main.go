package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	steps, max := getStepsAwayFromStart(getInput())
	fmt.Printf("The fewest number of steps to reach the child process is %v.\n", steps)
	fmt.Printf("The furthest away the child process ever got was %v steps.\n", max)

}

func getStepsAwayFromStart(directions []string) (steps, max float64) {
	x, y := 0.0, 0.0
	for _, d := range directions {
		switch d {
		case "n":
			y = y + 1
		case "ne":
			x, y = x+.5, y+.5
		case "nw":
			x, y = x-.5, y+.5
		case "s":
			y = y - 1
		case "se":
			x, y = x+.5, y-.5
		case "sw":
			x, y = x-.5, y-.5
		}
		if dist := math.Abs(x) + math.Abs(y); dist > max {
			max = dist
		}
	}
	steps = math.Abs(x) + math.Abs(y)
	return
}

func getInput() (directions []string) {
	in, _ := ioutil.ReadFile("input.txt")
	for _, d := range strings.Split(strings.Replace(string(in), "\n", "", -1), ",") {
		directions = append(directions, d)
	}
	return
}
