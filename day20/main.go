package main

import (
	"fmt"
	"math"
)

func main() {
	input := 29000000
	for i := 1; ; i++ {
		if findNumberOfPresents(i, 10, -1) >= 29000000 {
			fmt.Println("The first house to get at least", input, "presents is", i)
			break
		}
	}
	for i := 1; ; i++ {
		if findNumberOfPresents(i, 11, 50) >= 29000000 {
			fmt.Println("The first house to get at least", input, "presents with lazy elves is", i)
			break
		}
	}
}

func findDivisors(num int) (div []int) {
	div = []int{}
	for i := 1; i <= int(math.Ceil(math.Sqrt(float64(num)))); i++ {
		if num%i == 0 {
			switch i == num/i {
			case true:
				div = append(div, i)
			case false:
				div = append(div, []int{i, num / i}...)
			}
		}
	}
	return
}

func findNumberOfPresents(house, elfPresentFactor, elfHouseLimit int) (presents int) {
	for _, v := range findDivisors(house) {
		if elfHouseLimit < 0 || v*elfHouseLimit >= house {
			presents += elfPresentFactor * v
		}
	}
	return
}
