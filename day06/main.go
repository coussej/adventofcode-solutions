package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type grid [1000][1000]int

type position struct {
	r, c int
}

type instruction struct {
	command  string
	from, to position
}

func (g *grid) runInstruction(instr instruction) {
	for r := instr.from.r; r <= instr.to.r; r++ {
		for c := instr.from.c; c <= instr.to.c; c++ {
			switch instr.command {
			case "toggle":
				if g[r][c] == 0 {
					g[r][c] = 1
				} else {
					g[r][c] = 0
				}
			case "on":
				g[r][c] = 1
			case "off":
				g[r][c] = 0
			}
		}
	}
}

func (g *grid) runInstructionBrightness(instr instruction) {
	for r := instr.from.r; r <= instr.to.r; r++ {
		for c := instr.from.c; c <= instr.to.c; c++ {
			switch instr.command {
			case "toggle":
				g[r][c] = g[r][c] + 2
			case "on":
				g[r][c] = g[r][c] + 1
			case "off":
				if g[r][c] != 0 {
					g[r][c] = g[r][c] - 1
				}
			}
		}
	}
}

func (g *grid) burningLights() int {
	var counter int
	for r, _ := range g {
		for _, l := range g[r] {
			if l > 0 {
				counter++
			}
		}
	}
	return counter
}

func (g *grid) totalBrightness() int {
	var total int
	for r, _ := range g {
		for _, l := range g[r] {
			total = total + l
		}
	}
	return total
}

func instrTextToPosition(in string) position {
	pos := strings.Split(in, ",")
	if len(pos) != 2 {
		panic("Invalid input.")
	}
	r, _ := strconv.Atoi(pos[0])
	c, _ := strconv.Atoi(pos[1])
	return position{r, c}
}

func parseInput(input string) []instruction {
	lines := strings.Split(input, "\n")
	var instr []instruction
	for _, v := range lines {
		words := strings.Split(v, " ")
		p1 := instrTextToPosition(words[len(words)-3])
		p2 := instrTextToPosition(words[len(words)-1])
		command := words[len(words)-4]
		instr = append(instr, instruction{command, p1, p2})
	}
	return instr
}

func main() {
	in, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	instr := parseInput(string(in))

	g := grid{}
	for _, i := range instr {
		g.runInstruction(i)
	}
	fmt.Println("After following Santa's instructions,", g.burningLights(), "lights are lit.")

	g = grid{}
	for _, i := range instr {
		g.runInstructionBrightness(i)
	}
	fmt.Println("After following Santa's new instructions, the total brightness is", g.totalBrightness())
}
