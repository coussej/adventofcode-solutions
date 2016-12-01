package main

import (
	"bytes"
	"fmt"
)

func main() {
	input := "hepxcrrq"
	pw := nextPassWord(input)
	fmt.Println("Santa's next password is", pw)
	fmt.Println("The one after that is", nextPassWord(pw))

}

func nextPassWord(i string) string {
	in := []byte(i)
	for {
		in = nextValue(in)
		if containsStraightOfThree(in) &&
			containsTwoNonOverlappingPairs(in) &&
			!containsForbiddenLetter(in, []byte("iol")) {
			break
		}
	}
	return string(in)
}

func nextValue(v []byte) []byte {
	res := v
	for i := len(res) - 1; i >= 0; i-- {
		res[i]++
		if res[i] > []byte("z")[0] {
			res[i] = []byte("a")[0]
		} else {
			break
		}
	}
	return res
}

func containsStraightOfThree(v []byte) bool {
	for i := 0; i < len(v)-2; i++ {
		if v[i+2] == v[i+1]+1 && v[i+1] == v[i]+1 {
			return true
		}
	}
	return false
}

func containsTwoNonOverlappingPairs(v []byte) bool {
	var count int
	for i := 0; i < len(v)-1; i++ {
		if v[i+1] == v[i] {
			count++
			i++ // skip next letter
		}
	}
	if count > 1 {
		return true
	}
	return false
}

func containsForbiddenLetter(v []byte, forbidden []byte) bool {
	for _, b := range forbidden {
		if bytes.Contains(v, []byte{b}) {
			return true
		}
	}
	return false
}
