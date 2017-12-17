package main

import "fmt"

type circularBuffer struct {
	buf []int
	len int
	pos int
}

func main() {
	input := 367

	// part 1

	buf := []int{0}
	pos := 0
	for i := 1; i <= 2017; i++ {
		pos = (pos + input) % i
		buf = append(buf, 0)
		copy(buf[pos+2:], buf[pos+1:])
		buf[pos+1] = i
		pos = (pos + 1) % i
	}

	pos = (pos + 1) % len(buf)
	fmt.Printf("The value after 2017 in the buffer is %v.\n", buf[pos])

	// part 2: don't keep buffer. Only insertions after pos O are relevant.

	pos, val := 0, 0
	for i := 1; i <= 50000000; i++ {
		pos = (pos + input) % i
		if pos == 0 {
			val = i
		}
		pos++
	}
	fmt.Printf("The value after 0 at the moment 50000000 is inserted is %v.\n", val)
}

func next(cur, len, steps int) int {
	return (cur + steps) % len
}
