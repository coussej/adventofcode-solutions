package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	in, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	var nice1, nice2 int
	for _, s := range strings.Split(string(in), "\n") {
		if countVowels(s) > 2 && hasDoubleLetter(s) && !containsForbiddenSequence(s) {
			nice1++
		}
		if containsPairTwice(s) && containsAlternatingRepeat(s) {
			nice2++
		}
	}
	fmt.Println("According to the first set of rules,", nice1, "strings are nice.")
	fmt.Println("According to the second set of rules,", nice2, "strings are nice.")
}

func isNice2(s string) bool {
	return false
}

func countVowels(s string) int {
	var vowels int
	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			vowels++
		}
	}
	return vowels
}

func hasDoubleLetter(s string) bool {
	chars := strings.Split(s, "")
	for i := 0; i < len(chars)-1; i++ {
		if chars[i] == chars[i+1] {
			return true
		}
	}
	return false
}

func containsForbiddenSequence(s string) bool {
	ix := strings.Index
	if ix(s, "ab")+ix(s, "cd")+ix(s, "pq")+ix(s, "xy") != -4 {
		return true
	}
	return false
}

func containsPairTwice(s string) bool {
	ix := strings.Index
	chars := strings.Split(s, "")
	for i := 0; i < len(chars)-2; i++ {
		if ix(s[:i], chars[i]+chars[i+1])+ix(s[i+2:], chars[i]+chars[i+1]) > -2 {
			return true
		}
	}
	return false
}

func containsAlternatingRepeat(s string) bool {
	chars := strings.Split(s, "")
	for i := 0; i < len(chars)-2; i++ {
		if chars[i] == chars[i+2] {
			return true
		}
	}
	return false
}
