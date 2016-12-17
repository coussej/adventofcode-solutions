package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type position struct {
	x, y int
	path string
}

func (p position) getPossibleMoves(passcode string) (moves []string) {
	hash := getHashHex(passcode + p.path)
	for i, dir := range []string{"U", "D", "L", "R"} {
		if hash[i] > 'a' {
			moves = append(moves, dir)
		}
	}
	return
}

func (p position) move(dir string) (newPos position, success bool) {
	switch {
	case dir == "U" && p.y > 0:
		p.y, p.path, success = p.y-1, p.path+"U", true
	case dir == "D" && p.y < 3:
		p.y, p.path, success = p.y+1, p.path+"D", true
	case dir == "L" && p.x > 0:
		p.x, p.path, success = p.x-1, p.path+"L", true
	case dir == "R" && p.x < 3:
		p.x, p.path, success = p.x+1, p.path+"R", true
	}
	newPos = p
	return
}

func (p position) isFinal() bool {
	return p.x == 3 && p.y == 3
}

func getHashHex(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

func getShortestAndLongestPath(passcode string) (shortestPath, longestPath string) {
	steps, pos := 0, []position{position{0, 0, ""}}
	for len(pos) > 0 {
		steps++
		nPos := []position{}
		for _, p := range pos {
			moves := p.getPossibleMoves(passcode)
			for _, m := range moves {
				if np, success := p.move(m); success {
					if np.isFinal() {
						if shortestPath == "" {
							shortestPath = np.path
						}
						longestPath = np.path
					} else {
						nPos = append(nPos, np)
					}
				}
			}
		}
		pos = nPos
	}
	return
}

func main() {
	passcode := "awrkjxxr"

	s, l := getShortestAndLongestPath(passcode)

	fmt.Println("The shortest path is", s)
	fmt.Println("The length of the longest path is", len(l))
}
