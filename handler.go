package lab2

import (
  "io"
  "bufio"
  "unicode"
  "strings"
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

  elements := strings.Split(text, " ")
  for _, elem := range elements {
    if checker < 0 {
      return errors.New("Error in input text")
    } else if contains(mathSymbols, elem) {
      checker += 2
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

// func catchErr(err error) {
//   fmt.Fprintln(os.Stderr, err)
//   os.Exit(1)
// }

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
  Input io.Reader
  Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
  scanner := bufio.NewScanner(ch.Input)
  scanner.Scan()
  expression := scanner.Text()
  
  if scanner.Err() != nil {
    return scanner.Err()
  }

  resultExpression, calcErr := PrefixToPostfix(expression)
  if calcErr != nil {
    return calcErr
  }

  errorInExpressionSyntax := checkInputText(resultExpression)
  if errorInExpressionSyntax != nil {
    return errorInExpressionSyntax
  }
  
  writer := bufio.NewWriter(ch.Output) 
  _, writeErr := writer.WriteString(resultExpression)
  if writeErr != nil {
    return writeErr
  }
  writer.Flush()
  return nil 
}
