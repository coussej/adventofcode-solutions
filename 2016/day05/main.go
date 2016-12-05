package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func getHashHex(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

func main() {
	doorID := "ffykfhsq"
	pass1 := [8]string{"_", "_", "_", "_", "_", "_", "_", "_"}
	pass2 := [8]string{"_", "_", "_", "_", "_", "_", "_", "_"}
	pass1count := 0
	pass2count := 0
	pass2pos := map[int]bool{}
	var i int
	for {
		hash := getHashHex(doorID + strconv.Itoa(i))
		if strings.Index(hash, "00000") == 0 {
			// password first door
			if pass1count < 8 {
				pass1[pass1count] = hash[5:6]
				pass1count++
			}

			// password second door
			pos, err := strconv.Atoi(hash[5:6])
			if err == nil && pos < 8 && !pass2pos[pos] {
				pass2[pos] = hash[6:7]
				pass2pos[pos] = true
				pass2count++
			}

			// animation !
			fmt.Printf("\rDOOR1 > %v <  |  DOOR2  > %v <",
				strings.Join(pass1[:], ""),
				strings.Join(pass2[:], ""))

			// break if password of second door is complete
			if pass2count == 8 {
				fmt.Printf("\n")
				break
			}
		}
		i++
	}
}
