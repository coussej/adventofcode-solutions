package main

import (
	"fmt"
	"io/ioutil"
)

type coordinate struct {
	x, y int
}

func (c *coordinate) move(r rune) {
	switch r {
	case '^':
		c.y++
	case 'v':
		c.y--
	case '<':
		c.x++
	case '>':
		c.x--
	}
}

func main() {
	in, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	// create map of houses
	houses := make(map[coordinate]int)
	houses2 := make(map[coordinate]int)
	// santa starting position
	santa := coordinate{0, 0}
	santa2 := coordinate{0, 0}
	robosanta := coordinate{0, 0}
	// give present to first house
	houses[santa]++
	houses2[santa2]++
	houses2[robosanta]++
	// travel
	for i, v := range string(in) {
		santa.move(v)
		houses[santa]++
		if i%2 == 0 {
			santa2.move(v)
			houses2[santa2]++
		} else {
			robosanta.move(v)
			houses2[robosanta]++
		}
	}
	fmt.Println("Santa delivers presents in", len(houses), "houses")
	fmt.Println("When Robo-Santa helps Santa, they deliver presents in", len(houses2), "houses")
}
