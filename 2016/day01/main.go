package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type instruction struct {
	direction string
	steps     int
}

type position struct {
	lat, lon int
}

func (p *position) distanceFromOrigin() int {
	return int(math.Abs(float64(p.lat)) + math.Abs(float64(p.lon)))
}

type route struct {
	currentDirection          int
	totalStepsInEachDirection [4]int // [North, East, South, West]
	path                      []position
}

func (r *route) move(instr instruction) {
	// rotate to next direction
	if instr.direction == "R" {
		r.currentDirection++
		if r.currentDirection == 4 {
			r.currentDirection = 0
		}
	} else {
		r.currentDirection--
		if r.currentDirection == -1 {
			r.currentDirection = 3
		}
	}

	// take steps and record every position in path.
	for i := 0; i < instr.steps; i++ {
		r.totalStepsInEachDirection[r.currentDirection]++
		pos := position{
			lat: r.totalStepsInEachDirection[0] - r.totalStepsInEachDirection[2],
			lon: r.totalStepsInEachDirection[1] - r.totalStepsInEachDirection[3],
		}
		r.path = append(r.path, pos)
	}
}

func getInstructions() (instructions []instruction) {
	in, _ := ioutil.ReadFile("input.txt")
	for _, v := range strings.Split(strings.Replace(string(in), "\n", "", -1), ", ") {
		direction := v[0:1]
		steps, _ := strconv.Atoi(v[1:])
		instructions = append(instructions, instruction{direction, steps})
	}
	return
}

func main() {

	instructions := getInstructions()
	r := route{}

	for _, instr := range instructions {
		r.move(instr)
	}

	fmt.Println("Easter Bunny HQ is", r.path[len(r.path)-1].distanceFromOrigin(),
		"blocks away.")

	var firstRevisit position
firstRevisitSearch:
	for n, pos := range r.path {
		for i := 0; i < n; i++ {
			if pos == r.path[i] {
				firstRevisit = pos
				break firstRevisitSearch
			}
		}
	}

	fmt.Println("According to the new instructions on the back of the paper,",
		"Easter Bunny HQ is", firstRevisit.distanceFromOrigin(), "blocks away.")
}
