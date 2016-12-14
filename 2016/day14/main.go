package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func getFirstRepeatedChar(hash, char string, n int) string {
	switch char {
	case "":
		ix := -1
		for _, v := range []string{
			"0", "1", "2", "3", "4", "5", "6", "7",
			"8", "9", "a", "b", "c", "d", "e", "f"} {
			if i := strings.Index(hash, strings.Repeat(v, n)); i >= 0 && (i < ix || ix == -1) {
				ix = i
			}
		}
		if ix >= 0 {
			return string(hash[ix])
		}
	default:
		if strings.Contains(hash, strings.Repeat(char, n)) {
			return char
		}
	}
	return ""
}

func getHashHex(s string, stretches int) (hash string) {
	hashbyte := md5.Sum([]byte(s))
	hash = hex.EncodeToString(hashbyte[:])
	for i := 0; i < stretches; i++ {
		hashbyte = md5.Sum([]byte(hash))
		hash = hex.EncodeToString(hashbyte[:])
	}
	return
}

func findStepsToNrOfKeys(nKeys int, input string, stretchFactor int) (i int) {
	hashes, results := map[int]string{}, []string{}
	i = -1001

	for len(results) < nKeys {
		i++
		hashes[i+1000] = getHashHex(input+strconv.Itoa(i+1000), stretchFactor)
		if i < 0 {
			continue
		}
		c := getFirstRepeatedChar(hashes[i], "", 3)
		if c == "" {
			continue
		}
		for j := i + 1; j <= i+1000; j++ {
			if getFirstRepeatedChar(hashes[j], c, 5) != "" {
				results = append(results, hashes[i])
				break
			}
		}
	}
	return
}

func main() {
	input := "jlmsuwbz"

	fmt.Printf("It takes %v steps to generate 64 keys.\n",
		findStepsToNrOfKeys(64, input, 0))

	fmt.Printf("Using stretched hashes, it takes %v steps to generate 64 keys.\n",
		findStepsToNrOfKeys(64, input, 2016))
}
