package main

import (
	"flag"
	"fmt"
	"strings"
	"os"
	"log"
	lab2 "github.com/roman-mazur/architecture-lab-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile = flag.String("f", "", "File with expression to compute")
	outputFile = flag.String("o", "", "File for the result")
)

func main() {
	flag.Parse()

	handler := &lab2.ComputeHandler{}
	if *inputExpression != "" {
		handler.Input = strings.NewReader(*inputExpression)
	} else {
		file, err := os.Open(*inputFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		handler.Input = file
	}
	if *outputFile != "" {
		file, err := os.OpenFile(*outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		handler.Output = file
	} else {
		handler.Output = os.Stdout
	}

	res, _ := lab2.PrefixToPostfix("+ 2 2")
	fmt.Println(res)
}
