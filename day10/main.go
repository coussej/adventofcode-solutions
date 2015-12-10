package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := "1113222113"
	result := input
	for i := 0; i < 40; i++ {
		result = lookAndSay(result)
	}
	fmt.Println("After 40 iterations, the length of the result is", len(result))
	for i := 40; i < 50; i++ {
		result = lookAndSay(result)
	}
	fmt.Println("After 50 iterations, the length of the result is", len(result))
}

func lookAndSay(s string) string {
	var currSet, res string
	for i, _ := range s {
		switch {
		case len(currSet) == 0:
			currSet = string(s[i])
		case i > 0 && s[i] == s[i-1]:
			currSet = currSet + string(s[i])
		default:
			res = res + strconv.Itoa(len(currSet)) + string(currSet[0])
			currSet = string(s[i])
		}
	}
	res = res + strconv.Itoa(len(currSet)) + string(currSet[0])
	return res
}
