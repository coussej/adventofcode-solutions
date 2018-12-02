package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	ids := getIds()

	cs2, cs3 := 0, 0

	for _, id := range ids {
		has2, has3 := false, false
		for _, c := range id {
			if strings.Count(id, string(c)) == 2 {
				has2 = true
			}
			if strings.Count(id, string(c)) == 3 {
				has3 = true
			}
		}
		if has2 {
			cs2++
		}
		if has3 {
			cs3++
		}
	}

	fmt.Printf("The resulting checksum is %v\n", cs2*cs3)

	// Part two

	var result string

search:
	for _, id := range ids {
		for _, id2 := range ids {
			result = ""
			for i := 0; i < len(id); i++ {
				if id[i] == id2[i] {
					result += string(id[i])
				}
			}
			if len(id)-len(result) == 1 {
				break search
			}
		}
	}

	fmt.Printf("The letters the 2 IDs have in common are %v\n", result)

}

func getIds() (ids []string) {
	in, _ := ioutil.ReadFile("input.txt")
	for _, id := range strings.Split(string(in), "\n") {
		if len(id) > 0 {
			ids = append(ids, id)
		}
	}
	return
}
