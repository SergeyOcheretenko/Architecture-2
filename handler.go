package lab2

import (
  "io"
  "fmt"
  "log"
  "bufio"
)
// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
  Input io.Reader
  Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
  scanner := bufio.NewScanner(ch.Input)
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
  return nil
}
