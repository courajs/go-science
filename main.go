package main
import (
  "time"
  . "fmt"
  . "github.com/jtolds/gls"
)

type Any = interface{}

var c = NewContextManager()
var k = "key"

func printit(label string) {
  id, okid := GetGoroutineId()
  val, okval := c.GetValue(k)
  Println(label, okid, id, okval, val)
}

func printer(label string) func() {
  return func(){printit(label)}
}

func main() {
  c.SetValues(Values(map[Any]Any{k: 4}), func() {
    printit("setvalues")
    Go(printer("Go"))
    go printit("go")
    time.Sleep(time.Second)
  })
}
