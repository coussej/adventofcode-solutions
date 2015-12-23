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
	secret := "bgvyzdsv"
	var i int
	for {
		hash := getHashHex(secret + strconv.Itoa(i))
		if strings.Index(hash, "00000") == 0 {
			fmt.Println("The smallest positive number that generates a hash with 5 leading zeros is", i)
			break
		}
		i++
	}
	// no need to reset counter. we re-evalute the same i to see if it's also valid for 6.
	for {
		hash := getHashHex(secret + strconv.Itoa(i))
		if strings.Index(hash, "000000") == 0 {
			fmt.Println("The smallest positive number that generates a hash with 6 leading zeros is", i)
			break
		}
		i++
	}
}
