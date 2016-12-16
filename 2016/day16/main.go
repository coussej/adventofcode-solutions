package main

import (
	"fmt"
	"math"
	"strconv"
)

func fillDisk(data []int, disklenght int) []int {
	a := data
	b := a
	for len(b) < disklenght {
		b = append(b, 0)
		for i := len(a) - 1; i >= 0; i-- {
			b = append(b, int(math.Abs(float64(a[i]-1))))
		}
		a = b
	}
	return b[:disklenght]
}

func getChecksum(data []int) (cs string) {
	for len(data)%2 == 0 {
		new := []int{}
		for i := 0; i < len(data); i = i + 2 {
			if data[i] == data[i+1] {
				new = append(new, 1)
			} else {
				new = append(new, 0)
			}
		}
		data = new
	}
	for _, c := range data {
		cs = cs + strconv.Itoa(c)
	}
	return
}

func main() {
	input := "10111100110001111"
	data := []int{}
	for _, n := range input {
		num, _ := strconv.Atoi(string(n))
		data = append(data, num)
	}
	fmt.Println("The checksum of the disk with lenght 272 is",
		getChecksum(fillDisk(data, 272)))
	fmt.Println("The checksum of the disk with lenght 35651584 is",
		getChecksum(fillDisk(data, 35651584)))
}
