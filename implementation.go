package lab2

import (
  "strings"
  "unicode"
  "errors"
)

func isInt(s string) bool {
  for _, elem := range s {
    if !unicode.IsDigit(elem) {
      return false
    }
  }
  return true
}

func contains(slice []string, neededValue string) bool {
  for _, elem := range slice {
    if elem == neededValue {
      return true
    }
  }
  return false
}

func checkInputText(text string) error {
  mathSymbols := []string{"+", "/", "*", "-", "^", "div", "%"}
  checker := 0

  elements := strings.Split(Reverse(text), " ")
  for index, elem := range elements {
    if checker < 0 {
      return errors.New("Error in input text")
    } else if contains(mathSymbols, elem) {
      if (index == 0) {
        checker += 2
      } else {
        checker += 1
      }
    } else if isInt(elem) {
      checker -= 1
    } else {
      return errors.New("Should only be numbers or mathematical signs")
    }
  }

  if (checker != 0) {
    return errors.New("Error in input text")
  }

  return nil
}

func Reverse(s string) string {
    runes := strings.Split(s, " ")
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return strings.Join(runes, " ")
}

// TODO: document this function.
// PrefixToPostfix converts
func PrefixToPostfix(input string) (string, error) {
  return Reverse(input), checkInputText(input)
}
