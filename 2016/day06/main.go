package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func getMessages() (messages []string) {
	in, _ := ioutil.ReadFile("input.txt")
	in2 := strings.TrimSpace(regexp.MustCompile("\n *").ReplaceAllString(string(in), " "))
	return strings.Split(in2, " ")
}

func errorCorrenction(messages []string) (max, min string) {
	positions := make([]map[string]int, len(messages[0]), len(messages[0]))

	// fill with empty maps
	for i := range positions {
		positions[i] = make(map[string]int)
	}

	// count occurences
	for _, msg := range messages {
		for i, ltr := range msg {
			positions[i][string(ltr)]++
		}
	}

	// find max and min occurence
	for _, pos := range positions {
		maxcount, mincount := 0, 0
		maxltr, minltr := "", ""
		for ltr, count := range pos {
			if count > maxcount {
				maxcount, maxltr = count, ltr
			}
			if count < mincount || mincount == 0 {
				mincount, minltr = count, ltr
			}
		}
		max, min = max+maxltr, min+minltr
	}
	return
}

func main() {
	msgs := getMessages()
	max, min := errorCorrenction(msgs)
	fmt.Printf("The corrected message is %v with method 1, but %v with method 2.\n",
		max, min)
}
