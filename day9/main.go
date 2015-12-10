package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type routemap map[string]map[string]int

func main() {
	r := getRoutemapFromInput("input.txt")

	// for all startpoints
	var min, max int
	var minp, maxp []string
	for place, _ := range r {
		d, p := r.findShortestRemainingPath([]string{place})
		if d < min || min == 0 {
			min = d
			minp = p
		}
		d, p = r.findLongestRemainingPath([]string{place})
		if d > max || min == 0 {
			max = d
			maxp = p
		}
	}
	fmt.Println("Shortest Path:", minp, "| Distance:", min)
	fmt.Println("Longest Path: ", maxp, "| Distance:", max)
}

func (r routemap) addRoute(from, to string, distance int) {
	m, ok := r[from]
	if !ok {
		m = make(map[string]int)
		r[from] = m
	}
	r[from][to] = distance
}

func (r routemap) findShortestRemainingPath(currentPath []string) (int, []string) {
	var shortestPath []string
	var distance int
	for place, dist := range r[currentPath[len(currentPath)-1]] {
		if !sliceContains(currentPath, place) {
			//fmt.Println("Visiting", append(currentPath, place))
			d, path := r.findShortestRemainingPath(append(currentPath, place))
			if distance == 0 || d+dist < distance {
				distance = d + dist
				shortestPath = make([]string, len(path))
				copy(shortestPath, path)
			}
		}
	}
	if distance == 0 {
		shortestPath = make([]string, len(currentPath))
		copy(shortestPath, currentPath)
	}
	return distance, shortestPath
}

func (r routemap) findLongestRemainingPath(currentPath []string) (int, []string) {
	var longestPath []string
	var distance int
	for place, dist := range r[currentPath[len(currentPath)-1]] {
		if !sliceContains(currentPath, place) {
			//fmt.Println("Visiting", append(currentPath, place))
			d, path := r.findLongestRemainingPath(append(currentPath, place))
			if distance == 0 || d+dist > distance {
				distance = d + dist
				longestPath = make([]string, len(path))
				copy(longestPath, path)
			}
		}
	}
	if distance == 0 {
		longestPath = make([]string, len(currentPath))
		copy(longestPath, currentPath)
	}
	return distance, longestPath
}

func getRoutemapFromInput(file string) routemap {
	in, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	r := routemap{}
	for _, line := range strings.Split(string(in), "\n") {
		words := strings.Split(line, " ")
		// Add route and reverse route to routes
		dist, _ := strconv.Atoi(words[4])
		r.addRoute(words[0], words[2], dist)
		r.addRoute(words[2], words[0], dist)
	}
	return r
}

func sliceContains(s []string, v string) bool {
	for _, val := range s {
		if val == v {
			return true
		}
	}
	return false
}
