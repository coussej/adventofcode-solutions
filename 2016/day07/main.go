package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type IPv7 struct {
	address     string
	hypernetseq []string
	supernetseq []string
}

func getIPv7Adresses() (ips []IPv7) {
	in, _ := ioutil.ReadFile("input.txt")
	in2 := strings.TrimSpace(regexp.MustCompile("\n *").ReplaceAllString(string(in), " "))

	hypernetregex := regexp.MustCompile("\\[.*?\\]")
	for _, ip := range strings.Split(in2, " ") {
		hyper := hypernetregex.FindAllString(ip, -1)
		for i := range hyper {
			hyper[i] = hyper[i][1 : len(hyper[i])-1]
		}
		super := hypernetregex.Split(ip, -1)

		ips = append(ips, IPv7{ip, hyper, super})
	}
	return
}

func (ip *IPv7) supportsTLS() bool {
	for _, h := range ip.hypernetseq {
		if containsABBA(h) {
			return false
		}
	}
	for _, s := range ip.supernetseq {
		if containsABBA(s) {
			return true
		}
	}
	return false
}

func (ip *IPv7) supportsSSL() bool {
	abas := []string{}
	for _, h := range ip.hypernetseq {
		abas = append(abas, getABAs(h)...)
	}
	for _, s := range ip.supernetseq {
		for _, aba := range abas {
			if containsBABsforABA(s, aba) {
				return true
			}
		}
	}
	return false
}

func containsABBA(s string) bool {
	if len(s) < 4 {
		return false
	}
	for i := 0; i < len(s)-3; i++ {
		if s[i] != s[i+1] &&
			s[i+1] == s[i+2] &&
			s[i] == s[i+3] {
			return true
		}
	}
	return false
}

func getABAs(s string) (abas []string) {
	if len(s) < 3 {
		return
	}
	for i := 0; i < len(s)-2; i++ {
		if s[i] != s[i+1] &&
			s[i] == s[i+2] {
			abas = append(abas, s[i:i+3])
		}
	}
	return
}

func containsBABsforABA(s string, aba string) bool {
	if len(s) < 3 || len(aba) != 3 {
		return false
	}
	bab := string(aba[1]) + string(aba[0]) + string(aba[1])
	return strings.Contains(s, bab)
}

func main() {
	ips := getIPv7Adresses()

	countTLS, countSSL := 0, 0
	for _, ip := range ips {
		if ip.supportsTLS() {
			countTLS++
		}
		if ip.supportsSSL() {
			countSSL++
		}
	}
	fmt.Println("There are", countTLS, "IPv7 addresses supporting TSL, and",
		countSSL, "supporting SSL.")
}
