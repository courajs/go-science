package main
import (
  . "fmt"
)

type S struct {
  i int
}

type O S

func do(o O) {
  Println(o.i)
}

func main() {
  s := struct{i int}{4}
  do(s)

  // var s2 S = struct{i int}{4}
  // do(s2)
}
