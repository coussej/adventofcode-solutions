package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

type claim struct {
	id, left, top, width, height int
}

func main() {
	claims := getClaims()
	fabric := newFabric(1000)

	overlapping := 0

	for _, c := range claims {
		for x := c.left; x < c.left+c.width; x++ {
			for y := c.top; y < c.top+c.height; y++ {
				if fabric[x][y] == 1 {
					overlapping++
				}
				fabric[x][y]++
			}
		}
	}

	fmt.Printf("There are %v overlapping square inches.\n", overlapping)

	for _, c := range claims {
		overlap := false
		for x := c.left; x < c.left+c.width; x++ {
			for y := c.top; y < c.top+c.height; y++ {
				if fabric[x][y] > 1 {
					overlap = true
				}
			}
		}
		if !overlap {
			fmt.Printf("The claim without overlaps has id %v.\n", c.id)
			break
		}
	}

}

func getClaims() (claims []claim) {
	in, _ := ioutil.ReadFile("input.txt")
	pat := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)
	matches := pat.FindAllStringSubmatch(string(in), -1)
	for _, m := range matches {
		claims = append(claims, claim{
			id:     toInt(m[1]),
			left:   toInt(m[2]),
			top:    toInt(m[3]),
			width:  toInt(m[4]),
			height: toInt(m[5]),
		})
	}
	return
}

func newFabric(size int) (fabric [][]int) {
	fabric = make([][]int, size, size)
	for i := range fabric {
		fabric[i] = make([]int, size, size)
	}
	return
}

func toInt(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return
}
