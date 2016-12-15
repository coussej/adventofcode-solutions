package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type disc struct {
	total int
	start int
}

func (d disc) getPositionAt(time int) int {
	return (d.start + time) % d.total
}

func getDiscs() (discs []disc) {
	in, _ := ioutil.ReadFile("input.txt")
	for _, l := range strings.Split(string(in), "\n") {
		if l == "" {
			continue
		}
		words := strings.Split(strings.Replace(l, ".", " .", -1), " ")
		total, _ := strconv.Atoi(words[3])
		start, _ := strconv.Atoi(words[11])
		discs = append(discs, disc{total, start})
	}
	return
}

func getFirstDiskAlignment(discs []disc) int {
searchloop:
	for t := 0; ; t++ {
		for n, d := range discs {
			if d.getPositionAt(t+1+n) != 0 {
				continue searchloop
			}
		}
		return t
	}
}

func main() {
	discs := getDiscs()
	fmt.Printf("For the capsult to pass, you should press the button at t = %v\n",
		getFirstDiskAlignment(discs))
	discs = append(discs, disc{11, 0})
	fmt.Printf("After the new disc was added, you should press the button at t = %v\n",
		getFirstDiskAlignment(discs))
}
