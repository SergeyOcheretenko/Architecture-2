package main

import (
	"flag"
	"fmt"
	"strings"
	"os"
	"errors"
	lab2 "github.com/roman-mazur/architecture-lab-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile = flag.String("f", "", "File with expression to compute")
	outputFile = flag.String("o", "", "File for the result")
)

func catchErr(err error) {
    fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func checkErr(err error) {
    if err != nil {
		catchErr(err)
    }
}

func main() {
	flag.Parse()

	handler := &lab2.ComputeHandler{}
	if *inputExpression != "" && *inputFile != "" {
		catchErr(errors.New("There should be only one way to pass an expression"))
	}
	if *inputExpression != "" {
		handler.Input = strings.NewReader(*inputExpression)
	} else if *inputFile != "" {
		file, err := os.Open(*inputFile)
		checkErr(err)
		defer file.Close()
		handler.Input = file
	} else {
		catchErr(errors.New("Expression not passed"))
	}

	if *outputFile != "" {
		file, err := os.OpenFile(*outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		checkErr(err)
		defer file.Close()
		handler.Output = file
	} else {
		handler.Output = os.Stdout
	} 

	err := handler.Compute()
	checkErr(err)
}
