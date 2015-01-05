package main

import (
  "code.google.com/p/go-tour/wc"
  "strings"
)

func WordCount(s string) map[string]int {
  // map of each word and counts of its occurrence
  var wordMap = make(map[string]int)

  // split string into words by whitespace
  // http://golang.org/pkg/strings/#Fields
  var words []string = strings.Fields(s)

  // range form of `for` to loop through array
  // Use `_` to keep the compiler happy since
  // we won't actually use the index.
  for _, word := range words {

    // Use `_` again since we won't use the value.
    _, exists := wordMap[word]

    if !exists {
      // add it to the map and tally it
      wordMap[word] = 1

    } else {

      // increment the existing count
      wordMap[word]++
    }
  }

  return wordMap
}

func main() {
  wc.Test(WordCount)
}
