package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) (m map[string]int) {
	m = make(map[string]int)
	slice := strings.Fields(s)
	for _,v := range slice  {
		_, ok := m[v]
		if ok {
			m[v] = m[v]  + 1
		} else {
			m[v] = 1
		}
	}
	
	return
}

func WordCountTwo(s string) (m map[string]int) {
	m = make(map[string]int)
	a := strings.Fields(s)
	for _, v := range a {
		m[v]++
	}

	return
}

func main() {
	wc.Test(WordCountTwo)
}