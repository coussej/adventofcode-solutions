package main

import "fmt"

func main() {
	fmt.Println("The code required to power the weather machine is:", getCodeAtLocation(3010, 3019))
}

func getCodeAtLocation(row, col int) int {
	code := 20151125
	for i := 2; true; i++ {
		r := i
		c := 1
		for r > 0 {
			code = (code * 252533) % 33554393
			if r == row && c == col {
				return code
			}
			r--
			c++
		}
	}
	return 0
}
