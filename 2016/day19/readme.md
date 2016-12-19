package main

type elf struct {
  number, presents int
}

func main() {
  input := 5
  elves := []int{}
  for i:= 0, i<input; i++ {
    elves = append(elves, elf{i+1, 1})
  }
  curr := 0
  for len(elves) > 1 {
    // find elf to take presents from
    if (curr == len(elves) - 1) {
      elves[curr].presents += elves[0].presents
      elves = elves[1:]
      curr = 0
    } else {
      elves[curr].presents += elves[curr+1].presents
      elves = append(elves[:curr+1], elves[curr+2:]...)
      curr++
    }
  }
  fmt.Println(elves)
}
