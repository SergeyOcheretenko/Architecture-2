package lab2

import "strings"

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
  return Reverse(input), nil
}
