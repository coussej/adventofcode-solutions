package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

func main() {
	polymer := getPolymer()
	fmt.Printf("There are %v remaining units after fully reacting the polymer.\n", len(reactPolymer(polymer)))

	shortest := len(polymer)
	for r := rune('a'); r <= rune('z'); r++ {
		pol, i := getPolymer(), 0
		for i < len(pol) {
			if pol[i] == r || pol[i] == invertCap(r) {
				pol = append(pol[:i], pol[i+1:]...)
			} else {
				i++
			}
		}
		finalLen := len(reactPolymer(pol))
		if finalLen < shortest {
			shortest = finalLen
		}
	}
	fmt.Printf("The length of the shortest polymer you can produce is %v.\n", shortest)

}

func reactPolymer(polymer []rune) (result []rune) {

	result = make([]rune, len(polymer), len(polymer))
	copy(result, polymer)

	cur := 0
	for cur < len(result)-1 {
		//fmt.Println(string(polymer))
		if result[cur] == invertCap(result[cur+1]) {
			result = append(result[:cur], result[cur+2:]...)
			cur--
		} else {
			cur++
		}
		if cur < 0 {
			cur = 0
		}
	}
	return
}

func invertCap(r rune) rune {
	if unicode.IsLower(r) {
		return unicode.ToUpper(r)
	}
	return unicode.ToLower(r)
}

func getPolymer() (polymer []rune) {
	in, _ := ioutil.ReadFile("input.txt")
	for _, el := range strings.Split(string(in), "\n")[0] {
		polymer = append(polymer, el)
	}
	return
}
