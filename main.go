package main
import (
  . "fmt"
)

func takeTwo(a,b int) {
  Println(a,b)
}

func takeThree(a,b,c int) {
  Println(a,b,c)
}

func takeVar(a, b int, rest ...int) {
  Println(a, b, rest)
}

func makeNone() {}

func makeOne() int {
  return 1
}

func makeTwo() (int,int) {
  return 1,2
}

func makeThree() (a,b,c int) {
  return 1,2,3
}

func main() {
  takeTwo(makeTwo())
  takeVar(makeTwo())
  takeVar(makeThree())
  // takeTwo(makeThree()) // too many arguments in call to takeTwo
  // takeThree(makeTwo()) // not enough arguments in call to takeThree
  // takeThree(4, makeTwo()) // multiple-value makeTwo() in single-value context
  // takeThree(makeTwo(), 4) // multiple-value makeTwo() in single-value context
  // takeVar(makeNone()) // makeNone() used as value
  // takeVar(makeOne()) // not enough arguments in call to takeVar
}
