package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// assume grid is always 100x100, simplifies code
type grid [100][100]bool
type animation []grid

func main() {
	g := getInput()

	a := playAnimation(g, 100, false)
	fmt.Println("The total number of lights after 100 frames is", a[len(a)-1].totalLightsOn())

	a = playAnimation(g, 100, true)
	fmt.Println("The total number of lights with fixed corners after 100 frames is", a[len(a)-1].totalLightsOn())
}

func getInput() grid {
	in, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	res := grid{}
	for r, line := range strings.Split(string(in), "\n") {
		for c, s := range line {
			switch s {
			case '#':
				res[r][c] = true
			case '.':
				res[r][c] = false
			}
		}
	}
	return res
}

func playAnimation(g grid, n int, fixedCorners bool) (a animation) {
	a = animation{}
	if fixedCorners {
		g.cornersOn()
	}
	a = append(a, g)
	for i := 1; i <= n; i++ {
		gn := grid{}
		for r := 0; r <= 99; r++ {
			for c := 0; c <= 99; c++ {
				nb := a[i-1].getSumOfNeighbours(r, c)
				switch a[i-1][r][c] {
				case true:
					if nb == 2 || nb == 3 {
						gn[r][c] = true
					} else {
						gn[r][c] = false
					}
				case false:
					if nb == 3 {
						gn[r][c] = true
					} else {
						gn[r][c] = false
					}
				}
			}
		}
		if fixedCorners {
			gn.cornersOn()
		}
		a = append(a, gn)
	}
	return
}

func (g *grid) cornersOn() {
	g[0][0] = true
	g[0][99] = true
	g[99][0] = true
	g[99][99] = true
}

func (g grid) getSumOfNeighbours(r, c int) (sum int) {
	for i := max(0, r-1); i <= min(99, r+1); i++ {
		for j := max(0, c-1); j <= min(99, c+1); j++ {
			if !(i == r && j == c) {
				if g[i][j] {
					sum++
				}
			}
		}
	}
	return
}

func (g grid) totalLightsOn() (total int) {
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			if g[i][j] {
				total++
			}
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
