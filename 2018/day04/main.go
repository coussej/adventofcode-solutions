package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"time"
)

type logEntry struct {
	ts      time.Time
	raw     string
	guardID int
}

type sleepActivity struct {
	ts      time.Time
	guardID int
	minutes int
}

type summary struct {
	guardID            int
	sleepStats         []int
	totalMinutesAsleep int
	mostSleptMinute    int
	mostSleptNumber    int
}

func main() {
	activities := getActivities()

	guardSummary := map[int]*summary{}
	for _, a := range activities {
		if guardSummary[a.guardID] == nil {
			guardSummary[a.guardID] = &summary{sleepStats: make([]int, 60, 60), guardID: a.guardID}
		}
		for m := a.ts.Minute(); m < a.ts.Minute()+a.minutes; m++ {
			guardSummary[a.guardID].sleepStats[m]++
			guardSummary[a.guardID].totalMinutesAsleep++
			if guardSummary[a.guardID].sleepStats[m] > guardSummary[a.guardID].mostSleptNumber {
				guardSummary[a.guardID].mostSleptMinute = m
				guardSummary[a.guardID].mostSleptNumber = guardSummary[a.guardID].sleepStats[m]
			}
		}
	}

	strategyOneBest := &summary{}
	for _, activity := range guardSummary {
		if activity.totalMinutesAsleep > strategyOneBest.totalMinutesAsleep {
			strategyOneBest = activity
		}
	}
	fmt.Printf("For strategy 1, the ID of the guard multiplied by the minute is %v.\n", strategyOneBest.guardID*strategyOneBest.mostSleptMinute)

	strategyTwoBest := &summary{}
	for _, activity := range guardSummary {
		if activity.mostSleptNumber > strategyTwoBest.mostSleptNumber {
			strategyTwoBest = activity
		}
	}
	fmt.Printf("For strategy 2, the ID of the guard multiplied by the minute is %v.\n", strategyTwoBest.guardID*strategyTwoBest.mostSleptMinute)
}

func getActivities() (activities []sleepActivity) {
	in, _ := ioutil.ReadFile("input.txt")
	pat := regexp.MustCompile(`\[([-: 0-9]+)\] ([ #a-zA-Z]+(\d*)[ #a-zA-Z]+)`)
	matches := pat.FindAllStringSubmatch(string(in), -1)
	raw := []logEntry{}
	for _, m := range matches {
		layout := "2006-01-02 15:04"
		ts, _ := time.Parse(layout, m[1])
		id, _ := strconv.Atoi(m[3])
		raw = append(raw, logEntry{ts, m[2], id})
	}

	sort.Slice(raw, func(i, j int) bool { return raw[i].ts.Before(raw[j].ts) })

	currentGuard := 0
	for i, r := range raw {
		if r.guardID > 0 {
			currentGuard = r.guardID
			continue
		}
		if r.raw == "falls asleep" {
			activities = append(activities, sleepActivity{r.ts, currentGuard, raw[i+1].ts.Minute() - r.ts.Minute()})
		}
	}
	return
}
