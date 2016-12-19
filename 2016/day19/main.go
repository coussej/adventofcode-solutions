package main

import "fmt"

func main() {
	in := 3001330

	// part 1: result in function of the nearest lower 2^f
	f2 := 1
	for f2*2 <= in {
		f2 *= 2
	}
	winningElf := in/f2 + 2*(in%f2)
	fmt.Printf("Elf %v gets all the presents when stealing from the left.\n", winningElf)

	// part2: result in function of the nearest lower 3^f
	f3 := 1
	for f3*3 < in {
		f3 *= 3
	}
	winningElf = (in/(f3*3)+(in/f3)-1)*f3 + (in/f3)*(in%f3)
	fmt.Printf("Elf %v gets all the presents when stealing from across.\n", winningElf)
}
