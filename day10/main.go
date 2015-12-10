package main

import (
	"bytes"
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
	var current string
	var final bytes.Buffer
	var count int
	for i, _ := range s {
		switch {
		case count == 0:
			current = string(s[i])
			count++
		case i > 0 && s[i] == s[i-1]:
			count++
		default:
			final.WriteString(strconv.Itoa(count))
			final.WriteString(current)
			current = string(s[i])
			count = 1
		}
	}
	final.WriteString(strconv.Itoa(count))
	final.WriteString(current)
	return final.String()
}
