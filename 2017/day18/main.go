package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type instruction struct {
	cmd, param1, param2 string
}

func main() {
	instr := getInput()

	// part 1

	snd, rcv := make(chan int, 1000), make(chan int, 1000)

	freq, blocked := 0, false
	go runProgram(0, instr, snd, rcv, false)

	for !blocked {
		select {
		case freq = <-snd:
		case <-time.After(3 * time.Second):
			fmt.Printf("Program seems blocked. The last frequency sent was %v.\n", freq)
			blocked = true
		}
	}

	// part 2

	chan1, chan2, intercept := make(chan int, 1000), make(chan int, 1000), make(chan int, 1000)

	go runProgram(0, instr, chan1, chan2, true)
	go runProgram(1, instr, intercept, chan1, true)

	count, blocked := 0, false
	for !blocked {
		select {
		case val := <-intercept:
			count++
			chan2 <- val
		case <-time.After(3 * time.Second):
			fmt.Printf("Programs seem blocked. PID 1 has sent a value %v times.\n", count)
			blocked = true
		}
	}
}

func runProgram(pid int, instr []instruction, snd, rcv chan int, multithreaded bool) {
	register := map[string]int{"p": pid}
	for i := 0; i < len(instr); i++ {
		cmd, param1, param2 := instr[i].cmd, instr[i].param1, instr[i].param2
		val1, err := strconv.Atoi(param1)
		if err != nil {
			val1 = register[param1]
		}
		val2, err := strconv.Atoi(param2)
		if err != nil {
			val2 = register[param2]
		}
		switch cmd {
		case "snd":
			snd <- val1
		case "set":
			register[param1] = val2
		case "add":
			register[param1] += val2
		case "mul":
			register[param1] *= val2
		case "mod":
			register[param1] %= val2
		case "rcv":
			if multithreaded || val1 > 0 {
				register[param1] = <-rcv
			}
		case "jgz":
			if val1 > 0 {
				i = i + val2 - 1
			}
		}
	}
}

func getInput() (instr []instruction) {
	in, _ := ioutil.ReadFile("input.txt")
	for _, line := range strings.Split(string(in), "\n") {
		if len(line) > 0 {
			words := strings.Split(line, " ")
			cmd, param1, param2 := words[0], words[1], ""
			if len(words) > 2 {
				param2 = words[2]
			}
			instr = append(instr, instruction{cmd, param1, param2})
		}
	}
	return
}
