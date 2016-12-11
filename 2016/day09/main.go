package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func getInput() string {
	in, _ := ioutil.ReadFile("input.txt")
	return strings.TrimSpace(regexp.MustCompile("\n *").ReplaceAllString(string(in), ""))
}

func getDecompressedLength(data string, recursive bool) (length int) {
	for i := 0; i < len(data); i++ {
		if data[i] != '(' {
			length++
			continue
		}
		cb := strings.IndexRune(data[i:], ')')
		marker := strings.Split(data[i+1:i+cb], "x")
		n, _ := strconv.Atoi(marker[0])
		r, _ := strconv.Atoi(marker[1])
		if recursive {
			length += r * getDecompressedLength(data[i+cb+1:i+cb+1+n], true)
		} else {
			length += r * n
		}
		i += cb + n
	}
	return
}

func main() {
	data := getInput()
	fmt.Printf("Using version 1, the decompressed length of the file is %v.\n",
		getDecompressedLength(data, false))
	fmt.Printf("Using version 2, the decompressed length of the file is %v.\n",
		getDecompressedLength(data, true))
}
