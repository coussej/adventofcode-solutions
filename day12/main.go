package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	in, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("The sum of all numbers in the document is", totalAllIntsInString(string(in)))
	fmt.Println("After removing any object with a red property, the sum is", totalAllIntsInString(removeAllEntitiesContaining(string(in), "red")))
}

func totalAllIntsInString(s string) int {
	reg := regexp.MustCompile(`(-?\d+)`)
	var total int
	for _, v := range reg.FindAllString(string(s), -1) {
		i, _ := strconv.Atoi(v)
		total += i
	}
	return total
}

func removeAllEntitiesContaining(s, v string) string {
	res := s
	for i := 0; i < len(res); i++ {
		if string(res[i]) == "}" {
			// find opening bracket.
			balance := 1
			for j := i - 1; j >= 0; j-- {
				if string(res[j]) == "}" {
					balance++
				}
				if string(res[j]) == "{" {
					balance--
				}
				if balance == 0 {
					if strings.Contains(res[j:i+1], ":\""+v+"\"") {
						res = res[:j] + res[i+1:]
						i = j - 1
					}
					break
				}
			}
		}
	}

	return res
}
