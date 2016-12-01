package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Ingredient struct {
	Name       string
	Capacity   int
	Durability int
	Flavor     int
	Texture    int
	Calories   int
}

func main() {
	ing := getIngredients()

	var maxScore int
	var maxScoreInclCal int
	for _, c := range getPossibleCombos(len(ing), 100) {
		s := getCookieScore(ing, c, 0)
		if s > maxScore {
			maxScore = s
		}
		s = getCookieScore(ing, c, 500)
		if s > maxScoreInclCal {
			maxScoreInclCal = s
		}
	}

	fmt.Println("The cookie with the highest score has", maxScore, "points.")
	fmt.Println("The best scoring cookie with 500 cal has", maxScoreInclCal, "points.")
}

func getIngredients() []Ingredient {
	in, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	var ing []Ingredient
	for _, line := range strings.Split(string(in), "\n") {
		words := strings.Split(line, " ")
		c, _ := strconv.Atoi(words[2][:len(words[2])-1])
		d, _ := strconv.Atoi(words[4][:len(words[4])-1])
		f, _ := strconv.Atoi(words[6][:len(words[6])-1])
		t, _ := strconv.Atoi(words[8][:len(words[8])-1])
		cal, _ := strconv.Atoi(words[10])
		ing = append(ing,
			Ingredient{
				Name:       words[0],
				Capacity:   c,
				Durability: d,
				Flavor:     f,
				Texture:    t,
				Calories:   cal,
			})
	}
	return ing
}

func getPossibleCombos(lenght, total int) [][]int {
	var c [][]int
	if lenght > 1 {
		for i := total; i >= 0; i-- {
			n := []int{i}
			for _, v := range getPossibleCombos(lenght-1, total-i) {
				c = append(c, append(n, v...))
			}
		}
	} else {
		c = append(c, []int{total})
	}
	return c
}

func getCookieScore(ings []Ingredient, balance []int, calories int) int {
	var c, d, f, t, cal int
	for i, ing := range ings {
		c += ing.Capacity * balance[i]
		d += ing.Durability * balance[i]
		f += ing.Flavor * balance[i]
		t += ing.Texture * balance[i]
		cal += ing.Calories * balance[i]
	}
	if c < 0 {
		c = 0
	}
	if d < 0 {
		d = 0
	}
	if f < 0 {
		f = 0
	}
	if t < 0 {
		t = 0
	}
	if calories > 0 && cal != calories {
		return 0
	} else {
		return c * d * f * t
	}
}
