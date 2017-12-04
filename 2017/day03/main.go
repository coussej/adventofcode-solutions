package main

import (
	"fmt"
)

func main() {
	input := 312051

	fmt.Printf("%v steps are required to carry the data from the input square to the access port.", getStepsToCenter(input))
	fmt.Printf("The first value written that is larger than the input is %v.", getFirstValueAbove(input))
}

func getStepsToCenter(pos int) int {
	side := 1
	// get first odd square higher then input
	for pos > side*side {
		side += 2
	}

	// calculate distances
	x1 := (side - 1) / 2
	x2 := x1 - (side*side-pos)%(side-1)
	if x2 < 0 {
		x2 = -x2
	}
	return x1 + x2
}

func getFirstValueAbove(limit int) (value int) {
	memory := map[loc]int{}
	rotation := 0
	curLoc := loc{0, 0}
	memory[curLoc] = 1

	for value < limit {
		fmt.Println(curLoc)
		switch {
		case curLoc == (loc{rotation, -rotation}):
			curLoc = curLoc.right()
			rotation++
		case curLoc.x == rotation && curLoc.y < rotation:
			curLoc = curLoc.up()
		case curLoc.y == rotation && curLoc.x > -rotation:
			curLoc = curLoc.left()
		case curLoc.x == -rotation && curLoc.y > -rotation:
			curLoc = curLoc.down()
		case curLoc.y == -rotation && curLoc.x < rotation:
			curLoc = curLoc.right()
		}
		value = curLoc.Value(memory)
		memory[curLoc] = value
	}
	return
}

type loc struct {
	x, y int
}

func (l *loc) up() loc {
	return loc{x: l.x, y: l.y + 1}
}

func (l *loc) down() loc {
	return loc{x: l.x, y: l.y - 1}
}

func (l *loc) left() loc {
	return loc{x: l.x - 1, y: l.y}
}

func (l *loc) right() loc {
	return loc{x: l.x + 1, y: l.y}
}

func (l *loc) Value(mem map[loc]int) (value int) {
	value += mem[loc{l.x - 1, l.y - 1}] + mem[loc{l.x + 1, l.y + 1}]
	value += mem[loc{l.x + 1, l.y - 1}] + mem[loc{l.x - 1, l.y + 1}]
	value += mem[loc{l.x, l.y - 1}] + mem[loc{l.x, l.y + 1}]
	value += mem[loc{l.x - 1, l.y}] + mem[loc{l.x + 1, l.y}]
	return
}
