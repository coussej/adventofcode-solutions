package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type bot struct {
	low  int
	high int
}

func (b *bot) assign(val int) {
	if b.low == 0 {
		b.low = val
	} else if b.low < val {
		b.high = val
	} else if b.low != val {
		b.low, b.high = val, b.low
	}
}

type botlist map[int]bot

func (bl botlist) assign(botID, val int) {
	b := bl[botID]
	b.assign(val)
	bl[botID] = b
}

type valueInstruction struct {
	botID, value int
}

type giveInstruction struct {
	botID         int
	valueType     string
	recipientID   int
	recipientType string
}

func getInstructions() []string {
	in, _ := ioutil.ReadFile("input.txt")
	in2 := strings.TrimSpace(regexp.MustCompile("\n *").ReplaceAllString(string(in), ";"))
	return strings.Split(in2, ";")
}

func main() {
	bots := botlist{}
	outputs := map[int]int{}

	instructionFailed := true

	for instructionFailed {
		instructionFailed = false
		for _, instr := range getInstructions() {
			words := strings.Split(instr, " ")
			switch words[0] {
			case "value":
				val, _ := strconv.Atoi(words[1])
				id, _ := strconv.Atoi(words[5])
				bots.assign(id, val)
			case "bot":
				id, _ := strconv.Atoi(words[1])
				bot := bots[id]
				if bot.high > 0 {
					//assign low
					lowid, _ := strconv.Atoi(words[6])
					switch words[5] {
					case "output":
						outputs[lowid] = bot.low
					default:
						bots.assign(lowid, bot.low)
					}
					highid, _ := strconv.Atoi(words[11])
					switch words[10] {
					case "output":
						outputs[highid] = bot.high
					default:
						bots.assign(highid, bot.high)
					}
				} else {
					instructionFailed = true
				}
			}
		}
	}

	for id, b := range bots {
		if b.low == 17 && b.high == 61 {
			fmt.Printf("Bot %v compares value 17 and value 61 microchips.\n", id)
		}
	}
	fmt.Printf("The product of the values of the microchips in bins 0, 1 and 2 is %v.\n",
		outputs[0]*outputs[1]*outputs[2])
}
