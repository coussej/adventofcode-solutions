package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type grid [][]string

type book map[string]string

func main() {
	g := newGridFromPattern(".#./..#/###")
	rules := getInput()

	// part 1

	for i := 0; i < 5; i++ {
		g = g.enhance(rules)
	}
	fmt.Printf("There are %v pixels active after 5 iterations.\n", g.activePixels())

	// part 2

	for i := 5; i < 18; i++ {
		g = g.enhance(rules)
	}
	fmt.Printf("There are %v pixels active after 5 iterations.\n", g.activePixels())

}

func (g grid) enhance(rules book) (enh grid) {
	size := len(g)
	resolution := 2 + size%2
	enh = newGrid(size + size/resolution)
	for r := 0; r < size; r += resolution {
		for c := 0; c < size; c += resolution {
			// get the new pattern and loop over the elements, while adding them to the enhanced grid
			result := newGridFromPattern(rules[g.toPattern(r, c, resolution)])
			for i := range result {
				for j := range result[i] {
					enh[r+r/resolution+i][c+c/resolution+j] = result[i][j]
				}
			}
		}
	}
	return
}

func (g grid) rotate() (rot grid) {
	s := len(g)
	rot = newGrid(s)
	for r := range g {
		for c := range g {
			rot[c][s-1-r] = g[r][c]
		}
	}
	return
}

func (g grid) flip() (fl grid) {
	s := len(g)
	fl = newGrid(s)
	for r := range g {
		for c := range g {
			fl[r][c] = g[s-1-r][c]
		}
	}
	return
}

func (g grid) toPattern(row, column, size int) string {
	parts := []string{}
	for r := row; r < row+size; r++ {
		part := ""
		for c := column; c < column+size; c++ {
			part += g[r][c]
		}
		parts = append(parts, part)
	}
	return strings.Join(parts, "/")
}

func (g grid) activePixels() (count int) {
	for r := range g {
		for c := range g {
			if g[r][c] == "#" {
				count++
			}
		}
	}
	return
}

func newGridFromPattern(pattern string) (g grid) {
	for _, line := range strings.Split(pattern, "/") {
		g = append(g, strings.Split(line, ""))
	}
	return
}

func newGrid(size int) (g grid) {
	for r := 0; r < size; r++ {
		row := make([]string, size)
		g = append(g, row)
	}
	return
}

func getInput() (b book) {
	in, _ := ioutil.ReadFile("input.txt")
	b = book{}
	for _, line := range strings.Split(string(in), "\n") {
		if len(line) > 0 {
			parts := strings.Split(line, " ")
			g, result := newGridFromPattern(parts[0]), parts[2]
			for i := 0; i < 4; i++ {
				g = g.rotate()
				b[g.toPattern(0, 0, len(g))] = result
				b[g.flip().toPattern(0, 0, len(g))] = result
			}
		}
	}
	return
}
