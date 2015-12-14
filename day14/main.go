package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Herd []Reindeer

type Reindeer struct {
	Name             string
	Speed            int
	ActivityDuration int
	PauseDuration    int
	DistanceTraveled int
	TimesInLead      int
	Cycletime        int // negative is pause, positive is activity
}

func main() {
	herd := getHerd()
	herd.Race(2503)

	fmt.Println("The largest distance travelled was", herd.getMaxDistance(), "km.")
	fmt.Println("The the most points gathered was", herd.getMaxTimesInLead(), "points.")
}

func getHerd() Herd {
	in, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	var herd []Reindeer
	for _, line := range strings.Split(string(in), "\n") {
		words := strings.Split(line, " ")
		s, _ := strconv.Atoi(words[3])
		a, _ := strconv.Atoi(words[6])
		p, _ := strconv.Atoi(words[13])
		herd = append(herd,
			Reindeer{
				Name:             words[0],
				Speed:            s,
				ActivityDuration: a,
				PauseDuration:    p,
			})
	}
	return Herd(herd)
}

func (h *Herd) Race(seconds int) {
	for i := 0; i < seconds; i++ {
		for i, _ := range *h {
			if (*h)[i].Cycletime >= 0 {
				(*h)[i].DistanceTraveled += (*h)[i].Speed
			}
			(*h)[i].Cycletime++
			if (*h)[i].Cycletime == (*h)[i].ActivityDuration {
				(*h)[i].Cycletime = -(*h)[i].PauseDuration
			}
		}
		// part 2: give points when in lead.
		max := h.getMaxDistance()
		for i, _ := range *h {
			if (*h)[i].DistanceTraveled == max {
				(*h)[i].TimesInLead++
			}
		}
	}
}

func (h *Herd) getMaxDistance() int {
	var max int
	for _, r := range *h {
		if r.DistanceTraveled > max {
			max = r.DistanceTraveled
		}
	}
	return max
}

func (h *Herd) getMaxTimesInLead() int {
	var max int
	for _, r := range *h {
		if r.TimesInLead > max {
			max = r.TimesInLead
		}
	}
	return max
}
