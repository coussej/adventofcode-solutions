package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getInstructions() (instr []string) {
	in, _ := ioutil.ReadFile("input.txt")
	for _, line := range strings.Split(string(in), "\n") {
		if line != "" {
			instr = append(instr, line)
		}
	}
	return
}

type registry map[string]int

func (r registry) runInstructions(instr []string, cancel chan bool, out chan int) {
	n := 0
	for n < len(instr) {
		select {
		case <-cancel:
			return
		default:
			if n < 0 {
				n = 0
			}
			words := strings.Split(instr[n], " ")
			switch words[0] {
			case "cpy":
				x, err := strconv.Atoi(words[1])
				if err != nil {
					x = r[words[1]]
				}
				_, err = strconv.Atoi(words[2])
				if err != nil {
					r[words[2]] = x
				} else {
					fmt.Printf("Encountered invalid instruction %v\n", instr[n])
				}
			case "inc":
				r[words[1]]++
			case "dec":
				r[words[1]]--
			case "jnz":
				x, err := strconv.Atoi(words[1])
				if err != nil {
					x = r[words[1]]
				}
				y, err := strconv.Atoi(words[2])
				if err != nil {
					y = r[words[2]]
				}
				if x != 0 {
					n += y - 1 // because n++ at the end
				}
			case "tgl":
				x, err := strconv.Atoi(words[1])
				if err != nil {
					x = r[words[1]]
				}
				if n+x < 0 || n+x >= len(instr) {
					break
				}
				tglInstr := instr[n+x]
				newInstr := ""
				switch len(strings.Split(tglInstr, " ")) {
				case 2:
					newInstr = "inc"
					if tglInstr[0:3] == "inc" {
						newInstr = "dec"
					}
				case 3:
					newInstr = "jnz"
					if tglInstr[0:3] == "jnz" {
						newInstr = "cpy"
					}
				}
				instr[n+x] = newInstr + tglInstr[3:]
			case "out":
				out <- r[words[1]]
			}
			n++
		}
	}
	return
}

func main() {
searchloop:
	for i := 0; ; i++ {
		r := registry{"a": i}
		cancel, output := make(chan bool), make(chan int, 1000)
		go r.runInstructions(getInstructions(), cancel, output)
		prevValue := 0
		// check if the first 10 values are alternating
		for j := 0; j < 10; j++ {
			value := <-output
			if j > 0 && value == prevValue {
				cancel <- true
				continue searchloop
			}
			prevValue = value
		}
		// looks like we found our initial value!
		cancel <- true
		fmt.Printf("You should initialize register a to %v.\n", i)
		break
	}
}
