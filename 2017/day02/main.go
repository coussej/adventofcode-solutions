package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type spreadsheet [][]int

func main() {
	s := getSpreadsheet()
	fmt.Printf("The checksum is %v.\n", s.getCheckSum())
	fmt.Printf("The new checksum calculation results in %v.\n", s.getCheckSum2())
}

func (s *spreadsheet) getCheckSum() (checksum int) {
	for _, r := range *s {
		min, max := 0, 0
		for _, c := range r {
			if c < min || min == 0 {
				min = c
			}
			if c > max {
				max = c
			}
		}
		checksum += max - min
	}
	return
}

func (s *spreadsheet) getCheckSum2() (checksum int) {
rows:
	for _, row := range *s {
		r := make([]int, len(row))
		copy(r, row)
		sort.Ints(r)
		for i := len(r) - 1; i >= 0; i-- {
			for j := 0; j < i; j++ {
				if r[i]%r[j] == 0 {
					checksum += r[i] / r[j]
					continue rows
				}
			}
		}
	}
	return
}

func getSpreadsheet() (s spreadsheet) {
	in, _ := ioutil.ReadFile("input.txt")
	for _, row := range strings.Split(string(in), "\n") {
		if len(row) > 0 {
			r := []int{}
			for _, cell := range strings.Split(row, "\t") {
				num, _ := strconv.Atoi(cell)
				r = append(r, num)
			}
			s = append(s, r)
		}
	}
	return
}
