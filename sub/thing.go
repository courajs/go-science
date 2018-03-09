package sub

import . "fmt"
import "runtime"

func Place() {
  pc,file,line,ok := runtime.Caller(0)

  Println(pc, file, line, ok)

  f := runtime.FuncForPC(pc)
  file,line = f.FileLine(pc)
  Println(file, line)

  file,line = f.FileLine(f.Entry())
  Println(f.Name(), file, line)
}
