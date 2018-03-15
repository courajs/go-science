package main

import (
	. "fmt"
)

func act() (int, int) {
	return 1, 2
}

func delegate() (int, int) {
	return act()
}

func main() {
	Println(delegate())
}
