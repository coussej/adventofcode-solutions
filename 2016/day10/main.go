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

func (b *bot) isReadyToGive() bool {
	return b.high > 0
}

type botlist map[int]bot

func (bl botlist) assign(botID, val int) {
	b := bl[botID]
	b.assign(val)
	bl[botID] = b
}

type valueInstruction struct {
	botID int
	value int
}

type targetInstruction struct {
	botID         int
	valueType     string
	recipientID   int
	recipientType string
}

func getInstructions() (vInstr []valueInstruction, tInstr []targetInstruction) {
	in, _ := ioutil.ReadFile("input.txt")
	in2 := strings.TrimSpace(regexp.MustCompile("\n *").ReplaceAllString(string(in), ";"))
	for _, instr := range strings.Split(in2, ";") {
		words := strings.Split(instr, " ")
		switch words[0] {
		case "value":
			val, _ := strconv.Atoi(words[1])
			id, _ := strconv.Atoi(words[5])
			vInstr = append(vInstr, valueInstruction{id, val})
		case "bot":
			id, _ := strconv.Atoi(words[1])
			lowid, _ := strconv.Atoi(words[6])
			highid, _ := strconv.Atoi(words[11])
			tInstr = append(tInstr, targetInstruction{id, "low", lowid, words[5]})
			tInstr = append(tInstr, targetInstruction{id, "high", highid, words[10]})
		}
	}
	return
}

func main() {
	bots := botlist{}
	outputs := map[int]int{}

	vInstr, tInstr := getInstructions()

	// assign values from input bins
	for _, v := range vInstr {
		bots.assign(v.botID, v.value)
	}

	// execute target instructions until every instruction succeeds
	allInstructionsOK := false
	for !allInstructionsOK {
		allInstructionsOK = true
		for _, t := range tInstr {
			bot := bots[t.botID]
			if !bot.isReadyToGive() {
				allInstructionsOK = false
				continue
			}
			var val = 0
			if t.valueType == "low" {
				val = bot.low
			} else {
				val = bot.high
			}
			if t.recipientType == "bot" {
				bots.assign(t.recipientID, val)
			} else {
				outputs[t.recipientID] = val
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
