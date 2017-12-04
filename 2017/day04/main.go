package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	count1, count2 := countValidPassphrases(getPassphrases())
	fmt.Printf("The list contains %v passphases without duplicate words.\n", count1)
	fmt.Printf("It also contains %v entries that don't violate the anagram rule.\n", count2)
}

func countValidPassphrases(list [][]string) (method1, method2 int) {
	for _, p := range list {
		valid1, valid2 := true, true
		words, sortedWords := map[string]bool{}, map[string]bool{}

		for _, word := range p {
			// method 1
			if words[word] {
				valid1 = false
			}
			words[word] = true

			// method 2
			letters := strings.Split(word, "")
			sort.Strings(letters)
			sortedWord := strings.Join(letters, "")
			if sortedWords[sortedWord] {
				valid2 = false
			}
			sortedWords[sortedWord] = true
		}
		if valid1 {
			method1++
		}
		if valid2 {
			method2++
		}
	}
	return
}

func getPassphrases() (list [][]string) {
	in, _ := ioutil.ReadFile("input.txt")
	for _, row := range strings.Split(string(in), "\n") {
		if len(row) > 0 {
			r := strings.Split(row, " ")
			list = append(list, r)
		}
	}
	return
}
