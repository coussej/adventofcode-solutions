package main

import (
	"fmt"
	"strconv"
	"strings"
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
	var current, final string
	var count int
	for i, _ := range s {
		switch {
		case count == 0:
			current = string(s[i])
			count++
		case i > 0 && s[i] == s[i-1]:
			count++
		default:
			final = strings.Join([]string{final, strconv.Itoa(count), current}, "")
			current = string(s[i])
			count = 1
		}
	}
	final = strings.Join([]string{final, strconv.Itoa(count), current}, "")
	return final
}
