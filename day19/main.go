package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	r, m := getInput()
	fmt.Println("Possible distinct molecules:", countPossibleNextMolecules(r, m))
	fmt.Println("Steps from electron to molecule:", countStepsToElectron(m))
}

func countStepsToElectron(molecule string) int {
	//  Only possible transitions are (in reverse, with X not in {Rn, Ar, Y}):
	//    XX -> X          : Reduction by 1 = 1
	//    XX -> e          : Reduction by 1 = 1
	//    XRnXAr -> X      : Reduction by 3 = 1 + count(Rn|Ar)
	//    XRnXYXAr -> X    : Reduction by 5 = 1 + count(Rn|Ar) + 2*count(Y)
	//    XRnXYXYXAr -> X  : Reduction by 7 = 1 + count(Rn|Ar) + 2*count(Y)
	//  This means we can find the steps required by counting the occurences
	//  of (Rn|Ar) and (Y), en filling them in the formula.
	regAtom := regexp.MustCompile("([A-Z]{1}[a-z]*)")
	regRnAr := regexp.MustCompile("Rn|Ar")
	regY := regexp.MustCompile("Y")

	cAtom := len(regAtom.FindAllString(molecule, -1))
	cRnAr := len(regRnAr.FindAllString(molecule, -1))
	cY := len(regY.FindAllString(molecule, -1))
	return cAtom - cRnAr - 2*cY - 1
}

func countPossibleNextMolecules(rules map[string]string, molecule string) int {
	// Go over the rules and create the molecules by replacing each occurence once
	// and adding them to a map[string]bool. By storing the molecules as map keys,
	// we won't have any duplicates.
	pm := map[string]bool{}
	for to, from := range rules {
		i := strings.Index(molecule, from)
		for i >= 0 {
			mol := molecule[:i] + strings.Replace(molecule[i:], from, to, 1)
			pm[mol] = true
			if next := strings.Index(molecule[i+1:], from); next >= 0 {
				i += next + 1
			} else {
				break
			}
		}
	}
	return len(pm)
}

func getInput() (r map[string]string, m string) {
	in, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	// Rules and input molecule are separated by blank line.
	inputs := strings.Split(string(in), "\n\n")

	// Rules
	r = map[string]string{}
	for _, line := range strings.Split(inputs[0], "\n") {
		s := strings.Split(line, " ")
		r[s[2]] = s[0]
	}

	// Input molecule
	m = inputs[1]

	return
}
