package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	list := getList()
	listDouble := append(list, list...)

	var sum1, sum2 int
	for i := 0; i < len(list); i++ {
		if listDouble[i+1] == listDouble[i] {
			sum1 += listDouble[i]
		}
		if listDouble[i+len(list)/2] == listDouble[i] {
			sum2 += listDouble[i]
		}
	}

	fmt.Printf("The solution to the first captcha is %v\n", sum1)
	fmt.Printf("The solution to the second captcha is %v\n", sum2)
}

func getList() (list []int) {
	in, _ := ioutil.ReadFile("input.txt")
	for _, c := range strings.Replace(string(in), "\n", "", -1) {
		num, _ := strconv.Atoi(string(c))
		list = append(list, num)
	}
	return
}
