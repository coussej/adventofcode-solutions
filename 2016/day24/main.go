package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type duct struct {
	walls   map[position]bool
	relLocs map[int]position
}

type position struct {
	x, y int
}

func (p position) add(p2 position) (sum position) {
	return position{p.x + p2.x, p.y + p2.y}
}

func (d duct) getPossibleNeightboursForPosition(p position) (n []position) {
	moves := []position{position{1, 0}, position{0, 1},
		position{0, -1}, position{-1, 0}}
	for _, m := range moves {
		new := p.add(m)
		if !d.walls[new] {
			n = append(n, new)
		}
	}
	return
}

func getPossibleCombinations(nums []int) (combis [][]int) {
	for _, n := range nums {
		others := []int{}
		for _, o := range nums {
			if o != n {
				others = append(others, o)
			}
		}
		combi := []int{n}
		if len(others) > 0 {
			for _, pc := range getPossibleCombinations(others) {
				combis = append(combis, append(combi, pc...))
			}
		} else {
			combis = append(combis, []int{n})
		}
	}
	return
}

func getInput() (d duct) {
	d = duct{
		walls:   map[position]bool{},
		relLocs: map[int]position{},
	}
	in, _ := ioutil.ReadFile("input.txt")
	for y, line := range strings.Split(string(in), "\n") {
		for x, char := range strings.Split(line, "") {
			if char == "#" {
				d.walls[position{x, y}] = true
				continue
			}
			if num, err := strconv.Atoi(char); err == nil {
				d.relLocs[num] = position{x, y}
			}
		}
	}
	return
}

func (d duct) getShortestPath(from, to position) []position {
	status := map[position][]position{from: []position{}}
	visited := map[position]bool{from: true}
	for {
		nextStatus := map[position][]position{}
		for currentPos, path := range status {
			nextPos := d.getPossibleNeightboursForPosition(currentPos)
			for _, next := range nextPos {
				p := make([]position, len(path))
				copy(p, path)
				p = append(p, next)
				if next == to {
					return p
				}
				if !visited[next] {
					visited[next] = true
					nextStatus[next] = p
				}
			}
		}
		status = nextStatus
	}
}

func (d duct) getShortestTravel(returnToStart bool) (steps int) {
	start := d.relLocs[0]
	destinations := []int{}
	for k := range d.relLocs {
		if k != 0 {
			destinations = append(destinations, k)
		}
	}

	steps = 0
	pathlengths := make(chan int)
	paths := getPossibleCombinations(destinations)
	for i, path := range paths {
		fmt.Println("calculating path", i)
		path2 := make([]int, len(path))
		copy(path2, path)
		go func() {
			pathlength := 0
			for i := range path2 {
				from := start
				if i > 0 {
					from = d.relLocs[path2[i-1]]
				}
				to := d.relLocs[path2[i]]
				pathlength += len(d.getShortestPath(from, to))
			}
			if returnToStart {
				pathlength += len(d.getShortestPath(d.relLocs[path2[len(path2)-1]], start))
			}
			pathlengths <- pathlength
		}()
	}
	for _ = range paths {
		l := <-pathlengths
		if l < steps || steps == 0 {
			steps = l
		}
	}
	return
}

func main() {
	d := getInput()
	fmt.Printf("The fewest number of steps required to visit all positions is %v.\n",
		d.getShortestTravel(false))
	fmt.Printf("If you must return the robot to start, the fewest number of steps required to visit all positions is %v.\n",
		d.getShortestTravel(true))
}
