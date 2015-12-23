package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	in, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	var floor, firstTimeInBasement int
	for i, v := range string(in) {
		if v == '(' {
			floor++
		} else {
			floor--
		}
		if floor < 0 && firstTimeInBasement == 0 {
			firstTimeInBasement = i + 1
		}
	}
	fmt.Println("Santa ends up on the", floor, "th floor.")
	fmt.Println("The first time he hits the basement level is after", firstTimeInBasement, "instructions.")
}
