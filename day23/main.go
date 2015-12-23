package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type instruction struct {
	instr  string
	reg    string
	offset int
}

func main() {
	instr := getInstructions()

	a, b := runInstructions(0, 0, instr)
	fmt.Println("Result of program with (a0,b0) = (0,0): a =", a, "| b =", b)

	a, b = runInstructions(1, 0, instr)
	fmt.Println("Result of program with (a0,b0) = (1,0): a =", a, "| b =", b)

}
func runInstructions(a0, b0 int, instr []instruction) (a, b int) {
	a = a0
	b = b0
	for i := 0; i >= 0 && i < len(instr); {
		switch instr[i].instr {
		case "hlf":
			switch instr[i].reg {
			case "a":
				a /= 2
			case "b":
				b /= 2
			}
			i++
		case "tpl":
			switch instr[i].reg {
			case "a":
				a *= 3
			case "b":
				b *= 3
			}
			i++
		case "inc":
			switch instr[i].reg {
			case "a":
				a++
			case "b":
				b++
			}
			i++
		case "jmp":
			i += instr[i].offset
		case "jie":
			switch instr[i].reg {
			case "a":
				if a%2 == 0 {
					i += instr[i].offset
				} else {
					i++
				}
			case "b":
				if b%2 == 0 {
					i += instr[i].offset
				} else {
					i++
				}
			}
		case "jio":
			switch instr[i].reg {
			case "a":
				if a == 1 {
					i += instr[i].offset
				} else {
					i++
				}
			case "b":
				if b == 1 {
					i += instr[i].offset
				} else {
					i++
				}
			}
		}
	}
	return
}

func getInstructions() (instructions []instruction) {
	in, _ := ioutil.ReadFile("input.txt")
	for _, line := range strings.Split(string(in), "\n") {
		words := strings.Split(line, " ")
		switch words[0] {
		case "hlf", "tpl", "inc":
			instructions = append(instructions, instruction{words[0], words[1], 0})
		case "jmp":
			off, _ := strconv.Atoi(words[1])
			instructions = append(instructions, instruction{words[0], "", off})
		case "jie", "jio":
			off, _ := strconv.Atoi(words[2])
			instructions = append(instructions, instruction{words[0], words[1][:1], off})
		default:
			panic("invalid instruction")
		}
	}
	return
}
