package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type scrambler struct {
	instructions []string
}

func (scr scrambler) scramble(in string) (out string) {
	out = in
	for _, instr := range scr.instructions {
		words := strings.Split(instr, " ")
		switch words[0] {
		case "swap":
			switch words[1] {
			case "position":
				pos1, _ := strconv.Atoi(words[2])
				pos2, _ := strconv.Atoi(words[5])
				out = swapPositions(out, pos1, pos2)
			case "letter":
				out = swapLetters(out, words[2], words[5])
			}
		case "rotate":
			switch words[1] {
			case "left":
				pos, _ := strconv.Atoi(words[2])
				out = rotateLeft(out, pos)
			case "right":
				pos, _ := strconv.Atoi(words[2])
				out = rotateRight(out, pos)
			case "based":
				out = rotateByPositionOfLetter(out, words[6])
			}
		case "reverse":
			pos1, _ := strconv.Atoi(words[2])
			pos2, _ := strconv.Atoi(words[4])
			out = reverse(out, pos1, pos2)
		case "move":
			pos1, _ := strconv.Atoi(words[2])
			pos2, _ := strconv.Atoi(words[5])
			out = move(out, pos1, pos2)
		default:
			panic(words[0])
		}
	}
	return
}

func (scr scrambler) unscramble(in string) (out string) {
	out = in
	for i := len(scr.instructions) - 1; i >= 0; i-- {
		words := strings.Split(scr.instructions[i], " ")
		switch words[0] {
		case "swap":
			switch words[1] {
			case "position":
				pos1, _ := strconv.Atoi(words[2])
				pos2, _ := strconv.Atoi(words[5])
				out = swapPositions(out, pos1, pos2)
			case "letter":
				out = swapLetters(out, words[2], words[5])
			}
		case "rotate":
			switch words[1] {
			case "left":
				pos, _ := strconv.Atoi(words[2])
				out = rotateRight(out, pos)
			case "right":
				pos, _ := strconv.Atoi(words[2])
				out = rotateLeft(out, pos)
			case "based":
				out = unRotateByPositionOfLetter(out, words[6])
			}
		case "reverse":
			pos1, _ := strconv.Atoi(words[2])
			pos2, _ := strconv.Atoi(words[4])
			out = reverse(out, pos1, pos2)
		case "move":
			pos1, _ := strconv.Atoi(words[2])
			pos2, _ := strconv.Atoi(words[5])
			out = move(out, pos2, pos1)
		default:
			panic(words[0])
		}
	}
	return
}

func swapPositions(in string, pos1, pos2 int) (out string) {
	tmp := []rune(in)
	tmp[pos1], tmp[pos2] = tmp[pos2], tmp[pos1]
	out = string(tmp)
	return
}

func swapLetters(in, ltr1, ltr2 string) (out string) {
	out = swapPositions(in, strings.Index(in, ltr1), strings.Index(in, ltr2))
	return
}

func rotateLeft(in string, steps int) (out string) {
	out = in
	for i := 0; i < steps; i++ {
		out = out[1:] + string(out[0])
	}
	return
}

func rotateRight(in string, steps int) (out string) {
	out = in
	for i := 0; i < steps; i++ {
		out = string(out[len(out)-1]) + out[:len(out)-1]
	}
	return
}

func rotateByPositionOfLetter(in, ltr string) (out string) {
	pos := strings.Index(in, ltr)
	steps := 1 + pos
	if pos >= 4 {
		steps++
	}
	out = rotateRight(in, steps)
	return
}

func unRotateByPositionOfLetter(in, ltr string) (out string) {
	if len(in) != 8 {
		panic("unRotating only works for passwords with a length of 8!")
	}
	pos := strings.Index(in, ltr)
	steps := pos/2 + 1
	if pos%2 == 0 && pos != 0 {
		steps += 4
	}
	out = rotateLeft(in, steps)
	return
}

func reverse(in string, pos1, pos2 int) (out string) {
	out = in
	for pos1 < pos2 {
		out = swapPositions(out, pos1, pos2)
		pos1++
		pos2--
	}
	return
}

func move(in string, pos1, pos2 int) (out string) {
	ltr := string(in[pos1])
	out = in[:pos1] + in[pos1+1:] + ltr
	for strings.Index(out, ltr) != pos2 {
		out = swapPositions(out, strings.Index(out, ltr), strings.Index(out, ltr)-1)
	}
	return
}

func getInput() (scr scrambler) {
	in, _ := ioutil.ReadFile("input.txt")
	scr.instructions = []string{}
	for _, v := range strings.Split(string(in), "\n") {
		if len(v) > 0 {
			scr.instructions = append(scr.instructions, v)
		}
	}
	return
}

func main() {
	scrambler := getInput()
	fmt.Printf("abcdefgh is scrambled to %v.\n", scrambler.scramble("abcdefgh"))
	fmt.Printf("fbgdceah is unscrambled to %v.\n", scrambler.unscramble("fbgdceah"))
}
