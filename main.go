// Author: Taylor Privat

package main

import (
	"fmt"
	"text/scanner"
	"io/ioutil"
	"os"
	"strings"
)

var (
	tokenQueue []Token
	s scanner.Scanner
)

func main() {
	tokenQueue = []Token{}

	// Get the file name from user and read
	fmt.Println("Enter the name of file (in this directory) you wish to evaluate. e.g. program.txt")
	var prog string
	fmt.Scanln(&prog)
	content, err := ioutil.ReadFile(prog)
	if err != nil {  // If there is an error reading the file express that and exit
		fmt.Println(prog, " does not exists. Exiting")
		os.Exit(1)
	}
	// Otherwise convert to string for tokenization
	text := string(content)

	fmt.Println("Program to evaluate: ")
	fmt.Println(text)

	// Use the scanner to start tokenizing
	s.Init(strings.NewReader(text))
	getChar()
	for (nextToken != scanner.EOF) {
		lex()
	}

	fmt.Println("")
	fmt.Println("Program tokens: ")
	for _, v := range tokenQueue {
		fmt.Println("tok: ", v)
	}

	// Parse the tokens to get AST
	final_root := parseExp()
	fmt.Println("")
	fmt.Println("Progam AST: ")
	printTree(&final_root, 0)

	// Evaluate and return the answer
	answer := evaluate(final_root, empty_env())
	fmt.Println("")
	fmt.Println("The above program returns: ", answer)
}

