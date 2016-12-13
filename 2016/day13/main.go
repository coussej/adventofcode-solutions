package main

import (
	"fmt"
	"strings"
)

type position struct {
	x, y int
}

func (p position) add(p2 position) (sum position) {
	return position{p.x + p2.x, p.y + p2.y}
}

func (p position) isOpenSpace() bool {
	if p.x < 0 || p.y < 0 {
		return false
	}
	designersnumber := 1350
	a := p.x*p.x + 3*p.x + 2*p.x*p.y + p.y + p.y*p.y + designersnumber
	return strings.Count(fmt.Sprintf("%b", a), "1")%2 == 0
}

func (p position) getPossibleDestinations() (n []position) {
	moves := []position{position{1, 0}, position{0, 1},
		position{0, -1}, position{-1, 0}}
	for _, m := range moves {
		new := p.add(m)
		if new.isOpenSpace() {
			n = append(n, new)
		}
	}
	return
}

func (p position) moveTo(dest position, maxSteps int) (stepsTaken, posVisited int) {
	arrived := false
	currentPos := []position{p}
	visitedPos := map[position]bool{p: true}
	for !arrived && stepsTaken < maxSteps || maxSteps == 0 {
		stepsTaken++
		nextPos := []position{}
		for _, cp := range currentPos {
			for _, np := range cp.getPossibleDestinations() {
				if np == dest {
					return
				}
				if !visitedPos[np] {
					visitedPos[np] = true
					nextPos = append(nextPos, np)
				}
			}
		}
		posVisited = len(visitedPos)
		currentPos = nextPos
	}

	return
}

func main() {

	start := position{1, 1}
	dest := position{31, 39}

	steps, _ := start.moveTo(dest, 0)
	fmt.Printf("You can reach %v in %v steps.\n", dest, steps)

	_, visited := start.moveTo(dest, 50)
	fmt.Printf("With a max of 50 steps you can visit %v locations.\n", visited)
}
