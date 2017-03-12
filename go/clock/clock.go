// Clock stub file

// To use the right term, this is the package *clause*.
// You can document general stuff about the package here if you like.
package clock

import (
  "strconv"
  "strings"
  "math"
)

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

type Clock struct {
  Hour int
  Minute int
}

func New(hour, minute int) Clock {
  var remainingHours int
  var remainingNegativeHours int

  if minute >= 60 {
    remainingHours = minute / 60
    minute = minute % 60
  }

  if minute < 0 {
    remainingNegativeHours = 1
    remainingNegativeHours = remainingNegativeHours - ( minute / 60 )
    hour = hour - remainingNegativeHours
    remainder := minute % 60
    minute = 60 + remainder
  }

  hour = hour + remainingHours

  if hour >= 24 {
    hour = hour % 24
  }

  if hour < 0 {
    remainder := hour % 24
    hour = 24 + remainder
  }

  return Clock{
    Hour: hour,
    Minute: minute,
  }
}

func (myClock Clock) String() string {
  hourString := strconv.Itoa(int(math.Abs(float64(myClock.Hour))))
  minuteString := strconv.Itoa(int(math.Abs(float64(myClock.Minute))))
  time := leftPad(hourString, "0", 2) + ":" + leftPad(minuteString, "0", 2)
  return time
}

func (myClock Clock) Add(minutes int) Clock {
  var newMinutes int
  var newHours int
  var totalMinutes int

  // Convert hours to minutes
  // Add hours to the minutes for the clock
  paddingMinutes := (( int(math.Abs(float64(minutes))) / 1440 ) + 1 ) * 1440

  totalMinutes = ( myClock.Hour * 60 ) + myClock.Minute + paddingMinutes

  totalMinutes = ( totalMinutes + minutes ) % 1440

  // Convert the minutes back to hours and minutes
  newHours = ( totalMinutes / 60 ) % 24
  newMinutes = totalMinutes % 60

  myClock.Minute = newMinutes
  myClock.Hour = newHours

  return myClock
}

func leftPad(s string, padStr string, overallLen int) string{
  var strLen int
  var pLen int

  strLen = len(s)
  pLen = overallLen - strLen

  return strings.Repeat(padStr, pLen) + s
}
