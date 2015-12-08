package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	in, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// regex for hex escapes
	hexEscapeRegEx := regexp.MustCompile(`[\\][x][a-fA-F0-9]{2}`)
	// regex for \" and \\
	normalExcapeRegEx := regexp.MustCompile(`[\\][\\"]`)
	// regex for " at begining and end of string. Not really needed, as there are always two per line.
	quotesRegEx := regexp.MustCompile(`(^["]|["]$)`)

	var totalDiff, totalExtraCharsNewEncoding int
	for _, l := range strings.Split(string(in), "\n") {
		s := len(l)
		h := len(hexEscapeRegEx.FindAllString(l, -1))
		n := len(normalExcapeRegEx.FindAllString(l, -1))
		q := len(quotesRegEx.FindAllString(l, -1))

		// part 1
		totalDiff += s - (s - 3*h - n - q)

		// part 2:
		//  - hex escapes have one extra char, as \xaa becomes \\xaa
		//  - normal escapes have 2 extra chars, as \\ and \" become \\\\ and \\\"
		//  - begin and end quotes have 2 extra chars, as ^" and "$ becomes "\" and \""
		totalExtraCharsNewEncoding += 1*h + 2*n + 2*q
	}

	fmt.Println("The number of characters of code for string literals minus the number of characters in memory for the values of the strings in total for the entire file in", totalDiff)
	fmt.Println("the total number of characters to represent the newly encoded strings minus the number of characters of code in each original string literal.", totalExtraCharsNewEncoding)
}
