package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	score, garbage := getScore(getInput())
	fmt.Printf("The total score of all groups in the input is %v.\n", score)
	fmt.Printf("There are %v non-canceled characters in the garbage.\n", garbage)
}

func getScore(stream string) (score, garbage int) {
	groupLevel, garbageOpen := 0, false
	for i := 0; i < len(stream); i++ {
		c := stream[i]
		switch {
		case c == '!':
			i++
		case garbageOpen && c == '>':
			garbageOpen = false
		case garbageOpen:
			garbage++
		case !garbageOpen && c == '<':
			garbageOpen = true
		case !garbageOpen && c == '{':
			groupLevel++
			score += groupLevel
		case !garbageOpen && c == '}' && groupLevel > 0:
			groupLevel--
		}
	}
	return
}

func getInput() (stream string) {
	in, _ := ioutil.ReadFile("input.txt")
	return strings.Split(string(in), "\n")[0]
}
