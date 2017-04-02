package acronym

import (
  "strings"
  "unicode"
)

const testVersion = 2

func Abbreviate(termString string) string {
  var acro string

  // Split all of the words apart
  words := SplitAllWords(termString)

  for _,word := range words {
    letter := GetFirstLetterUpper(word)
    acro = acro + string(letter)
  }

  // Return the acronym
  return acro
}

func GetFirstLetterUpper (word string) string {
  var upperLetter string

  letter := string(word[0])
  upperLetter = strings.ToUpper(letter)

  return upperLetter
}

func SplitAllWords (phrase string) []string {
  var splitWords []string

  // Split the string by spaces
  words := strings.Fields(phrase)

  // Loop through each of the words and see if we can break it down any further
  for _,word := range words {
    lastSplit := 0

    // If this letter is a dash, split the word up to the letter before this and move along
    dashSplitWords := strings.Split(word, string('-'))

    if (len(dashSplitWords) > 1) {
      for _,dashWord := range dashSplitWords {
        splitWords = append(splitWords, dashWord)
      }
    } else {
      // check if there are multiple upper case letters separated by lower case letters
      for index,letter := range word {
        var previousLetter int

        // If we're at the end of the word, add it to the splitWords array
        if len(word) == index+1 {
          splitWord := string(word[lastSplit:index+1])
          splitWords = append(splitWords, splitWord)
        }

        // If we're not at the end of the word yet
        if (index+1 < len(word) && index >= 1) {
          previousLetter = int(word[index-1])

          // If this letter is upper case and the next letter is not upper case, split the word
          if unicode.IsUpper(letter) && !(unicode.IsUpper(rune(previousLetter))) {
            splitWord := string(word[lastSplit:index-1])
            lastSplit = index

            splitWords = append(splitWords, splitWord)
          }
        }
      }
    }
  }

  return splitWords
}
