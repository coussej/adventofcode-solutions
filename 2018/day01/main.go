package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	list := getList()

	sum := 0

	for _, n := range list {
		sum += n
	}

	fmt.Printf("The resulting frequency is %v\n", sum)

	sum = 0
	hist := map[int]bool{0: true}

search:
	for {
		for _, n := range list {
			sum += n
			if hist[sum] {
				break search
			}
			hist[sum] = true
		}
	}

	fmt.Printf("The first frequency seen twice is %v\n", sum)
}

func getList() (list []int) {
	in, _ := ioutil.ReadFile("input.txt")
	for _, c := range strings.Split(string(in), "\n") {
		num, err := strconv.Atoi(string(c))
		if err == nil {
			list = append(list, num)
		}
	}
	return
}
