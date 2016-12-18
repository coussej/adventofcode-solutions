package main

import (
	"fmt"
	"io/ioutil"
)

func getFirstRow() (row []int) {
	in, _ := ioutil.ReadFile("input.txt")
	for _, char := range string(in) {
		switch char {
		case '^':
			row = append(row, 0) // trap
		case '.':
			row = append(row, 1) // safe
		}
	}
	return
}

func getNextRow(cur []int) (next []int) {
	for i := range cur {
		left, center, right := 1, 1, 1
		if i > 0 {
			left = cur[i-1]
		}
		center = cur[i]
		if i < len(cur)-1 {
			right = cur[i+1]
		}
		if (left == 0 && center == 0 && right == 1) ||
			(left == 1 && center == 0 && right == 0) ||
			(left == 0 && center == 1 && right == 1) ||
			(left == 1 && center == 1 && right == 0) {
			next = append(next, 0)
		} else {
			next = append(next, 1)
		}
	}
	return
}

func getTotalSafeTiles(room [][]int) (safe int) {
	for r := range room {
		for c := range room[r] {
			safe += room[r][c]
		}
	}
	return
}

func main() {
	room := [][]int{getFirstRow()}
	for i := 1; i < 400000; i++ {
		room = append(room, getNextRow(room[i-1]))
		if i == 39 {
			fmt.Printf("A room with 40 rows has %v safe tiles.\n",
				getTotalSafeTiles(room))
		}
	}
	fmt.Printf("A room with 400000 rows has %v safe tiles.\n",
		getTotalSafeTiles(room))
}
