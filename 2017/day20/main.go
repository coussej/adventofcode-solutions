package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type particle struct {
	X, Y, Z, vX, vY, vZ, aX, aY, aZ int
}

type system []particle

func main() {

	// part 1
	s1 := getInput()
	s1.simulate(1000)
	fmt.Printf("Particle %v will stay closest to the origin in the long term.\n", s1.closestToOrigin())

	// part 2
	s2 := getInput()
	for i := 0; i < 1000; i++ {
		s2.simulate(1)
		s2.removeCollisions()
	}
	fmt.Printf("There are %v particles left after collisions are resolved.\n", len(s2))
}

func (p *particle) accelerate() {
	p.vX += p.aX
	p.vY += p.aY
	p.vZ += p.aZ
}

func (p *particle) move() {
	p.X += p.vX
	p.Y += p.vY
	p.Z += p.vZ
}

func (p particle) hasSamePositionAs(p2 particle) bool {
	return p.X == p2.X && p.Y == p2.Y && p.Z == p2.Z
}

func (p *particle) distanceToOrigin() int {
	return int(math.Abs(float64(p.X)) + math.Abs(float64(p.Y)) + math.Abs(float64(p.Z)))
}

func (s system) simulate(steps int) {
	for i := 0; i < steps; i++ {
		for id := range s {
			s[id].accelerate()
			s[id].move()
		}
	}
}

func (s system) closestToOrigin() int {
	min, id := math.MaxInt64, 0
	for i, p := range s {
		if d := p.distanceToOrigin(); d < min {
			min, id = d, i
		}
	}
	return id
}

func (s *system) removeCollisions() {
	for j := 0; j < len(*s); {
		collisionFound := false
		for k := len(*s) - 1; k > j; k-- {
			if (*s)[k].hasSamePositionAs((*s)[j]) {
				collisionFound = true
				*s = append((*s)[:k], (*s)[k+1:]...)
			}
		}
		if collisionFound {
			*s = append((*s)[:j], (*s)[j+1:]...)
		} else {
			j++
		}
	}
}

func getInput() (s system) {
	in, _ := ioutil.ReadFile("input.txt")
	for _, line := range strings.Split(string(in), "\n") {
		if len(line) > 0 {
			parts := regexp.MustCompile("([-0-9]+)[>,]").FindAllStringSubmatch(line, -1)
			p := particle{}
			p.X, _ = strconv.Atoi(parts[0][1])
			p.Y, _ = strconv.Atoi(parts[1][1])
			p.Z, _ = strconv.Atoi(parts[2][1])
			p.vX, _ = strconv.Atoi(parts[3][1])
			p.vY, _ = strconv.Atoi(parts[4][1])
			p.vZ, _ = strconv.Atoi(parts[5][1])
			p.aX, _ = strconv.Atoi(parts[6][1])
			p.aY, _ = strconv.Atoi(parts[7][1])
			p.aZ, _ = strconv.Atoi(parts[8][1])
			s = append(s, p)
		}
	}
	return
}
