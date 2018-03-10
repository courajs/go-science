package main

import (
	. "fmt"
)

func main() {
	m1 := make(map[int]int, 2)
	m2 := m1
	m1[7] = 9
	val, ok := m2[7]
	Println(val, ok)
}
