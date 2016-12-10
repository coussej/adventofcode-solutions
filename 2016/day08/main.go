package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type screen [][]int

func newScreen(width, height int) (s screen) {
	for r := 0; r < height; r++ {
		s = append(s, []int{})
		for c := 0; c < width; c++ {
			s[r] = append(s[r], 0)
		}
	}
	return
}

func (s screen) rect(x, y int) {
	for r := 0; r < y; r++ {
		for c := 0; c < x; c++ {
			s[r][c] = 1
		}
	}
}

func (s screen) rotateRow(y, n int) {
	new := []int{}
	for i := 0; i < len(s[y]); i++ {
		new = append(new, s[y][((i-n+len(s[y]))%len(s[y]))])
	}
	s[y] = new
}

func (s screen) rotateCol(x, n int) {
	new := []int{}
	for i := 0; i < len(s); i++ {
		new = append(new, s[((i - n + len(s)) % len(s))][x])
	}
	for i := 0; i < len(s); i++ {
		s[i][x] = new[i]
	}
}

func (s screen) voltage() int {
	v := 0
	for r := 0; r < len(s); r++ {
		for c := 0; c < len(s[r]); c++ {
			v += s[r][c]
		}
	}
	return v
}

func (s screen) print() {
	for _, r := range s {
		for i, c := range r {
			if i%5 == 0 {
				fmt.Print("   ")
			}
			if c == 1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

func getInstructions() (instructions []string) {
	in, _ := ioutil.ReadFile("input.txt")
	in2 := strings.TrimSpace(regexp.MustCompile("\n *").ReplaceAllString(string(in), ";"))
	return strings.Split(in2, ";")
}

func main() {
	s := newScreen(50, 6)
	for _, instr := range getInstructions() {
		switch {
		case strings.Contains(instr, "rect"):
			xy := strings.Split(strings.Replace(instr, "rect ", "", -1), "x")
			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])
			s.rect(x, y)
		case strings.Contains(instr, "rotate row"):
			yn := strings.Split(strings.Replace(instr, "rotate row y=", "", -1), " by ")
			y, _ := strconv.Atoi(yn[0])
			n, _ := strconv.Atoi(yn[1])
			s.rotateRow(y, n)
		case strings.Contains(instr, "rotate column"):
			xn := strings.Split(strings.Replace(instr, "rotate column x=", "", -1), " by ")
			x, _ := strconv.Atoi(xn[0])
			n, _ := strconv.Atoi(xn[1])
			s.rotateCol(x, n)
		}
	}
	fmt.Println("The screen has", s.voltage(), "pixels lit.")
	fmt.Println("The code it is trying to display is:")
	s.print()
}
