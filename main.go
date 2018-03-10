package main

import (
	. "fmt"
)

type S struct {
	m map[int]int
}

func main() {
	var s S
	s.m[4] = 5
	val, ok := s.m[4]
	Println(s, val, ok)
}
