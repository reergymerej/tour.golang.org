package main

import (
  "fmt"
  "strconv"
)

type IPAddr [4]byte

func (ipAddr IPAddr) String() string {
  var str string = ""

  for i, num := range ipAddr {
    str += strconv.FormatUint(uint64(num), 10)
    if i < len(ipAddr)-1 {
      str += "."
    }
  }

  return str
}

func main() {
  addrs := map[string]IPAddr{
    "loopback":  {127, 0, 0, 1},
    "googleDNS": {8, 8, 8, 8},
  }
  for n, a := range addrs {
    fmt.Printf("%v: %v\n", n, a)
  }
}
