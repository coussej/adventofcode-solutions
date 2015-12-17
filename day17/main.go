package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	combos := getPossibleCombos(getInput())

	// Part 1:
	var count int
	for _, c := range combos {
		if Sum(c) == 150 {
			count++
		}
	}
	fmt.Println(count, "combinations are possible for containing 150l")

	// Part 2:
	count = 0
	min := 0
	for _, c := range combos {
		if Sum(c) == 150 {
			switch {
			case len(c) < min || min == 0:
				count = 1
				min = len(c)
			case len(c) == min:
				count++
			}
		}
	}
	fmt.Println(count, "combinations are possible for containing 150l with the lowest number of containers,", min)
}

func getInput() []int {
	in, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	res := []int{}
	for _, line := range strings.Split(string(in), "\n") {
		n, _ := strconv.Atoi(line)
		res = append(res, n)
	}
	return res
}

func getPossibleCombos(vals []int) [][]int {
	var c [][]int
	for i, val := range vals {
		n := []int{val}
		c = append(c, n)
		for _, com := range getPossibleCombos(append([]int{}, vals[i+1:]...)) {
			//fmt.Println(com)
			c = append(c, append(n, com...))
		}
	}
	return c
}

func Sum(s []int) (sum int) {
	for _, v := range s {
		sum += v
	}
	return
}
