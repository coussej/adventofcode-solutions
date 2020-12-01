package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	expenses := getList()
	e1, e2 := getTwoExpensesThatMatchTotal(2020, expenses)
	fmt.Printf("The result of multiplying the two entries that sum to 2020 is %v.\n", e1*e2)
	e1, e2, e3 := getThreeExpensesThatMatchTotal(2020, expenses)
	fmt.Printf("The result of multiplying the three entries that sum to 2020 is %v.\n", e1*e2*e3)
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

func getTwoExpensesThatMatchTotal(sum int, expenses []int) (e1, e2 int) {
	for _, e1 := range expenses {
		for _, e2 := range expenses {
			if e1 != e2 && e1+e2 == sum {
				return e1, e2
			}
		}
	}
	return 0, 0
}

func getThreeExpensesThatMatchTotal(sum int, expenses []int) (e1, e2, e3 int) {
	for _, e1 := range expenses {
		for _, e2 := range expenses {
			for _, e3 := range expenses {
				if e1+e2+e3 == sum {
					return e1, e2, e3
				}
			}
		}
	}
	return 0, 0, 0
}
