package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type present struct {
	l, w, h int
}

func (p *present) area() int {
	return 2*p.l*p.w + 2*p.l*p.h + 2*p.w*p.h
}

func (p *present) areaSmallestSide() int {
	a := []int{p.l * p.w, p.l * p.h, p.w * p.h}
	return smallestIntFromSlice(a)
}

func (p *present) smallestPerimeter() int {
	a := []int{2*p.l + 2*p.w, 2*p.l + 2*p.h, 2*p.w + 2*p.h}
	return smallestIntFromSlice(a)
}

func (p *present) volume() int {
	return p.l * p.w * p.h
}

func smallestIntFromSlice(i []int) int {
	min := i[0]
	for _, v := range i[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func main() {
	in, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	var paper, ribbon int
	for _, v := range strings.Split(string(in), "\n") {
		s := strings.Split(v, "x")
		l, _ := strconv.Atoi(s[0])
		w, _ := strconv.Atoi(s[1])
		h, _ := strconv.Atoi(s[2])
		p := present{l, w, h}
		paper = paper + p.area() + p.areaSmallestSide()
		ribbon = ribbon + p.smallestPerimeter() + p.volume()
	}
	fmt.Println("The elves need to order", paper, "sq ft of paper.")
	fmt.Println("The elves need to order", ribbon, "ft of ribbon.")
}
