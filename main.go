package main
import (
  . "fmt"
)

func main() {
  is := make([]int, 0, 10)
  xs := make([][]int, 0, 10)
  _ = append(xs, []int{
    1,
    2,
  })
  is = append(is, 4)
  _ = append(is, 8)
  i2 := is[:2]
  Println(xs, is, i2)
}
