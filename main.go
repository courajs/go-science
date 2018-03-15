package main

import (
	. "fmt"
)

type None struct{}

func seq(n int) []None {
	return make([]None, n)
}

func main() {
	for i := range seq(9) {
		Println(i/3, i%3)
	}
}
