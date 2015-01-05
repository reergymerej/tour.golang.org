package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
  lastNum := 0
  secondToLastNum := 0

  return func() int {
    var nextNum int

    if (lastNum == 0) {
      nextNum = 1
    } else {
      nextNum = lastNum + secondToLastNum
    }

    secondToLastNum = lastNum
    lastNum = nextNum
    return lastNum
  }
}

func main() {
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Println(f())
  }
}
