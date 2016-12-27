package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type position struct {
	x, y int
}

func (p position) add(p2 position) (sum position) {
	return position{p.x + p2.x, p.y + p2.y}
}

func (p position) isWithinLimits(xmax, ymax int) bool {
	return p.x >= 0 && p.y >= 0 && p.x <= xmax && p.y <= ymax
}

func (p position) getPossibleDests(xmax, ymax int) (n []position) {
	moves := []position{position{1, 0}, position{0, 1},
		position{0, -1}, position{-1, 0}}
	for _, m := range moves {
		new := p.add(m)
		if new.isWithinLimits(xmax, ymax) {
			n = append(n, new)
		}
	}
	return
}

type node struct {
	size, used, avail int
}

type grid map[position]node

func (g grid) getLimits() (x, y int) {
	for p := range g {
		if p.x > x {
			x = p.x
		}
		if p.y > y {
			y = p.y
		}
	}
	return
}

func (g grid) getStartPosition() (p position) {
	for p, n := range g {
		if n.used == 0 {
			return p
		}
	}
	return
}

func (g grid) getNonMovableNodePositions() (nonMovable map[position]bool) {
	nonMovable = make(map[position]bool)
	// find node with smallest size
	smallest := 0
	for _, n := range g {
		if smallest == 0 || n.size < smallest {
			smallest = n.size
		}
	}
	// find nodes that don't fit on the smallest node.
	for p, n := range g {
		if n.used > smallest {
			nonMovable[p] = true
		}
	}
	return
}

func getInput() (g grid) {
	g = make(grid)
	in, _ := ioutil.ReadFile("input.txt")
	in2 := regexp.MustCompile(" +").ReplaceAllString(string(in), "-")
	for _, v := range strings.Split(in2, "\n") {
		if strings.Contains(v, "/dev/") {
			words := strings.Split(v, "-")
			x, _ := strconv.Atoi(words[1][1:])
			y, _ := strconv.Atoi(words[2][1:])
			size, _ := strconv.Atoi(words[3][:len(words[3])-1])
			used, _ := strconv.Atoi(words[4][:len(words[4])-1])
			avail, _ := strconv.Atoi(words[5][:len(words[5])-1])
			g[position{x, y}] = node{size, used, avail}
		}
	}
	return
}

func (g grid) getNrOfViablePairs() (count int) {
	for a, na := range g {
		for b, nb := range g {
			if na.used != 0 && a != b && na.used <= nb.avail {
				count++
			}
		}
	}
	return
}

func (g grid) getShortestPath(from, to position, illegal map[position]bool) []position {
	xmax, ymax := g.getLimits()
	// map with current positions and the path they have taken.
	status := map[position][]position{from: []position{}}
	visited := map[position]bool{from: true}
	for {
		nextStatus := map[position][]position{}
		for currentPos, path := range status {
			nextPos := currentPos.getPossibleDests(xmax, ymax)
			for _, next := range nextPos {
				p := make([]position, len(path))
				copy(p, path)
				p = append(p, next)
				if next == to {
					return p
				}
				if !visited[next] && !illegal[next] {
					visited[next] = true
					nextStatus[next] = p
				}
			}
		}
		status = nextStatus
	}
}

func (g grid) getMinimumStepsToMoveGoalData() (steps int) {
	xmax, _ := g.getLimits()
	start := g.getStartPosition()
	illegal := g.getNonMovableNodePositions()

	// get the length of the path from the starting node to the position left of
	// the target node
	steps += len(g.getShortestPath(start, position{xmax - 1, 0}, illegal))

	// get the shortest path the target node must follow to the top left
	targetPath := g.getShortestPath(position{xmax, 0}, position{0, 0}, illegal)

	// follow the path and for each step, calculate the shortest path for
	// the empty node to reach the other side
	currentTarget := position{xmax, 0}
	for i := 0; i < len(targetPath); i++ {
		// swap empty node and targetNode
		steps++

		if i < len(targetPath)-1 {
			// add goal to the illegal nodes
			illegal2 := map[position]bool{targetPath[i]: true}
			for k, v := range illegal {
				illegal2[k] = v
			}
			// get shortest path from position of emptyNode (= currentTarget, cause we
			// just swapped them) to the next position in the path:
			emptyPath := g.getShortestPath(currentTarget, targetPath[i+1], illegal2)
			steps += len(emptyPath)
			currentTarget = targetPath[i]
		}
	}
	return
}

func main() {
	g := getInput()
	fmt.Printf("There are %v viable pairs of nodes in the storage cluster.\n",
		g.getNrOfViablePairs())
	fmt.Printf("You need a minimum of %v steps to move the goal data to position 0,0.\n",
		g.getMinimumStepsToMoveGoalData())
}
