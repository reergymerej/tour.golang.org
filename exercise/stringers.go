package main

import "fmt"

type IPAddr [4]byte

func (ipAddr IPAddr) String() string {
  var str string = ""

  for _, num := range ipAddr {
    // TODO: This conversion creates an empty string.  Why?
    // TODO: Remove the trailing .
    str += string(num) + "."
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
