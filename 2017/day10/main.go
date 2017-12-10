package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type list []int

func main() {
	// part 1
	lengths, list := getInput(), newList()
	list.process(lengths, 0, 0)

	fmt.Printf("The result of multiplying the first two numbers in the list is %v.\n", list[0]*list[1])

	// part 2
	lengths2, list2 := getInputASCII(), newList()
	lengths2 = append(lengths2, 17, 31, 73, 47, 23)
	curPos, skipSize := 0, 0
	for i := 0; i < 64; i++ {
		curPos, skipSize = list2.process(lengths2, curPos, skipSize)
	}
	hash := list2.toDenseHash().toHexString()

	fmt.Printf("The knot hash of the input is %v.\n", hash)
}

func newList() (l list) {
	for i := 0; i < 256; i++ {
		l = append(l, i)
	}
	return
}

func (l list) process(lengths []int, iCurPos, iSkipSize int) (curPos, skipSize int) {
	curPos, skipSize = iCurPos, iSkipSize
	for _, length := range lengths {
		temp := list{}
		temp = append(append(temp, l...), l...)
		sub := temp[curPos : curPos+length]
		sub.reverse()
		replacePos := curPos
		for _, v := range sub {
			l[replacePos] = v
			replacePos = nextPos(replacePos, 1, len(l))
		}
		curPos = nextPos(curPos, length+skipSize, len(l))
		skipSize++
	}
	return
}

func (l *list) reverse() {
	for i, j := 0, len(*l)-1; i < j; i, j = i+1, j-1 {
		(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
	}
}

func (l list) toDenseHash() (hash list) {
	for i := 0; i < len(l); i = i + 16 {
		h := l[i]
		for j := 1; j < 16; j++ {
			h = h ^ l[i+j]
		}
		hash = append(hash, h)
	}
	return
}

func (l list) toHexString() (hash string) {
	for _, v := range l {
		str := fmt.Sprintf("%x", v)
		if len(str) == 1 {
			str = "0" + str
		}
		hash = hash + str
	}
	return
}

func nextPos(cur, jump, max int) int {
	return (cur + jump) % max
}

func getInput() (lengths []int) {
	in, _ := ioutil.ReadFile("input.txt")
	for _, num := range strings.Split(strings.Replace(string(in), "\n", "", -1), ",") {
		l, _ := strconv.Atoi(num)
		lengths = append(lengths, l)
	}
	return
}

func getInputASCII() (lengths []int) {
	in, _ := ioutil.ReadFile("input.txt")
	for _, r := range strings.Replace(string(in), "\n", "", -1) {
		lengths = append(lengths, int(r))
	}
	return
}
