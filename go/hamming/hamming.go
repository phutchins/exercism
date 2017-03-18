package hamming

import (
  "fmt"
)

const testVersion = 5

func Distance(a, b string) (int, error) {
  var hammDiff int

  aLen := len(a)
  bLen := len(b)

  if aLen != bLen {
    return -1, fmt.Errorf("Strings are not the same length")
  }

  for i := 0; i < aLen; i++ {
    if a[i] != b[i] {
      hammDiff++
    }
  }

  return hammDiff, nil
}
