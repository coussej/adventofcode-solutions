package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type room struct {
	name     string
	sectorID int
	checksum string
}

func (r *room) isReal() bool {
	frequencies := map[string]int{}
	mostfrequent := []string{}
	maxoccur := 0

	for _, l := range strings.Replace(r.name, "-", "", -1) {
		frequencies[string(l)]++
		if o := frequencies[string(l)]; o > maxoccur {
			maxoccur = o
		}
	}

	for i := maxoccur; i > 0; i-- {
		letters := []string{}
		for l, f := range frequencies {
			if f == i {
				letters = append(letters, l)
			}
		}
		sort.Strings(letters)
		mostfrequent = append(mostfrequent, letters...)
	}
	return r.checksum == strings.Join(mostfrequent[0:5], "")
}

func (r *room) decryptName() string {
	var result bytes.Buffer
	for _, ltr := range r.name {
		if ltr == '-' {
			result.WriteRune(' ')
		} else {
			newLtr := ltr
			for i := 0; i < r.sectorID; i++ {
				newLtr = newLtr + 1
				if newLtr == '{' {
					newLtr = 'a'
				}
			}
			result.WriteRune(newLtr)
		}
	}
	return result.String()
}

func getRooms() (rooms []room) {
	in, _ := ioutil.ReadFile("input.txt")
	in2 := strings.TrimSpace(regexp.MustCompile("\n *").ReplaceAllString(string(in), " "))

	for _, v := range strings.Split(in2, " ") {
		ld := strings.LastIndex(v, "-")
		fb := strings.Index(v, "[")
		sectorID, _ := strconv.Atoi(v[ld+1 : fb])
		r := room{
			v[:ld],
			sectorID,
			v[fb+1 : len(v)-1],
		}
		rooms = append(rooms, r)
	}
	return
}

func main() {
	sum := 0
	nosID := 0
	for _, r := range getRooms() {
		if r.isReal() {
			sum += r.sectorID
		}
		if r.decryptName() == "northpole object storage" {
			nosID = r.sectorID
		}
	}
	fmt.Println("The sum of the sector ID's from the real rooms is", sum, ".")
	fmt.Println("The northpole object storage has sector ID", nosID, ".")
}
