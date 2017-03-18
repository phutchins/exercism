package raindrops

import (
  "strconv"
)

const testVersion = 2

func Convert(input int) string {
  var convertedString string

  if ( input % 3 ) == 0 {
    convertedString += "Pling"
  }

  if ( input % 5 ) == 0 {
    convertedString += "Plang"
  }

  if ( input % 7 ) == 0 {
    convertedString += "Plong"
  }

  if (convertedString == "") {
    return strconv.Itoa(input)
  }

  return convertedString
}

// The test program has a benchmark too.  How fast does your Convert convert?
