package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	packs := getPackageWeights()
	fmt.Println("When divided in 3 groups, the smallest group has a QE of", getQEofBestConfig(packs, 3))
	fmt.Println("When divided in 4 groups, the smallest group has a QE of", getQEofBestConfig(packs, 4))
}

func getQEofBestConfig(packs []int, numberOfGroups int) int {
	var n_min, qe_min int
	for _, group := range getPossibleCombinationsForTotal(packs, sum(packs)/numberOfGroups) {
		if len(group) < n_min || n_min == 0 || (len(group) == n_min && qe_min > product(group)) {
			n_min = len(group)
			qe_min = product(group)
		}
	}
	return qe_min
}

func getPossibleCombinationsForTotal(numbers []int, total int) (result [][]int) {
	for i, n := range numbers {
		switch {
		case total-n == 0:
			// solution found. Add to result.
			result = append(result, []int{n})
		case total-n > 0:
			// recurse.
			remNumbers := []int{}
			copy(remNumbers, numbers[:i])
			remNumbers = append(remNumbers, numbers[i+1:]...)
			for _, r := range getPossibleCombinationsForTotal(remNumbers, total-n) {
				res := []int{n}
				res = append(res, r...)
				result = append(result, res)
			}
		}
	}
	return
}

func getPackageWeights() (weights []int) {
	in, _ := ioutil.ReadFile("input.txt")
	for _, line := range strings.Split(string(in), "\n") {
		w, _ := strconv.Atoi(line)
		weights = append(weights, w)
	}
	return
}

func sum(s []int) (sum int) {
	for _, v := range s {
		sum += v
	}
	return
}

func product(s []int) (product int) {
	product = 1
	for _, v := range s {
		product *= v
	}
	return
}
