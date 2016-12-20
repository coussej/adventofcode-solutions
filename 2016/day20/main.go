package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type ipRange struct {
	start, end int
}

type blacklist []ipRange

func (b blacklist) Len() int      { return len(b) }
func (b blacklist) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b blacklist) Less(i, j int) bool {
	if b[i].start < b[j].start {
		return true
	}
	if b[i].start > b[j].start {
		return false
	}
	return b[i].end < b[j].end
}

func (b *blacklist) clean() {
	bl := *b
	sort.Sort(bl)
	i := 1
	for i < len(bl) {
		if bl[i].start <= bl[i-1].end+1 {
			if bl[i].end >= bl[i-1].end {
				bl[i-1].end = bl[i].end
			}
			bl = append(bl[:i], bl[i+1:]...)
		} else {
			i++
		}
	}
	*b = bl
}

func getInput() (b blacklist) {
	in, _ := ioutil.ReadFile("input.txt")
	in2 := strings.TrimSpace(regexp.MustCompile("\n *").ReplaceAllString(string(in), " "))
	for _, v := range strings.Split(in2, " ") {
		nums := strings.Split(string(v), "-")
		start, _ := strconv.Atoi(nums[0])
		end, _ := strconv.Atoi(nums[1])
		b = append(b, ipRange{start, end})
	}
	return
}

func main() {
	bl := getInput()
	bl.clean()
	fmt.Printf("The lowest-valued ip that is not blocked is %v.\n", bl[0].end+1)
	count := bl[0].start
	for i := 1; i < len(bl); i++ {
		count += bl[i].start - bl[i-1].end - 1
		if i == len(bl)-1 {
			count += 4294967295 - bl[i].end
		}
	}
	fmt.Printf("In total, %v addresses are allowed.\n", count)
}
