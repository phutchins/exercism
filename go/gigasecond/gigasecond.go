// Package clause.
package gigasecond

import (
  "time"
  "math"
)

// Constant declaration.
const testVersion = 4 // find the value in gigasecond_test.go

// Add a gigasecond to the provided startTime
func AddGigasecond(startTime time.Time) time.Time {
  // Get how many seconds are in a gigasecond
  gigaseconds := math.Pow(10, float64(9))

  // Confert gigaseconds to a time.Duration so that we can add it to a date
  gigaDuration := time.Duration(gigaseconds)*time.Second

  // Add gigaDuration to the provided date
  calculatedDate := startTime.Add(gigaDuration)

  return calculatedDate
}
