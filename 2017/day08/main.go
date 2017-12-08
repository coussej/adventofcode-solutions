package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type instruction struct {
	target          string
	action          string
	amount          int
	conditionTarget string
	condition       string
	conditionAmount int
}

type register map[string]int

func main() {
	reg, instr := getInput()

	overallMax := math.MinInt64
	for _, i := range instr {
		reg.execInstruction(i)
		localMax := reg.getHighestValue()
		if localMax > overallMax {
			overallMax = localMax
		}
	}

	fmt.Printf("The highest final value in the register is %v.\n", reg.getHighestValue())
	fmt.Printf("The highest value that was encountered during execution was %v.\n", overallMax)
}

func (r register) execInstruction(i instruction) {
	// check condition
	if !evaluateCondition(r[i.conditionTarget], i.conditionAmount, i.condition) {
		return
	}

	// execute action
	switch i.action {
	case "inc":
		r[i.target] += i.amount
	case "dec":
		r[i.target] -= i.amount
	}
	return
}

func (r register) getHighestValue() int {
	max := math.MinInt64
	for _, v := range r {
		if v > max {
			max = v
		}
	}
	return max
}

func evaluateCondition(x1, x2 int, cond string) bool {
	switch cond {
	case "==":
		return x1 == x2
	case ">":
		return x1 > x2
	case "<":
		return x1 < x2
	case ">=":
		return x1 >= x2
	case "<=":
		return x1 <= x2
	case "!=":
		return x1 != x2
	default:
		panic("invalid condition: " + cond)
	}
}

func getInput() (reg register, instr []instruction) {
	reg = register{}
	in, _ := ioutil.ReadFile("input.txt")
	for _, line := range strings.Split(string(in), "\n") {
		if len(line) > 0 {
			words := strings.Split(line, " ")
			reg[words[0]] = 0
			amount, _ := strconv.Atoi(words[2])
			conditionAmount, _ := strconv.Atoi(words[6])
			instr = append(instr, instruction{
				target:          words[0],
				action:          words[1],
				amount:          amount,
				conditionTarget: words[4],
				condition:       words[5],
				conditionAmount: conditionAmount,
			})
		}
	}
	return
}
