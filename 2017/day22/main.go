package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type cluster struct {
	curNode      coordinate
	nodeStates   map[coordinate]int
	curDirection int
	directions   []coordinate
	states       []string
}

type coordinate struct {
	x, y int
}

func main() {
	fmt.Printf("%v bursts result in infection.\n", getInput().simulate(10000, 2))
	fmt.Printf("%v bursts result in infection using the evolved virus.\n", getInput().simulate(10000000, 1))
}

func (c cluster) simulate(bursts, stateJump int) (infections int) {
	for i := 0; i < bursts; i++ {
		switch c.states[c.nodeStates[c.curNode]] {
		case "clean":
			c.turnLeft()
		case "infected":
			c.turnRight()
		case "flagged":
			c.turnRight()
			c.turnRight()
		}
		c.currentNodeNextState(stateJump)
		if c.states[c.nodeStates[c.curNode]] == "infected" {
			infections++
		}
		c.move()
	}
	return
}

func (c *cluster) turnLeft() {
	c.curDirection = (c.curDirection + len(c.directions) - 1) % len(c.directions)
}

func (c *cluster) turnRight() {
	c.curDirection = (c.curDirection + 1) % len(c.directions)
}

func (c *cluster) move() {
	c.curNode.x += c.directions[c.curDirection].x
	c.curNode.y += c.directions[c.curDirection].y
}

func (c *cluster) currentNodeNextState(jump int) {
	c.nodeStates[c.curNode] = (c.nodeStates[c.curNode] + jump) % len(c.states)
}

func getInput() (c cluster) {
	in, _ := ioutil.ReadFile("input.txt")
	c = cluster{
		curNode:    coordinate{},
		nodeStates: map[coordinate]int{},
		directions: []coordinate{
			coordinate{0, 1},
			coordinate{1, 0},
			coordinate{0, -1},
			coordinate{-1, 0},
		},
		states: []string{"clean", "weakened", "infected", "flagged"},
	}
	lines := strings.Split(strings.TrimSuffix(string(in), "\n"), "\n")
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				c.nodeStates[coordinate{x - len(line)/2, -y + len(lines)/2}] = 2
			}
		}
	}
	return
}
