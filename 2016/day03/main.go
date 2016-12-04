package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type triangle struct {
	a, b, c int
}

func (t *triangle) isPossible() bool {
	return t.a+t.b > t.c &&
		t.a+t.c > t.b &&
		t.b+t.c > t.a
}

func getTriangles(offset int) (triangles []triangle) {
	in, _ := ioutil.ReadFile("input.txt")
	// put everything on one line and split
	in2 := regexp.MustCompile("\n *").ReplaceAll(in, []byte(" "))
	sides := regexp.MustCompile(" +").Split(strings.TrimSpace(string(in2)), -1)

	for i := 0; i < len(sides)-3; i++ {
		a, _ := strconv.Atoi(sides[i])
		b, _ := strconv.Atoi(sides[i+offset])
		c, _ := strconv.Atoi(sides[i+2*offset])
		triangles = append(triangles, triangle{a, b, c})

		// take a jump when a multiple of the offset is reached
		if (i+1)%offset == 0 {
			i += 2 * offset
		}
	}
	return
}

func main() {
	count := 0
	for _, t := range getTriangles(1) {
		if t.isPossible() {
			count++
		}
	}
	fmt.Println("There are", count, "possible triangles when reading by rows.")
	count = 0
	for _, t := range getTriangles(3) {
		if t.isPossible() {
			count++
		}
	}
	fmt.Println("There are", count, "possible triangles when reading by columns.")
}
